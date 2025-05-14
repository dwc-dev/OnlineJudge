package strategies

import (
	"backend/code-sandbox/internal/docker"
	"backend/code-sandbox/internal/strategies/c"
	"backend/code-sandbox/internal/strategies/cpp"
	"backend/code-sandbox/internal/strategies/golang"
	"backend/code-sandbox/internal/strategies/java"
	"backend/code-sandbox/internal/strategies/python"
	"backend/code-sandbox/internal/strategies/rust"
	"backend/code-sandbox/types"
	"context"
	"errors"
)

type LanguageStrategy interface {
	Prepare(context.Context) error
	Compile(context.Context) (*types.Result, error)
	Execute(context.Context) (*types.Result, error)
}

func GetStrategy(config *types.RunConfig, dockerClient *docker.DockerClient) (LanguageStrategy, error) {
	switch config.Language {
	case "c":
		return c.NewCStrategy(config, dockerClient), nil
	case "cpp":
		return cpp.NewCppStrategy(config, dockerClient), nil
	case "java":
		return java.NewJavaStrategy(config, dockerClient), nil
	case "python":
		return python.NewPythonStrategy(config, dockerClient), nil
	case "golang":
		return golang.NewGolangStrategy(config, dockerClient), nil
	case "rust":
		return rust.NewRustStrategy(config, dockerClient), nil
	}
	return nil, errors.New("unsupported language")
}
