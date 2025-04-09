package python

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

type PythonStrategy struct {
	code        string
	inputList   []string
	containerID string
}

func NewPythonStrategy(code string, inputList []string) *PythonStrategy {
	return &PythonStrategy{code: code, inputList: inputList}
}

func (s *PythonStrategy) Prepare() error {
	var err error
	s.containerID, err = docker.CreateContainer("python")
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
	return docker.StringToContainerFile(context.Background(), s.containerID, "/sandbox", "main.py", s.code)
}

func (s *PythonStrategy) Compile() error {
	// python无需编译
	return nil
}

func (s *PythonStrategy) Execute() ([]types.ExecResult, error) {
	defer docker.CleanupContainer(s.containerID, true)
	var res []types.ExecResult
	for idx := range len(s.inputList) {
		inputPath := "/sandbox/input" + strconv.Itoa(idx) + ".txt"
		result, err := docker.ExecInContainer(s.containerID, ExecTimeOutSeconds,
			[]string{"sh", "-c", "python /sandbox/main.py < " + inputPath})
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
