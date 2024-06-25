package response

import (
	"github.com/gin-gonic/gin"

	"github.com/twgcode/mbox/exception"
)

// FailedWithHttpCode  和 Failed 基本一致 只不过默认加了 WithHttpCode, 当然用户也可以添加新的 WithHttpCode 用来覆盖默认的 WithHttpCode
func FailedWithHttpCode(c *gin.Context, err error, opts ...Option) {
	var (
		errCode exception.Code
		reason  string
		data    interface{}
		msg     string
	)

	switch t := err.(type) {
	case exception.APIException:
		errCode = t.ErrorCode()
		reason = t.Reason()
		data = t.Data()
	default:
		errCode = exception.UnKnownException
		reason = exception.CodeReason(exception.UnKnownException)
	}

	// 获取 msg信息
	if err == nil { // 避免因为err 为nil时 直接调用  err.Error() 从而导致程序 panic
		msg = exception.CodeReason(exception.UnKnownException)
	} else {
		msg = err.Error()
	}
	// 获取 http code
	httpCode := exception.CodeToHttpCode(errCode)

	// 加入 默认的 WithHttpCode
	newOpts := make([]Option, 0, len(opts)+1)
	newOpts = append(newOpts, WithHttpCode(httpCode))
	opts = append(newOpts, opts...)

	resp := NewData(errCode, reason, msg, data, newOpts...)

	c.JSON(resp.HttpCode(), resp)
}

// Failed use to response error message
func Failed(c *gin.Context, err error, opts ...Option) {
	var (
		errCode exception.Code
		reason  string
		data    interface{}
		msg     string
	)

	switch t := err.(type) {
	case exception.APIException:
		errCode = t.ErrorCode()
		reason = t.Reason()
		data = t.Data()
	default:
		errCode = exception.UnKnownException
		reason = exception.CodeReason(exception.UnKnownException)
	}

	// 获取 msg信息
	if err == nil { // 避免因为err 为nil时 直接调用  err.Error() 从而导致程序 panic
		msg = exception.CodeReason(exception.UnKnownException)
	} else {
		msg = err.Error()
	}

	resp := NewData(errCode, reason, msg, data, opts...)

	c.JSON(resp.HttpCode(), resp)
}

// FailedWithNoErrorDetails 不对外暴露错误细节 不暴露错误细节， 目前是强制执行  WithInternalServerError() WithUnKnownException()
func FailedWithNoErrorDetails(c *gin.Context, err error, opts ...Option) {
	_opts := make([]Option, 0, len(opts)+2)
	_opts = append(_opts, opts...)
	_opts = append(_opts, WithUnKnownException(), WithInternalServerError())
	Failed(c, err, _opts...)
}

// Success use to response success data
func Success(c *gin.Context, data interface{}, opts ...Option) {
	resp := NewData(exception.Success, exception.CodeReason(exception.Success), "", data, opts...)
	c.JSON(resp.HttpCode(), resp)
}
