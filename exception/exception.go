package exception

// NewAPIException 创建一个API异常
// 用于其他模块自定义异常
func NewAPIException(code Code, format string, a ...interface{}) APIException {
	// 0表示正常状态, 但是要排除变量的零值
	if code == Success {
		code = UnKnownException
	}
	return newException(code, format, a...)
}

// NewUnKnownException 未知异常
func NewUnKnownException(format string, a ...interface{}) APIException {
	return newException(UnKnownException, format, a...)
}

// NewUnauthorized 未认证
func NewUnauthorized(format string, a ...interface{}) APIException {
	return newException(Unauthorized, format, a...)
}

// NewPermissionDeny 没有权限访问
func NewPermissionDeny(format string, a ...interface{}) APIException {
	return newException(Forbidden, format, a...)
}

// NewBadRequest 请求不合法
func NewBadRequest(format string, a ...interface{}) APIException {
	return newException(BadRequest, format, a...)
}

// NewNotFound 资源找不到
func NewNotFound(format string, a ...interface{}) APIException {
	return newException(NotFound, format, a...)
}

// NewMethodNotAllowed 405
func NewMethodNotAllowed(format string, a ...interface{}) APIException {
	return newException(MethodNotAllowed, format, a...)
}

// NewConflict 资源冲突, 已经存在
func NewConflict(format string, a ...interface{}) APIException {
	return newException(Conflict, format, a...)
}

// NewInternalServerError 500
func NewInternalServerError(format string, a ...interface{}) APIException {
	return newException(InternalServerError, format, a...)
}

// NewAccessTokenExpired 访问token过期
func NewAccessTokenExpired(format string, a ...interface{}) APIException {
	return newException(AccessTokenExpired, format, a...)
}

// NewRefreshTokenExpired 刷新token过期
func NewRefreshTokenExpired(format string, a ...interface{}) APIException {
	return newException(RefreshTokenExpired, format, a...)
}

// NewAccessTokenIllegal 访问token不合法
func NewAccessTokenIllegal(format string, a ...interface{}) APIException {
	return newException(AccessTokenIllegal, format, a...)
}

// NewRefreshTokenIllegal 刷新token不合法
func NewRefreshTokenIllegal(format string, a ...interface{}) APIException {
	return newException(RefreshTokenIllegal, format, a...)
}

// NewInvalidParam 无效参数, 参数不符合要求
func NewInvalidParam(format string, a ...interface{}) APIException {
	return newException(InvalidParam, format, a...)
}

// NewIllegalOperation 非法操作,
func NewIllegalOperation(format string, a ...interface{}) APIException {
	return newException(IllegalOperation, format, a...)
}

func FromError(err error) (a APIException, ok bool) {
	if err == nil {
		return nil, false
	}
	a, ok = err.(APIException)
	return
}

// IsBadRequestError 判断是否是 BadRequest
func IsBadRequestError(err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(APIException)
	if !ok {
		return false
	}

	return e.ErrorCode() == BadRequest
}

// IsNotFoundError 判断是否是NotFoundError
func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(APIException)
	if !ok {
		return false
	}

	return e.ErrorCode() == NotFound
}

// IsConflictError 判断是否是Conflict
func IsConflictError(err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(APIException)
	if !ok {
		return false
	}

	return e.ErrorCode() == Conflict
}

// IsPermissionDenyError 判断是否是 Forbidden
func IsPermissionDenyError(err error) bool {
	if err == nil {
		return false
	}

	e, ok := err.(APIException)
	if !ok {
		return false
	}

	return e.ErrorCode() == Forbidden
}
