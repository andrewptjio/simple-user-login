package util

//DefaultResponse default payload response
type DefaultResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ResponseSuccess struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ResponseError struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

type ResponsePagination struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Page    int         `json:"page"`
	PerPage int         `json:"per_page"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(data interface{}) ResponseSuccess {
	return ResponseSuccess{
		Code:    200,
		Message: "Successful Operation",
		Data:    data,
	}
}

func ErrorResponse(code int, err interface{}) ResponseError {
	return ResponseError{
		Code:    code,
		Message: err,
	}
}

//NewInternalServerErrorResponse default internal server error response
func NewSuccessOperationResponse() DefaultResponse {
	return DefaultResponse{
		200,
		"Successful Operation",
	}
}

//NewInternalServerErrorResponse default internal server error response
func NewInternalServerErrorResponse() DefaultResponse {
	return DefaultResponse{
		500,
		"Internal Server Error",
	}
}

//NewNotFoundResponse default not found error response
func NewNotFoundResponse() DefaultResponse {
	return DefaultResponse{
		404,
		"Not Found",
	}
}

//NewBadRequestResponse default bad request error response
func NewBadRequestResponse() DefaultResponse {
	return DefaultResponse{
		400,
		"Bad Request",
	}
}

//NewConflictResponse default conflict response error response
func NewConflictResponse() DefaultResponse {
	return DefaultResponse{
		409,
		"Data Has Been Modified",
	}
}

//NewStatusNotAccepted default not
func NewStatusNotAcceptable() DefaultResponse {
	return DefaultResponse{
		406,
		"Not Accepted",
	}
}

//NewStatusForbidden default not
func NewStatusForbidden() DefaultResponse {
	return DefaultResponse{
		403,
		"Forbidden",
	}
}
