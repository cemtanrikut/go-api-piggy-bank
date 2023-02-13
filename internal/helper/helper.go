package internal

type ResponseModel struct {
	StatusCode string
	Message    string
	Error      string
}

func SetResponse(statusCode, message, error string) ResponseModel {
	data := ResponseModel{
		StatusCode: statusCode,
		Message:    message,
		Error:      error,
	}

	return data
}
