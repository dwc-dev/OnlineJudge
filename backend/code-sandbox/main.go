package main

import (
	"code-sandbox/internal/docker"
	"code-sandbox/internal/sandbox"
	"fmt"
)

func main() {
	var testCode = `#include <stdio.h>
					int main() {
						int a, b;
						scanf("%d%d", &a, &b);
						printf("%d\n", a + b);
						return 0;
					}`
	docker.InitClient()
	_, err := sandbox.RunCode("c", testCode, []string{"1 2", "3 4", "5 6"})
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Success")
	}
}
