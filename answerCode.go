package answer

import (
	"errors"
)

/*
ConvertCodeToStatus - assistent funktion.
Dazutun fur konvertierung code in Nachricht.
*/
func ConvertCodeToStatus(code ErrorCode) (string, error) {
	if ok := StatusCode[code]; ok == "" {
		return "", errors.New("такой код не существует")
	}
	return StatusCode[code], nil
}

// AppendCode - dazutun Individuell code das weitere Verwendung.
func AppendCode(code ErrorCode, response string) error {
	if ok := StatusCode[code]; ok != "" {
		return errors.New("такой код уже есть")
	}
	StatusCode[code] = response
	return nil
}

// ChangeCode - verändern bestehende in map code.
func ChangeCode(code ErrorCode, response string) error {
	if ok := StatusCode[code]; ok == "" {
		return errors.New("такой код не существует")
	}
	StatusCode[code] = response
	return nil
}

// DeleteCode - entfernt Individuell code.
func DeleteCode(code ErrorCode) error {
	if ok := StatusCode[code]; ok == "" {
		return errors.New("такой код не существует")
	}
	delete(StatusCode, code)
	return nil
}
