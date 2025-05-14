package sandbox

import (
	"backend/code-sandbox/internal/docker"
	"backend/code-sandbox/internal/strategies"
	"backend/code-sandbox/types"
	"context"
)

type Sandbox struct {
	dockerClient *docker.DockerClient
}

func NewSandbox() (*Sandbox, error) {
	dockerClient, err := docker.NewDockerClient()
	if err != nil {
		return nil, err
	}
	return &Sandbox{dockerClient: dockerClient}, nil
}

func (s *Sandbox) RunCode(ctx context.Context, config *types.RunConfig) (*types.Result, error) {
	// 两类错误：系统错误返回 error，业务错误返回 Result
	languageStrategy, err := strategies.GetStrategy(config, s.dockerClient)
	if err != nil {
		return nil, err
	}

	// 准备
	err = languageStrategy.Prepare(ctx)
	if err != nil {
		return nil, err
	}

	// 编译
	result, err := languageStrategy.Compile(ctx)
	if err != nil {
		return nil, err
	}
	if result != nil {
		return result, nil
	}

	// 运行
	res, err := languageStrategy.Execute(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
