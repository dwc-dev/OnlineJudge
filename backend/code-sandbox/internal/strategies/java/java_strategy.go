package java

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

type JavaStrategy struct {
	code        string
	inputList   []string
	containerID string
}

func NewJavaStrategy(code string, inputList []string) *JavaStrategy {
	return &JavaStrategy{code: code, inputList: inputList}
}

func (s *JavaStrategy) Prepare() error {
	var err error
	s.containerID, err = docker.CreateContainer("openjdk")
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
	return docker.StringToContainerFile(context.Background(), s.containerID, "/sandbox", "Main.java", s.code)
}

func (s *JavaStrategy) Compile() error {
	result, err := docker.ExecInContainer(s.containerID, ExecTimeOutSeconds,
		[]string{"javac", "/sandbox/Main.java"})
	if err != nil {
		return err
	}
	if result.ExitCode != 0 {
		return fmt.Errorf("编译失败,错误输出:\n%s", result.ErrorOutput)
	}
	return nil
}

func (s *JavaStrategy) Execute() ([]types.ExecResult, error) {
	defer docker.CleanupContainer(s.containerID, true)
	var res []types.ExecResult
	for idx := range len(s.inputList) {
		inputPath := "/sandbox/input" + strconv.Itoa(idx) + ".txt"
		result, err := docker.ExecInContainer(s.containerID, ExecTimeOutSeconds,
			[]string{"sh", "-c", "java -cp /sandbox/ Main < " + inputPath})
		if err != nil {
			return []types.ExecResult{}, err
		}
		if result.ExitCode != 0 {
			return []types.ExecResult{}, fmt.Errorf("运行失败,错误输出:\n%s", result.ErrorOutput)
		}
		res = append(res, types.ExecResult{
			Output:      result.StandardOutput,
			Time:        result.Time,
			MemoryUsage: 0,
		})
	}
	return res, nil
}
