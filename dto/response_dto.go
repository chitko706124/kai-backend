package dto

// swagger:model Response
type Response[T any] struct {
	// HTTP status code
	// example: 200
	Code int `json:"code"`
	// Response message
	// example: Success
	Message string `json:"message"`
	// Response data
	Data T `json:"data"`
}

func CreateResponse[T any](code int, message string, data T) Response[T] {
	return Response[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func CreateResponseWithoutData(code int, message string) Response[any] {
	return Response[any]{
		Code:    code,
		Message: message,
	}
}
