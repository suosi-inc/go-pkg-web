package rest

// Result is a generic json result type
type Result struct {
	Code    int     `json:"code"`
	State   bool    `json:"state"`
	Message *string `json:"message,omitempty"`
	Error   *string `json:"error,omitempty"`
	Data    any     `json:"data,omitempty"`
}

// Success return success result
func Success() *Result {
	return &Result{
		Code:    0,
		State:   true,
		Message: nil,
		Error:   nil,
		Data:    nil,
	}
}

// SuccessMessage return success result with message
func SuccessMessage(message string) *Result {
	return &Result{
		Code:    0,
		State:   true,
		Message: &message,
		Error:   nil,
		Data:    nil,
	}
}

// Error return error result
func Error(code int) *Result {
	return &Result{
		Code:    code,
		State:   false,
		Message: nil,
		Error:   nil,
		Data:    nil,
	}
}

// ErrorMessage return error result with error message
func ErrorMessage(code int, error string) *Result {
	return &Result{
		Code:    code,
		State:   false,
		Message: nil,
		Error:   &error,
		Data:    nil,
	}
}

// Fail return fail result
func Fail() *Result {
	return &Result{
		Code:    0,
		State:   false,
		Message: nil,
		Error:   nil,
		Data:    nil,
	}
}

// FailMessage return fail result with error message
func FailMessage(error string) *Result {
	return &Result{
		Code:    0,
		State:   false,
		Message: nil,
		Error:   &error,
		Data:    nil,
	}
}

// Data return success result with data result
func Data(data any) *Result {
	return &Result{
		Code:    0,
		State:   true,
		Message: nil,
		Error:   nil,
		Data:    data,
	}
}

// DataMessage return success result with data result and message
func DataMessage(data any, message string) *Result {
	return &Result{
		Code:    0,
		State:   true,
		Message: &message,
		Error:   nil,
		Data:    data,
	}
}
