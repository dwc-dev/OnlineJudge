package main

import (
	"code-sandbox/internal/codes"
	"code-sandbox/internal/docker"
	"code-sandbox/internal/sandbox"
	"fmt"
)

func main() {
	docker.InitClient()

	res, err := sandbox.RunCode("c", codes.CCode, codes.CInput)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("c:\n", res)
	}

	res, err = sandbox.RunCode("cpp", codes.CppCode, codes.CppInput)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("cpp:\n", res)
	}

	res, err = sandbox.RunCode("java", codes.JavaCode, codes.JavaInput)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("java:\n", res)
	}

	res, err = sandbox.RunCode("python", codes.PythonCode, codes.PythonInput)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("python:\n", res)
	}

	res, err = sandbox.RunCode("golang", codes.GoCode, codes.GoInput)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("golang:\n", res)
	}

	res, err = sandbox.RunCode("rust", codes.RustCode, codes.RustInput)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("rust:\n", res)
	}
}
