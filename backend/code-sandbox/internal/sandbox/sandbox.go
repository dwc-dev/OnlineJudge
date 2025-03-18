package sandbox

import (
	"code-sandbox/internal/strategies"
	"code-sandbox/internal/types"
)

func RunCode(language, code string, inputList []string) ([]types.ExecResult, error) {
	languageStrategy, err := strategies.GetStrategy(language, code, inputList)
	if err != nil {
		return []types.ExecResult{}, err
	}
	err = languageStrategy.Prepare()
	if err != nil {
		return []types.ExecResult{}, err
	}
	err = languageStrategy.Compile()
	if err != nil {
		return []types.ExecResult{}, err
	}
	res, err := languageStrategy.Execute()
	if err != nil {
		return []types.ExecResult{}, err
	}
	return res, nil
}
