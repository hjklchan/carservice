package api

import (
	"fmt"
	"net/http"
)

type ApiCoder interface {
	HttpCode() int16
	Code() string
	Message() string
	WithDetails(...string)
	Details() []string
	Error() string
}

// apiCode is used to API error code.
type apiCode struct {
	httpCode int16
	code     string
	msg      string
	details  []string
}

// NewApiCode is create a error code entity.
func NewApiCode(httpCode int16, code, msg string) *apiCode {
	details := []string{}
	return &apiCode{httpCode, code, msg, details}
}

// HttpCode returns http status code.
func (ac *apiCode) HttpCode() int16 {
	return ac.httpCode
}

// Code returns error code.
func (ac *apiCode) Code() string {
	return "E" + ac.code
}

func (ac *apiCode) SetMessage(m string) *apiCode {
	ac.msg = m
	return ac
}

// Message returns error message.
func (ac *apiCode) Message() string {
	return ac.msg
}

// WithDetails could put more error details of type of string in response data.
func (ac *apiCode) WithDetails(dts ...string) *apiCode {
	ac.details = append(ac.details, dts...)
	return ac
}

// Details returns error details of type of string.
func (ac *apiCode) Details() []string {
	return ac.details
}

func (ac *apiCode) Error() string {
	return fmt.Sprintf("[%s] - %s", ac.Code(), ac.Message())
}

// Default error code.
var OK = NewApiCode(http.StatusOK, "0", "欧克")
