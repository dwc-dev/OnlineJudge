package golang

import (
	"backend/code-sandbox/docker"
	"backend/code-sandbox/internal/types"
	"context"
	"fmt"
	"strconv"
)

var (
	ExecTimeOutSeconds = 5
)

type GolangStrategy struct {
	code        string
	inputList   []string
	containerID string
}

func NewGolangStrategy(code string, inputList []string) *GolangStrategy {
	return &GolangStrategy{code: code, inputList: inputList}
}

func (s *GolangStrategy) Prepare() error {
	var err error
	s.containerID, err = docker.CreateContainer("golang")
	if err != nil {
		return err
	}
	err = docker.StartContainer(s.containerID)
	if err != nil {
		return err
	}
	_, err = docker.ExecInContainer(s.containerID, ExecTimeOutSeconds, []string{"mkdir", "/sandbox"})
	if err != nil {
		return err
	}
	for idx, input := range s.inputList {
		err = docker.StringToContainerFile(context.Background(), s.containerID, "/sandbox", "input"+strconv.Itoa(idx)+".txt", input)
		if err != nil {
			return err
		}
	}
	return docker.StringToContainerFile(context.Background(), s.containerID, "/sandbox", "main.go", s.code)
}

func (s *GolangStrategy) Compile() error {
	fmt.Println("开始编译")
	result, err := docker.ExecInContainer(s.containerID, ExecTimeOutSeconds,
		[]string{"go", "build", "-o", "/sandbox/main", "/sandbox/main.go"})
	if err != nil {
		return err
	}
	if result.ExitCode != 0 {
		return fmt.Errorf("编译失败,错误输出:\n%s", result.ErrorOutput)
	}
	fmt.Println("编译完成")
	return nil
}

func (s *GolangStrategy) Execute() ([]types.ExecResult, error) {
	fmt.Println("开始运行")
	// defer docker.CleanupContainer(s.containerID, true)
	var res []types.ExecResult
	for idx := range len(s.inputList) {
		inputPath := "/sandbox/input" + strconv.Itoa(idx) + ".txt"
		result, err := docker.ExecInContainer(s.containerID, ExecTimeOutSeconds,
			[]string{"sh", "-c", "/sandbox/main < " + inputPath})
		if err != nil {
			return []types.ExecResult{}, err
		}
		if result.ExitCode != 0 {
			return []types.ExecResult{}, fmt.Errorf("运行失败,错误输出:\n%s", result.ErrorOutput)
		}
		fmt.Println("运行完成")
		res = append(res, types.ExecResult{
			Output:      result.StandardOutput,
			Time:        result.Time,
			MemoryUsage: 0,
		})
	}
	docker.CleanupContainer(s.containerID, true)
	return res, nil
}
