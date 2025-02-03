package ansError

import "errors"

func ConvertCodeToStatus(code int, stc map[int]string) (string, error) {
	if ok := stc[code]; ok == "" {
		return "", errors.New("такой код не существует")
	}
	return stc[code], nil
}

func AppendCode(code int, response string, stc map[int]string) error {
	if ok := stc[code]; ok != "" {
		return errors.New("такой код уже есть")
	}
	stc[code] = response
	return nil
}

func DeleteCode(code int, stc map[int]string) error {
	if ok := stc[code]; ok == "" {
		return errors.New("такой код не существует")
	}
	delete(stc, code)
	return nil
}
