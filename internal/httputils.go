package internal

func IsSuccessful(statusCode int) bool {
	return (statusCode >= 200 && statusCode <= 299)
}

func isFail(statusCode int) bool {
	return (statusCode >= 400 && statusCode <= 599)
}
