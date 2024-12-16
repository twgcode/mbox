package exception

import "net/http"

type Code int64

func (c Code) Is(code Code) bool {
	return c == code
}

const (
	UnKnownException Code = -1 // 未知异常
	Success          Code = 0  // 成功, 无异常
)

const (
	// Unauthorized 未认证
	Unauthorized = http.StatusUnauthorized
	// BadRequest 请求不合法
	BadRequest = http.StatusBadRequest
	// Forbidden 无权限
	Forbidden = http.StatusForbidden
	// NotFound 资源未找到
	NotFound = http.StatusNotFound
	// MethodNotAllowed 方法不允许
	MethodNotAllowed = http.StatusMethodNotAllowed
	// Conflict 资源冲突, 已经存在
	Conflict = http.StatusConflict
	// UnprocessableEntity 语义错误 ; 适用于值在逻辑或业务上不合法的情况(比如用于表示参数通过了语法检查，但语义上不合法)。
	UnprocessableEntity = http.StatusUnprocessableEntity
	InternalServerError = http.StatusInternalServerError

	// AccessTokenExpired token过期
	AccessTokenExpired Code = 1002
	// RefreshTokenExpired token过期
	RefreshTokenExpired Code = 1003
	// AccessTokenIllegal 访问token不合法
	AccessTokenIllegal Code = 1004
	// RefreshTokenIllegal 刷新token不合法
	RefreshTokenIllegal Code = 1005
)
const (
	// InvalidParam 无效的参数, 参数不符合要求
	InvalidParam Code = 2000
	// IllegalOperation 非法操作
	IllegalOperation Code = 2001
)

var (
	reasonMap = map[Code]string{
		UnKnownException:    "未知异常",
		InternalServerError: "系统内部错误",
		Success:             "success",

		Unauthorized:        "认证失败",
		Forbidden:           "访问未授权",
		NotFound:            "Not Found",
		MethodNotAllowed:    "Method Not Allowed",
		Conflict:            "资源已经存在",
		BadRequest:          "请求不合法",
		UnprocessableEntity: "参数语义错误",

		AccessTokenIllegal:  "访问令牌不合法",
		RefreshTokenIllegal: "刷新令牌不合法",
		AccessTokenExpired:  "访问过期, 请刷新",
		RefreshTokenExpired: "刷新过期, 请登录",

		InvalidParam:     "无效的参数",
		IllegalOperation: "非法操作",
	}
)

func CodeReason(code Code) string {
	v, ok := reasonMap[code]
	if !ok {
		v = reasonMap[UnKnownException]
	}
	return v
}

func SetCodeReasonMap(code Code, reason string) {
	reasonMap[code] = reason
}

func CodeToHttpCode(code Code) (httpCode int) {
	if code == UnKnownException {
		return http.StatusInternalServerError
	}
	if code == Success {
		return http.StatusOK
	}
	if code == InvalidParam || code == IllegalOperation {
		return BadRequest
	}
	// token 有关的
	if code >= AccessTokenExpired && code <= RefreshTokenIllegal {
		return http.StatusUnauthorized
	}

	if code >= 100 && code <= 599 {
		return int(code)
	}

	// 最后兜底
	return http.StatusInternalServerError
}

func ErrToHttpCode(err error) (httpCode int) {
	var (
		a  APIException
		ok bool
	)
	if err == nil {
		return http.StatusOK
	}
	a, ok = FromError(err)
	if !ok {
		return http.StatusInternalServerError
	}

	return CodeToHttpCode(a.ErrorCode())
}
