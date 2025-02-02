package ansError

func ConvertCodeToStatus(status int) string {
	return statusCode[status]
}
