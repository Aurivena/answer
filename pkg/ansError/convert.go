package ansError

func ConvertCodeToStatus(processStatus int) string {
	return statusCode[processStatus]
}
