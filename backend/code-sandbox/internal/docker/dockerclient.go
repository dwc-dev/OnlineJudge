package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"github.com/docker/go-units"
)

type DockerClient struct {
	cli *client.Client
}

type DockerExecResult struct {
	ErrorOutput    string
	StandardOutput string
	ExitCode       int
}

func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}
	return &DockerClient{cli: cli}, nil
}

func (c *DockerClient) CreateContainer(ctx context.Context, image string, memoryLimitMiB int64, stackLimitMiB int64, cpuCount int64) (string, error) {
	if c.cli == nil {
		return "", fmt.Errorf("docker client not initialized")
	}
	config := &container.Config{
		Image:           image,
		Tty:             true,
		NetworkDisabled: true,
	}
	ulimits := []*units.Ulimit{
		{
			Name: "stack",
			Soft: stackLimitMiB * 1024 * 1024,
			Hard: stackLimitMiB * 1024 * 1024,
		},
	}
	hostConfig := &container.HostConfig{
		AutoRemove: true,
		Resources: container.Resources{
			Memory:     memoryLimitMiB * 1024 * 1024,
			MemorySwap: memoryLimitMiB * 1024 * 1024,
			CPUCount:   cpuCount,
			Ulimits:    ulimits,
		},
	}
	resp, err := c.cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func (c *DockerClient) StartContainer(ctx context.Context, containerID string) error {
	if c.cli == nil {
		return fmt.Errorf("docker client not initialized")
	}
	return c.cli.ContainerStart(ctx, containerID, container.StartOptions{})
}

func (c *DockerClient) ExecInContainer(ctx context.Context, containerID string, command []string) (*DockerExecResult, error) {
	if c.cli == nil {
		return nil, fmt.Errorf("docker client not initialized")
	}

	execConfig := container.ExecOptions{
		Cmd:          command,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	}

	response, err := c.cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		return nil, err
	}

	attachResp, err := c.cli.ContainerExecAttach(ctx, response.ID, container.ExecAttachOptions{})
	if err != nil {
		return nil, err
	}
	defer attachResp.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	if _, err := stdcopy.StdCopy(&stdoutBuf, &stderrBuf, attachResp.Reader); err != nil {
		return nil, err
	}

	execInspect, err := c.cli.ContainerExecInspect(ctx, response.ID)
	if err != nil {
		return nil, err
	}

	return &DockerExecResult{
		StandardOutput: stdoutBuf.String(),
		ErrorOutput:    stderrBuf.String(),
		ExitCode:       execInspect.ExitCode,
	}, nil
}

func (c *DockerClient) CleanupContainer(ctx context.Context, containerID string, force bool) error {
	if c.cli == nil {
		return fmt.Errorf("docker client not initialized")
	}
	var timeout int
	if force {
		timeout = 0
	} else {
		timeout = 10
	}
	err := c.cli.ContainerStop(ctx, containerID, container.StopOptions{Timeout: &timeout})
	if err != nil {
		return err
	}
	err = c.cli.ContainerRemove(ctx, containerID, container.RemoveOptions{RemoveVolumes: true, RemoveLinks: true, Force: true})
	if err != nil {
		return err
	}
	return nil
}

func (c *DockerClient) StringToContainerFile(ctx context.Context, containerID string, dstPath string, fileName string, str string) error {
	if c.cli == nil {
		return fmt.Errorf("docker client not initialized")
	}

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)
	content := []byte(str)
	header := &tar.Header{
		Name: fileName,
		Mode: 0644,
		Size: int64(len(content)),
	}
	if err := tw.WriteHeader(header); err != nil {
		return err
	}
	if _, err := tw.Write(content); err != nil {
		return err
	}
	if err := tw.Close(); err != nil {
		return err
	}

	return c.cli.CopyToContainer(ctx, containerID, dstPath, &buf, container.CopyToContainerOptions{})
}
