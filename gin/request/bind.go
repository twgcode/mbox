/**
@Author: twgcode
@Email: 17600113577@163.com
@Date: 2022/10/20 19:29
@Description:
*/

package request

import (
	"fmt"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/twgcode/mbox/exception"
	"github.com/twgcode/mbox/gin/response"
)

func getFuncName() string {
	pc, _, _, _ := runtime.Caller(2)
	return runtime.FuncForPC(pc).Name()
}

func ShouldBind(c *gin.Context, obj interface{}, logger *zap.Logger) (err error) {
	funcName := getFuncName()
	if err = c.ShouldBind(obj); err != nil {
		// 请求参数有误，直接返回响应
		logger.Error(fmt.Sprintf("%s with invalid param", funcName), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Failed(c, exception.NewBadRequest(err.Error()))
			return
		}
		response.Failed(c, exception.NewInvalidParam(err.Error()))
		return
	}
	return
}

func ShouldBindWithHttpCode(c *gin.Context, obj interface{}, logger *zap.Logger) (err error) {
	funcName := getFuncName()
	if err = c.ShouldBind(obj); err != nil {
		// 请求参数有误，直接返回响应
		logger.Error(fmt.Sprintf("%s with invalid param", funcName), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			response.FailedWithHttpCode(c, exception.NewBadRequest(err.Error()))
			return
		}
		response.FailedWithHttpCode(c, exception.NewInvalidParam(err.Error()))
		return
	}
	return
}

func ShouldBindUri(c *gin.Context, obj interface{}, logger *zap.Logger) (err error) {
	funcName := getFuncName()
	if err = c.ShouldBindUri(obj); err != nil {
		// 请求参数有误，直接返回响应
		logger.Error(fmt.Sprintf("%s with invalid param", funcName), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			response.Failed(c, exception.NewBadRequest(err.Error()))
			return
		}
		response.Failed(c, exception.NewInvalidParam(err.Error()))
		return
	}
	return
}

func ShouldBindUriWithHttpCode(c *gin.Context, obj interface{}, logger *zap.Logger) (err error) {
	funcName := getFuncName()
	if err = c.ShouldBindUri(obj); err != nil {
		// 请求参数有误，直接返回响应
		logger.Error(fmt.Sprintf("%s with invalid param", funcName), zap.Error(err))
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			response.FailedWithHttpCode(c, exception.NewBadRequest(err.Error()))
			return
		}
		response.FailedWithHttpCode(c, exception.NewInvalidParam(err.Error()))
		return
	}
	return
}
