package rest

// Result is a generic json result type
type Result struct {
	Code    int    `json:"code"`
	State   bool   `json:"state"`
	Message string `json:"message"`
	Error   string `json:"error"`
	Data    any    `json:"data"`
}

// Success return success result
func Success() *Result {
	return &Result{
		Code:    0,
		State:   true,
		Message: "",
		Error:   "",
		Data:    nil,
	}
}

// SuccessMessage return success result with message
func SuccessMessage(message string) *Result {
	return &Result{
		Code:    0,
		State:   true,
		Message: message,
		Error:   "",
		Data:    nil,
	}
}

// Error return error result
func Error(code int) *Result {
	return &Result{
		Code:    code,
		State:   false,
		Message: "",
		Error:   "",
		Data:    nil,
	}
}

// ErrorMessage return error result with error message
func ErrorMessage(code int, error string) *Result {
	return &Result{
		Code:    code,
		State:   false,
		Message: "",
		Error:   error,
		Data:    nil,
	}
}

// Fail return fail result
func Fail() *Result {
	return &Result{
		Code:    0,
		State:   false,
		Message: "",
		Error:   "",
		Data:    nil,
	}
}

// FailMessage return fail result with error message
func FailMessage(error string) *Result {
	return &Result{
		Code:    0,
		State:   false,
		Message: "",
		Error:   error,
		Data:    nil,
	}
}

// Data return success result with data result
func Data(data any) *Result {
	return &Result{
		Code:    0,
		State:   true,
		Message: "",
		Error:   "",
		Data:    data,
	}
}
