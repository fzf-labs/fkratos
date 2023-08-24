package errorx

import "net/http"

var (
	InternalServerError = New(http.StatusInternalServerError, "InternalServerError", "服务崩溃了,请稍后再试")
)

// 参数相关
var (
	ParamBindErr        = New(http.StatusBadRequest, "ParamBindErr", "Param Exception")
	ParamErr            = New(http.StatusBadRequest, "ParamErr", "Param Exception")
	ParamValidationErr  = New(http.StatusBadRequest, "ParamValidationErr", "Param Exception")
	ParamNotJSONRequest = New(http.StatusBadRequest, "ParamNotJSONRequest", "Param Exception")
	ParamHeaderErr      = New(http.StatusBadRequest, "ParamHeaderErr", "Param Exception")
)

// Data数据相关
var (
	DataSQLErr           = New(http.StatusInternalServerError, "DataSQLErr", "Data Exception")
	DataRedisErr         = New(http.StatusInternalServerError, "DataRedisErr", "Data Exception")
	DataRecordNotFound   = New(http.StatusInternalServerError, "DataRecordNotFound", "Data Exception")
	DataDuplicateRecords = New(http.StatusInternalServerError, "DataDuplicateRecords", "Data Exception")
	DataFormattingError  = New(http.StatusInternalServerError, "DataFormattingError", "Data Exception")
	DataProcessingError  = New(http.StatusInternalServerError, "DataProcessingError", "Data Exception")
)

// Account账号验证
var (
	AccountNotExist            = New(http.StatusConflict, "AccountNotExist", "Account Exception")
	AccountIsLocked            = New(http.StatusConflict, "AccountIsLocked", "Account Exception")
	AccountIsLogout            = New(http.StatusConflict, "AccountIsLogout", "Account Exception")
	AccountError               = New(http.StatusConflict, "AccountError", "Account Exception")
	AccountWrongPassword       = New(http.StatusConflict, "AccountWrongPassword", "Account Exception")
	AccountIsBanned            = New(http.StatusConflict, "AccountIsBanned", "Account Exception")
	AccountUpdateFailed        = New(http.StatusConflict, "AccountUpdateFailed", "Account Exception")
	AccountDuplicateUsername   = New(http.StatusConflict, "AccountDuplicateUsername", "Account Exception")
	AccountAbnormalStatus      = New(http.StatusConflict, "AccountAbnormalStatus", "Account Exception")
	AccountNicknameViolation   = New(http.StatusConflict, "AccountNicknameViolation", "Account Exception")
	AccountNotBoundRole        = New(http.StatusConflict, "AccountNotBoundRole", "Account Exception")
	AccountVerificationCodeErr = New(http.StatusConflict, "AccountVerificationCodeErr", "Account Exception")
)

// Token验证
var (
	TokenNotRequest          = New(http.StatusConflict, "TokenNotRequest", "Token Exception")
	TokenFormatErr           = New(http.StatusConflict, "TokenFormatErr", "Token Exception")
	TokenErr                 = New(http.StatusConflict, "TokenErr", "Token Exception")
	TokenInvalidErr          = New(http.StatusConflict, "TokenInvalidErr", "Token Exception")
	TokenExpired             = New(http.StatusConflict, "TokenExpired", "Token Exception")
	TokenVerificationFailed  = New(http.StatusConflict, "TokenVerificationFailed", "Token Exception")
	TokenRefreshErr          = New(http.StatusConflict, "TokenRefreshErr", "Token Exception")
	TokenStorageFailed       = New(http.StatusConflict, "TokenStorageFailed", "Token Exception")
	TokenErrSignatureParam   = New(http.StatusConflict, "TokenErrSignatureParam", "Token Exception")
	TokenWrongTypeOfBusiness = New(http.StatusConflict, "TokenWrongTypeOfBusiness", "Token Exception")
	TokenGenerationFailed    = New(http.StatusConflict, "TokenGenerationFailed", "Token Exception")
)
