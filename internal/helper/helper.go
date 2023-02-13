package internal

type ResponseModel struct {
	StatusCode string
	Message    string
	Error      string
	Data       interface{}
}

func SetResponse(statusCode, message, error string, data interface{}) ResponseModel {
	respData := ResponseModel{
		StatusCode: statusCode,
		Message:    message,
		Error:      error,
		Data:       data,
	}

	return respData
}
