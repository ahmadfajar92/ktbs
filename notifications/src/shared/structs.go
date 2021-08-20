package shared

import "net/http"

type (
	// Result struct
	Result struct {
		Success bool        `json:"success"`
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Error   error       `json:"error"`
		Data    interface{} `json:"data"`
	}
)

func FailedResult(s string, e error) *Result {
	return &Result{
		Success: false,
		Code:    http.StatusBadRequest,
		Error:   e,
		Message: s,
	}
}

func SuccessResult(d interface{}) *Result {
	return &Result{
		Success: true,
		Code:    http.StatusOK,
		Message: "Success",
		Data:    d,
	}
}
