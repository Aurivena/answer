package ansError

func ConvertCodeToStatus(processStatus int) string {
	return statusCode[processStatus]
}

func AppendCode(code int, msg string) {
	statusCode[code] = msg
}

func DeleteCode(code int) {
	delete(statusCode, code)
}
