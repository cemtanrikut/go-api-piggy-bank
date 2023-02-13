package internal

type ResponseModel struct {
	StatusCode string
	Message    string
	Error      string
	Data       interface{}
}

func SetResponse(statusCode, message, err string, data interface{}) ResponseModel {
	respData := ResponseModel{
		StatusCode: statusCode,
		Message:    message,
		Error:      err,
		Data:       data,
	}

	return respData
}
