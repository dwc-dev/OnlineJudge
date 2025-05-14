package rust

import (
	"backend/code-sandbox/internal/docker"
	"backend/code-sandbox/internal/tools"
	"backend/code-sandbox/types"
	"context"
	"fmt"
	"strconv"
	"strings"
)

type RustStrategy struct {
	code               string
	memoryLimitMiB     int64
	stackLimitMiB      int64
	timeoutMillisecond int
	inputList          []string
	containerID        string
	dockerClient       *docker.DockerClient
}

var rustRunner = `
#!/bin/bash

input_file="$1"
timeout_sec="$2"

# 开始时间
start=$(date +%s%3N)

# 运行程序并通过 /usr/bin/time 获取内存信息
timeout "$timeout_sec"s /usr/bin/time -v /sandbox/main < "$input_file"
exit_code=$?  # 获取 timeout 的退出码，如果超时，退出码为124

# 结束时间
end=$(date +%s%3N)

# 耗时
elapsed=$((end - start))

# 输出运行时间到 stderr
echo "Time elapsed(ms): ${elapsed}" >&2

# 返回退出码
exit $exit_code
`

func NewRustStrategy(runConfig *types.RunConfig, dockerClient *docker.DockerClient) *RustStrategy {
	return &RustStrategy{
		code:               runConfig.Code,
		inputList:          runConfig.InputList,
		memoryLimitMiB:     runConfig.MemoryLimitMiB,
		stackLimitMiB:      runConfig.StackLimitMiB,
		timeoutMillisecond: runConfig.TimeoutMillisecond,
		dockerClient:       dockerClient,
	}
}

func (s *RustStrategy) Prepare(ctx context.Context) error {
	var err error
	s.containerID, err = s.dockerClient.CreateContainer(ctx, "oj-rust", s.memoryLimitMiB+128, s.stackLimitMiB, 1)
	if err != nil {
		return err
	}
	err = s.dockerClient.StartContainer(ctx, s.containerID)
	if err != nil {
		return err
	}
	_, err = s.dockerClient.ExecInContainer(ctx, s.containerID, []string{"mkdir", "/sandbox"})
	if err != nil {
		return err
	}
	for idx, input := range s.inputList {
		err = s.dockerClient.StringToContainerFile(ctx, s.containerID, "/sandbox", "input_"+strconv.Itoa(idx)+".txt", input)
		if err != nil {
			return err
		}
	}
	err = s.dockerClient.StringToContainerFile(ctx, s.containerID, "/sandbox", "main.rs", s.code)
	if err != nil {
		return err
	}
	err = s.dockerClient.StringToContainerFile(ctx, s.containerID, "/sandbox", "rust_runner.sh", rustRunner)
	if err != nil {
		return err
	}
	_, err = s.dockerClient.ExecInContainer(ctx, s.containerID, []string{"chmod", "+x", "/sandbox/rust_runner.sh"})
	if err != nil {
		return err
	}
	return nil
}

func (s *RustStrategy) Compile(ctx context.Context) (*types.Result, error) {
	result, err := s.dockerClient.ExecInContainer(ctx, s.containerID,
		[]string{"rustc", "-o", "/sandbox/main", "/sandbox/main.rs"})
	if err != nil {
		return nil, err
	}
	if result.ExitCode != 0 {
		return &types.Result{
			CompileError:       true,
			CompileErrorOutput: result.ErrorOutput,
		}, nil
	}
	return nil, nil
}

func (s *RustStrategy) Execute(ctx context.Context) (*types.Result, error) {
	defer s.dockerClient.CleanupContainer(ctx, s.containerID, true)
	var res []*types.ExecDetail
	for idx := range len(s.inputList) {
		inputPath := "/sandbox/input_" + strconv.Itoa(idx) + ".txt"
		timeout := float64(s.timeoutMillisecond) / 1000 // 毫秒转秒
		result, err := s.dockerClient.ExecInContainer(ctx, s.containerID,
			[]string{"sh", "-c",
				"/sandbox/rust_runner.sh " + inputPath + " " + fmt.Sprintf("%.3f", timeout)})
		if err != nil {
			return nil, err
		}
		isTimeout := false
		isRuntimeError := false
		isMemoryLimitOut := false
		isStackOverflow := false
		// 从错误输出中解析出内存使用和运行时间
		memoryUsage, elapsedTime := tools.Parse(result.ErrorOutput)
		if result.ExitCode == 124 {
			// 如果退出代码为124，则认为超时
			result.StandardOutput = ""
			isTimeout = true
		} else if result.ExitCode == 139 || strings.Contains(result.ErrorOutput, "signal 11") {
			// 如果退出代码为139，则认为栈溢出
			result.StandardOutput = ""
			isStackOverflow = true
		} else if result.ExitCode == 137 {
			// 如果这个程序内存超出限制，它会被容器限制强制结束，退出码为137
			result.StandardOutput = ""
			isMemoryLimitOut = true
		} else if memoryUsage > s.memoryLimitMiB*1024 {
			// 内存超出限制
			result.StandardOutput = ""
			isMemoryLimitOut = true
		} else if result.ExitCode != 0 {
			// 如果退出代码不为0，则认为运行时错误
			result.StandardOutput = ""
			isRuntimeError = true
		}
		res = append(res, &types.ExecDetail{
			Output:           result.StandardOutput,
			TimeMilliseconds: elapsedTime, // 单位：ms
			MemoryUsage:      memoryUsage, // 单位：kbytes
			Timeout:          isTimeout,
			RuntimeError:     isRuntimeError,
			MemoryOut:        isMemoryLimitOut,
			StackOverflow:    isStackOverflow,
		})
	}
	return &types.Result{
		ExecDetails:  res,
		CompileError: false,
	}, nil
}
