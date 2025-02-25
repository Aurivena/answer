package ansCode

import (
	"errors"
	"github.com/Aurivena/answer/pkg/ansError"
)

func ConvertCodeToStatus(code ansError.ErrorCode, stc map[ansError.ErrorCode]string) (string, error) {
	if ok := stc[code]; ok == "" {
		return "", errors.New("такой код не существует")
	}
	return stc[code], nil
}

func AppendCode(code ansError.ErrorCode, response string, stc map[ansError.ErrorCode]string) error {
	if ok := stc[code]; ok != "" {
		return errors.New("такой код уже есть")
	}
	stc[code] = response
	return nil
}

func DeleteCode(code ansError.ErrorCode, stc map[ansError.ErrorCode]string) error {
	if ok := stc[code]; ok == "" {
		return errors.New("такой код не существует")
	}
	delete(stc, code)
	return nil
}
