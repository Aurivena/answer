package answer

import (
	"errors"
)

func ConvertCodeToStatus(code ErrorCode, stc map[ErrorCode]string) (string, error) {
	if ok := stc[code]; ok == "" {
		return "", errors.New("такой код не существует")
	}
	return stc[code], nil
}

func AppendCode(code ErrorCode, response string, stc map[ErrorCode]string) error {
	if ok := stc[code]; ok != "" {
		return errors.New("такой код уже есть")
	}
	stc[code] = response
	return nil
}

func ChangeCode(code ErrorCode, response string, stc map[ErrorCode]string) error {
	if ok := stc[code]; ok == "" {
		return errors.New("такой код не существует")
	}
	stc[code] = response
	return nil
}

func DeleteCode(code ErrorCode, stc map[ErrorCode]string) error {
	if ok := stc[code]; ok == "" {
		return errors.New("такой код не существует")
	}
	delete(stc, code)
	return nil
}
