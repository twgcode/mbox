package response

import (
	"fmt"
	"net/http"

	"github.com/twgcode/mbox/exception"
)

const (
	CodeSuccess   int64  = 0
	SuccessReason string = "success"
)

func NewData(code exception.Code, reason, msg string, data interface{}, opts ...Option) *Data {
	resp := &Data{
		Code:   code,
		Reason: reason,
		Msg:    msg,
		Data:   data,
	}
	for _, opt := range opts {
		opt.apply(resp)
	}
	return resp
}

// Data to be used by controllers.
type Data struct {
	Code     exception.Code `json:"code"`   // 自定义返回码  0:表示正常
	Reason   string         `json:"reason"` // 异常原因, 简单信息
	Msg      string         `json:"msg"`    // 关于这次响应的说明信息
	Data     interface{}    `json:"data"`   // 返回的具体数据
	httpCode *int           // 不用对外, json序列化后不能出现该字段
}

func (d *Data) Error() error {
	if d.Code == exception.Success {
		return nil
	}
	return fmt.Errorf("code: %d message: %s", d.Code, d.Msg)
}
func (d *Data) HttpCode() int {
	if d.httpCode == nil {
		return http.StatusOK
	}
	return *d.httpCode
}

// Option configures how we set up the data.
type Option interface {
	apply(*Data)
}

func newFuncOption(f func(*Data)) Option {
	return &funcOption{
		f: f,
	}
}

type funcOption struct {
	f func(*Data)
}

func (fdo *funcOption) apply(do *Data) {
	fdo.f(do)
}

func WithHttpCode(httpCode int) Option {
	return newFuncOption(func(o *Data) {
		o.httpCode = &httpCode
	})
}

// WithJoinReasonMsg 合并  Reason 和 Msg 到 Msg
func WithJoinReasonMsg() Option {
	return newFuncOption(func(o *Data) {
		var msg string
		if o.Reason != "" {
			msg = o.Msg
		}
		if o.Msg != "" {
			if msg == "" {
				msg = o.Msg
			} else {
				msg = fmt.Sprintf("%s, %s", msg, o.Msg)
			}
		}
		o.Msg = msg
	})
}

// WithInternalServerError 遇到 exception.InternalServerError 时 把 Msg 替换为 Reason, 避免对外暴露太多的错误细节
func WithInternalServerError() Option {
	return newFuncOption(func(o *Data) {
		if o.Code.Is(exception.InternalServerError) {
			o.Msg = o.Reason
		}
	})
}

// WithUnKnownException 遇到 exception.UnKnownException 时 把 Msg 替换为 Reason, 避免对外暴露太多的错误细节
func WithUnKnownException() Option {
	return newFuncOption(func(o *Data) {
		if o.Code.Is(exception.UnKnownException) {
			o.Msg = o.Reason
		}
	})
}
