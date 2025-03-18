package strategies

import (
	"code-sandbox/internal/types"
)

func GetStrategy(language string, code string, inputList []string) (types.LanguageStrategy, error) {
	switch language {
	case "c":
		return NewC_CPP_Strategy(code, inputList), nil
	case "cpp":
		return nil, nil
	}
	return nil, nil
}
