package libraries

type ResponseSuccess struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}
type ResponseError struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
}

func StatusOk(data interface{}) interface{} {
	return &ResponseSuccess{
		Code:   200,
		Status: "OK",
		Data:   data,
	}
}
func StatusCreated(data interface{}) interface{} {
	return &ResponseSuccess{
		Code:   201,
		Status: "Created",
		Data:   data,
	}
}
func StatusBadRequest(errors interface{}) interface{} {
	return &ResponseError{
		Code:   400,
		Status: "Bad Request",
		Errors: errors,
	}
}
func StatusNotFound(errors interface{}) interface{} {
	return &ResponseError{
		Code:   404,
		Status: "Not Found",
		Errors: errors,
	}
}
func StatusInternalServerError(errors interface{}) interface{} {
	return &ResponseError{
		Code:   500,
		Status: "Internal Server Error",
		Errors: errors,
	}
}
func StatusNoContent(data interface{}) interface{} {
	return &ResponseSuccess{
		Code:   204,
		Status: "No Content",
		Data:   data,
	}
}
