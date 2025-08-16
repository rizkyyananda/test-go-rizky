package helper

type Response struct {
	Status bool        `json:"status"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
	Info   interface{} `json:"info"`
}

type AppError struct {
	Code        int         `json:"code,omitempty"`
	Object      interface{} `json:"object"`
	Field       interface{} `json:"field"`
	MessageData interface{} `json:"messageData"`
}

func ResponseSuccess(data interface{}) Response {
	res := Response{
		Status: true,
		Data:   data,
		Error:  nil,
		Info:   nil,
	}
	return res
}

func ResponseError(message interface{}, code int) Response {
	return Response{
		Status: false,
		Data:   nil,
		Error: AppError{
			Code:        code,
			Object:      nil,
			Field:       nil,
			MessageData: message,
		},
		Info: nil,
	}
}
