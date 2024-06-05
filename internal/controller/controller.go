package controller;

type Response struct {
	Success bool `json:"success"`
	Data interface{} `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func SuccessResponse (data interface{}) *Response {
	return &Response {
		Success: true,
		Data: data,
	}
}

func ErrorResponse (message string) *Response {
	return &Response {
		Success: false,
		Message: message,
	}
}
