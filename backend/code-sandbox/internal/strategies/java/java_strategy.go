package java

import (
	"backend/code-sandbox/internal/docker"
	"backend/code-sandbox/internal/tools"
	"backend/code-sandbox/types"
	"context"
	"fmt"
	"strconv"
	"strings"
)

type JavaStrategy struct {
	code               string
	memoryLimitMiB     int64
	stackLimitMiB      int64
	timeoutMillisecond int
	inputList          []string
	containerID        string
	dockerClient       *docker.DockerClient
}

var javaRunner = `
#!/bin/bash

input_file="$1"
timeout_sec="$2"
stack_limit_mb="$3"  # 第三个参数，指定 Java 栈大小，单位：MB
heap_limit_mb="$4"  # 第四个参数，指定 Java 堆大小，单位：MB

# 开始时间
start=$(date +%s%3N)

# 使用 timeout + /usr/bin/time 运行 Java 程序
timeout "${timeout_sec}"s /usr/bin/time -v java -Xss"${stack_limit_mb}"m -Xmx"${heap_limit_mb}"m -cp /sandbox Main < "$input_file"
exit_code=$?

# 结束时间
end=$(date +%s%3N)

# 耗时（ms）
elapsed=$((end - start))

# 输出运行时间到 stderr
echo "Time elapsed(ms): ${elapsed}" >&2

# 返回程序或 timeout 的退出码
exit $exit_code
`

func NewJavaStrategy(runConfig *types.RunConfig, dockerClient *docker.DockerClient) *JavaStrategy {
	return &JavaStrategy{
		code:               runConfig.Code,
		inputList:          runConfig.InputList,
		memoryLimitMiB:     runConfig.MemoryLimitMiB,
		stackLimitMiB:      runConfig.StackLimitMiB,
		timeoutMillisecond: runConfig.TimeoutMillisecond,
		dockerClient:       dockerClient,
	}
}

func (s *JavaStrategy) Prepare(ctx context.Context) error {
	var err error
	s.containerID, err = s.dockerClient.CreateContainer(ctx, "oj-java", s.memoryLimitMiB+128, 8, 1)
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
	err = s.dockerClient.StringToContainerFile(ctx, s.containerID, "/sandbox", "Main.java", s.code)
	if err != nil {
		return err
	}
	err = s.dockerClient.StringToContainerFile(ctx, s.containerID, "/sandbox", "java_runner.sh", javaRunner)
	if err != nil {
		return err
	}
	_, err = s.dockerClient.ExecInContainer(ctx, s.containerID, []string{"chmod", "+x", "/sandbox/java_runner.sh"})
	if err != nil {
		return err
	}
	return nil
}

func (s *JavaStrategy) Compile(ctx context.Context) (*types.Result, error) {
	result, err := s.dockerClient.ExecInContainer(ctx, s.containerID, []string{"javac", "/sandbox/Main.java"})
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

func (s *JavaStrategy) Execute(ctx context.Context) (*types.Result, error) {
	defer s.dockerClient.CleanupContainer(ctx, s.containerID, true)
	var res []*types.ExecDetail
	for idx := range len(s.inputList) {
		inputPath := "/sandbox/input_" + strconv.Itoa(idx) + ".txt"
		timeout := float64(s.timeoutMillisecond) / 1000 // 毫秒转秒
		result, err := s.dockerClient.ExecInContainer(ctx, s.containerID,
			[]string{"sh", "-c",
				"/sandbox/java_runner.sh " + inputPath + " " +
					fmt.Sprintf("%.3f", timeout) + " " +
					strconv.Itoa(int(s.stackLimitMiB)) + " " +
					strconv.Itoa(int(s.memoryLimitMiB))})
		if err != nil {
			return nil, err
		}
		isTimeout := false
		isRuntimeError := false
		isMemoryLimitOut := false
		isStackOverflow := false
		memoryUsage, elapsedTime := tools.Parse(result.ErrorOutput) // 从错误输出中解析出内存使用和运行时间
		if result.ExitCode == 124 {
			// 如果退出代码为124，则认为超时
			result.StandardOutput = ""
			isTimeout = true
		} else if result.ExitCode == 1 && strings.Contains(result.ErrorOutput, "java.lang.StackOverflowError") {
			// 栈溢出
			result.StandardOutput = ""
			isStackOverflow = true
		} else if result.ExitCode == 137 || memoryUsage > s.memoryLimitMiB*1024 {
			//如果这个程序内存超出限制，它会被容器限制强制结束，退出码为137
			result.StandardOutput = ""
			isMemoryLimitOut = true
		} else if result.ExitCode == 1 && strings.Contains(result.ErrorOutput, "java.lang.OutOfMemoryError") {
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
