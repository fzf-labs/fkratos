// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 未登录
func IsNotLoggedIn(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_NOT_LOGGED_IN.String() && e.Code == 401
}

// 未登录
func ErrorNotLoggedIn(format string, args ...interface{}) *errors.Error {
	return errors.New(401, SysErrorReason_NOT_LOGGED_IN.String(), fmt.Sprintf(format, args...))
}

// 访问被禁止的
func IsAccessForbidden(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_ACCESS_FORBIDDEN.String() && e.Code == 403
}

// 访问被禁止的
func ErrorAccessForbidden(format string, args ...interface{}) *errors.Error {
	return errors.New(403, SysErrorReason_ACCESS_FORBIDDEN.String(), fmt.Sprintf(format, args...))
}

// 找不到资源
func IsResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_RESOURCE_NOT_FOUND.String() && e.Code == 404
}

// 找不到资源
func ErrorResourceNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, SysErrorReason_RESOURCE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 方法不允许
func IsMethodNotAllowed(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_METHOD_NOT_ALLOWED.String() && e.Code == 405
}

// 方法不允许
func ErrorMethodNotAllowed(format string, args ...interface{}) *errors.Error {
	return errors.New(405, SysErrorReason_METHOD_NOT_ALLOWED.String(), fmt.Sprintf(format, args...))
}

// 请求超时
func IsRequestTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_REQUEST_TIMEOUT.String() && e.Code == 408
}

// 请求超时
func ErrorRequestTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(408, SysErrorReason_REQUEST_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 服务器内部错误
func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_INTERNAL_SERVER_ERROR.String() && e.Code == 500
}

// 服务器内部错误
func ErrorInternalServerError(format string, args ...interface{}) *errors.Error {
	return errors.New(500, SysErrorReason_INTERNAL_SERVER_ERROR.String(), fmt.Sprintf(format, args...))
}

// 不可执行
func IsNotImplemented(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_NOT_IMPLEMENTED.String() && e.Code == 501
}

// 不可执行
func ErrorNotImplemented(format string, args ...interface{}) *errors.Error {
	return errors.New(501, SysErrorReason_NOT_IMPLEMENTED.String(), fmt.Sprintf(format, args...))
}

// 网络错误
func IsNetworkError(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_NETWORK_ERROR.String() && e.Code == 502
}

// 网络错误
func ErrorNetworkError(format string, args ...interface{}) *errors.Error {
	return errors.New(502, SysErrorReason_NETWORK_ERROR.String(), fmt.Sprintf(format, args...))
}

// 服务不可用
func IsServiceUnavailable(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_SERVICE_UNAVAILABLE.String() && e.Code == 503
}

// 服务不可用
func ErrorServiceUnavailable(format string, args ...interface{}) *errors.Error {
	return errors.New(503, SysErrorReason_SERVICE_UNAVAILABLE.String(), fmt.Sprintf(format, args...))
}

// 网络超时
func IsNetworkTimeout(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_NETWORK_TIMEOUT.String() && e.Code == 504
}

// 网络超时
func ErrorNetworkTimeout(format string, args ...interface{}) *errors.Error {
	return errors.New(504, SysErrorReason_NETWORK_TIMEOUT.String(), fmt.Sprintf(format, args...))
}

// 请求不支持
func IsRequestNotSupport(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_REQUEST_NOT_SUPPORT.String() && e.Code == 505
}

// 请求不支持
func ErrorRequestNotSupport(format string, args ...interface{}) *errors.Error {
	return errors.New(505, SysErrorReason_REQUEST_NOT_SUPPORT.String(), fmt.Sprintf(format, args...))
}

// 未找到用户
func IsUserNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_USER_NOT_FOUND.String() && e.Code == 600
}

// 未找到用户
func ErrorUserNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(600, SysErrorReason_USER_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 密码错误
func IsIncorrectPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_INCORRECT_PASSWORD.String() && e.Code == 599
}

// 密码错误
func ErrorIncorrectPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(599, SysErrorReason_INCORRECT_PASSWORD.String(), fmt.Sprintf(format, args...))
}

// 用户冻结
func IsUserFreeze(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_USER_FREEZE.String() && e.Code == 598
}

// 用户冻结
func ErrorUserFreeze(format string, args ...interface{}) *errors.Error {
	return errors.New(598, SysErrorReason_USER_FREEZE.String(), fmt.Sprintf(format, args...))
}

// 用户ID无效
func IsInvalidUserid(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_INVALID_USERID.String() && e.Code == 101
}

// 用户ID无效
func ErrorInvalidUserid(format string, args ...interface{}) *errors.Error {
	return errors.New(101, SysErrorReason_INVALID_USERID.String(), fmt.Sprintf(format, args...))
}

// 密码无效
func IsInvalidPassword(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_INVALID_PASSWORD.String() && e.Code == 102
}

// 密码无效
func ErrorInvalidPassword(format string, args ...interface{}) *errors.Error {
	return errors.New(102, SysErrorReason_INVALID_PASSWORD.String(), fmt.Sprintf(format, args...))
}

// token过期
func IsTokenExpired(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_TOKEN_EXPIRED.String() && e.Code == 103
}

// token过期
func ErrorTokenExpired(format string, args ...interface{}) *errors.Error {
	return errors.New(103, SysErrorReason_TOKEN_EXPIRED.String(), fmt.Sprintf(format, args...))
}

// token无效
func IsInvalidToken(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_INVALID_TOKEN.String() && e.Code == 104
}

// token无效
func ErrorInvalidToken(format string, args ...interface{}) *errors.Error {
	return errors.New(104, SysErrorReason_INVALID_TOKEN.String(), fmt.Sprintf(format, args...))
}

// token不存在
func IsTokenNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_TOKEN_NOT_EXIST.String() && e.Code == 105
}

// token不存在
func ErrorTokenNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(105, SysErrorReason_TOKEN_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 用户不存在
func IsUserNotExist(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_USER_NOT_EXIST.String() && e.Code == 106
}

// 用户不存在
func ErrorUserNotExist(format string, args ...interface{}) *errors.Error {
	return errors.New(106, SysErrorReason_USER_NOT_EXIST.String(), fmt.Sprintf(format, args...))
}

// 验证码错误
func IsIncorrectVerificationCode(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == SysErrorReason_INCORRECT_VERIFICATION_CODE.String() && e.Code == 107
}

// 验证码错误
func ErrorIncorrectVerificationCode(format string, args ...interface{}) *errors.Error {
	return errors.New(107, SysErrorReason_INCORRECT_VERIFICATION_CODE.String(), fmt.Sprintf(format, args...))
}