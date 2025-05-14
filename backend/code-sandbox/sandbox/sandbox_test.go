package sandbox

import (
	"backend/code-sandbox/internal/codes"
	"backend/code-sandbox/types"
	"context"
	"encoding/json"
	"testing"
)

func TestCCode(t *testing.T) {
	// t.Skip("skip c code test")
	sandbox, err := NewSandbox()
	if err != nil {
		t.Fatal(err)
	}

	config := &types.RunConfig{
		Language:           "c",
		Code:               codes.CCode,
		InputList:          codes.CInput,
		MemoryLimitMiB:     128,
		StackLimitMiB:      8,
		TimeoutMillisecond: 1000,
	}

	result, err := sandbox.RunCode(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}

	json, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(json))
}

func TestCppCode(t *testing.T) {
	// t.Skip("skip cpp code test")
	sandbox, err := NewSandbox()
	if err != nil {
		t.Fatal(err)
	}

	config := &types.RunConfig{
		Language:           "cpp",
		Code:               codes.CppCode,
		InputList:          codes.CppInput,
		MemoryLimitMiB:     128,
		StackLimitMiB:      8,
		TimeoutMillisecond: 1000,
	}

	result, err := sandbox.RunCode(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}

	json, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(json))
}

func TestJavaCode(t *testing.T) {
	// t.Skip("skip java code test")
	sandbox, err := NewSandbox()
	if err != nil {
		t.Fatal(err)
	}

	config := &types.RunConfig{
		Language:           "java",
		Code:               codes.JavaCode,
		InputList:          codes.JavaInput,
		MemoryLimitMiB:     128,
		StackLimitMiB:      8,
		TimeoutMillisecond: 1000,
	}

	result, err := sandbox.RunCode(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}

	json, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(json))
}

func TestJavaCodeStackOverflow(t *testing.T) {
	// t.Skip("skip java code stack overflow test")
	sandbox, err := NewSandbox()
	if err != nil {
		t.Fatal(err)
	}
	config := &types.RunConfig{
		Language:           "java",
		Code:               codes.JavaCodeStackOverflow,
		InputList:          codes.JavaInputStackOverflow,
		MemoryLimitMiB:     128,
		StackLimitMiB:      8,
		TimeoutMillisecond: 1000,
	}
	result, err := sandbox.RunCode(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}
	json, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(json))
}

func TestPythonCode(t *testing.T) {
	// t.Skip("skip python code test")
	sandbox, err := NewSandbox()
	if err != nil {
		t.Fatal(err)
	}

	config := &types.RunConfig{
		Language:           "python",
		Code:               codes.PythonCode,
		InputList:          codes.PythonInput,
		MemoryLimitMiB:     128,
		StackLimitMiB:      8,
		TimeoutMillisecond: 1000,
	}

	result, err := sandbox.RunCode(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}

	json, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(json))
}

func TestPythonCodeStackOverflow(t *testing.T) {
	// t.Skip("skip python code stack overflow test")
	sandbox, err := NewSandbox()
	if err != nil {
		t.Fatal(err)
	}
	config := &types.RunConfig{
		Language:           "python",
		Code:               codes.PythonCodeStackOverflow,
		InputList:          codes.PythonInputStackOverflow,
		MemoryLimitMiB:     128,
		StackLimitMiB:      4,
		TimeoutMillisecond: 1000,
	}
	result, err := sandbox.RunCode(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}
	json, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(json))
}

func TestGolangCode(t *testing.T) {
	// t.Skip("skip golang code test")
	sandbox, err := NewSandbox()
	if err != nil {
		t.Fatal(err)
	}
	config := &types.RunConfig{
		Language:           "golang",
		Code:               codes.GolangCode,
		InputList:          codes.GolangInput,
		MemoryLimitMiB:     128,
		StackLimitMiB:      8,
		TimeoutMillisecond: 1000,
	}
	result, err := sandbox.RunCode(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}
	json, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(json))
}

func TestRustCode(t *testing.T) {
	// t.Skip("skip rust code test")
	sandbox, err := NewSandbox()
	if err != nil {
		t.Fatal(err)
	}
	config := &types.RunConfig{
		Language:       "rust",
		Code:           codes.RustCode,
		InputList:      codes.RustInput,
		MemoryLimitMiB: 128,
		StackLimitMiB:  8,
	}
	result, err := sandbox.RunCode(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}
	json, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(json))
}
