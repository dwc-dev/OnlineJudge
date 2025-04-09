package strategies

import (
	"backend/code-sandbox/internal/strategies/c"
	"backend/code-sandbox/internal/strategies/cpp"
	"backend/code-sandbox/internal/strategies/golang"
	"backend/code-sandbox/internal/strategies/java"
	"backend/code-sandbox/internal/strategies/python"
	"backend/code-sandbox/internal/strategies/rust"
	"backend/code-sandbox/internal/types"
	"errors"
)

func GetStrategy(language string, code string, inputList []string) (types.LanguageStrategy, error) {
	switch language {
	case "c":
		return c.NewCStrategy(code, inputList), nil
	case "cpp":
		return cpp.NewCppStrategy(code, inputList), nil
	case "rust":
		return rust.NewRustStrategy(code, inputList), nil
	case "python":
		return python.NewPythonStrategy(code, inputList), nil
	case "java":
		return java.NewJavaStrategy(code, inputList), nil
	case "golang":
		return golang.NewGolangStrategy(code, inputList), nil
	}
	return nil, errors.New("不支持的语言")
}
