package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// 全局 Docker client 对象
var cli *client.Client

type DockerExecResult struct {
	ErrorOutput    string
	StandardOutput string
	Time           time.Duration
	ExitCode       int
}

func InitClient() error {
	var err error = nil
	cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	return err
}

func CreateContainer(image string) (string, error) {
	ctx := context.Background()
	config := &container.Config{
		Image:           image,
		Tty:             true,
		NetworkDisabled: true,
	}
	hostConfig := &container.HostConfig{
		AutoRemove: true,
		Resources: container.Resources{
			Memory:     512 * 1024 * 1024, // 512MB
			MemorySwap: 512 * 1024 * 1024, // 512MB (内存 + 交换分区)
			CPUCount:   1,                 // 1 个 CPU
		},
	}
	resp, err := cli.ContainerCreate(ctx, config, hostConfig, nil, nil, "")
	if err != nil {
		return "", err
	}
	return resp.ID, nil
}

func StartContainer(containerId string) error {
	ctx := context.Background()
	return cli.ContainerStart(ctx, containerId, container.StartOptions{})
}

func ExecInContainer(containerID string, timeoutSeconds int, command []string) (*DockerExecResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSeconds)*time.Second)
	defer cancel()

	execConfig := container.ExecOptions{
		Cmd:          command,
		AttachStdin:  true,
		AttachStdout: true,
		AttachStderr: true,
	}

	response, err := cli.ContainerExecCreate(ctx, containerID, execConfig)
	if err != nil {
		return &DockerExecResult{}, err
	}

	start := time.Now()

	attachResp, err := cli.ContainerExecAttach(ctx, response.ID, container.ExecAttachOptions{})
	if err != nil {
		return &DockerExecResult{}, err
	}
	defer attachResp.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	if _, err := stdcopy.StdCopy(&stdoutBuf, &stderrBuf, attachResp.Reader); err != nil {
		return &DockerExecResult{}, err
	}

	execInspect, err := cli.ContainerExecInspect(ctx, response.ID)
	if err != nil {
		return &DockerExecResult{}, err
	}

	return &DockerExecResult{
		StandardOutput: stdoutBuf.String(),
		ErrorOutput:    stderrBuf.String(),
		Time:           time.Since(start),
		ExitCode:       execInspect.ExitCode,
	}, nil
}

func CleanupContainer(containerID string, force bool) error {
	ctx := context.Background()
	var timeout int
	if force {
		timeout = 0
	} else {
		timeout = 10
	}
	err := cli.ContainerStop(ctx, containerID, container.StopOptions{Timeout: &timeout})
	if err != nil {
		return err
	}
	err = cli.ContainerRemove(ctx, containerID, container.RemoveOptions{RemoveVolumes: true, RemoveLinks: false, Force: true})
	if err != nil {
		return err
	}
	return nil
}

func StringToContainerFile(ctx context.Context, containerID string, dstPath string, fileName string, str string) error {

	var buf bytes.Buffer
	tw := tar.NewWriter(&buf)

	content := []byte(str)

	hdr := &tar.Header{
		Name: fileName,
		Mode: 0644,
		Size: int64(len(content)),
	}

	if err := tw.WriteHeader(hdr); err != nil {
		return err
	}

	if _, err := tw.Write(content); err != nil {
		return err
	}

	if err := tw.Close(); err != nil {
		return err
	}

	err := cli.CopyToContainer(ctx, containerID, dstPath, &buf, container.CopyToContainerOptions{})
	if err != nil {
		return err
	}

	return nil
}
