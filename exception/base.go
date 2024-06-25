package exception

import (
	"fmt"
)

// APIException API异常
type APIException interface {
	error
	ErrorCode() Code
	WithMeta(m interface{}) APIException
	Meta() interface{}
	WithData(d interface{}) APIException
	Data() interface{}
	IsCode(code Code) bool
	Reason() string
}

func newException(code Code, format string, a ...interface{}) *exception {
	return &exception{
		code:    code,
		reason:  CodeReason(code),
		message: fmt.Sprintf(format, a...),
	}
}

// APIException is implementation for api exception
type exception struct {
	code    Code
	reason  string
	message string
	meta    interface{}
	data    interface{}
}

func (e *exception) Error() string {
	if e == nil {
		return ""
	}
	return e.message
}

// ErrorCode Code exception's code, 如果code不存在返回-1
func (e *exception) ErrorCode() Code {
	if e == nil {
		return Success
	}
	return e.code
}

// WithMeta 携带一些额外信息
func (e *exception) WithMeta(m interface{}) APIException {
	if e == nil {
		return nil
	}
	e.meta = m
	return e
}

func (e *exception) Meta() interface{} {
	if e == nil {
		return nil
	}
	return e.meta
}

func (e *exception) WithData(d interface{}) APIException {
	if e == nil {
		return nil
	}
	e.data = d
	return e
}

func (e *exception) Data() interface{} {
	if e == nil {
		return nil
	}
	return e.data
}

func (e *exception) IsCode(code Code) bool {
	if e == nil {
		return Success.Is(code)
	}
	return e.code.Is(code)
}

func (e *exception) Reason() string {
	if e == nil {
		return ""
	}
	return e.reason
}
