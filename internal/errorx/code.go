package errorx

import "net/http"

// 参数相关
var (
	ParamBindErr        = New(http.StatusBadRequest, "ParamBindErr", "Param Exception")
	ParamErr            = New(http.StatusBadRequest, "ParamErr", "Param Exception")
	ParamValidationErr  = New(http.StatusBadRequest, "ParamValidationErr", "Param Exception")
	ParamNotJsonRequest = New(http.StatusBadRequest, "ParamNotJsonRequest", "Param Exception")
	ParamHeaderErr      = New(http.StatusBadRequest, "ParamHeaderErr", "Param Exception")
)

var (
	DataSqlErr           = New(http.StatusInternalServerError, "DataSqlErr", "Data Exception")
	DataRedisErr         = New(http.StatusInternalServerError, "DataRedisErr", "Data Exception")
	DataRecordNotFound   = New(http.StatusInternalServerError, "DataRecordNotFound", "Data Exception")
	DataDuplicateRecords = New(http.StatusInternalServerError, "DataDuplicateRecords", "Data Exception")
	DataFormattingError  = New(http.StatusInternalServerError, "DataFormattingError", "Data Exception")
	DataProcessingError  = New(http.StatusInternalServerError, "DataProcessingError", "Data Exception")
)

var (
	AccountNotExist          = New(http.StatusUnprocessableEntity, "AccountNotExist", "Account Exception")
	AccountIsLocked          = New(http.StatusUnprocessableEntity, "AccountIsLocked", "Account Exception")
	AccountIsLogout          = New(http.StatusUnprocessableEntity, "AccountIsLogout", "Account Exception")
	AccountError             = New(http.StatusUnprocessableEntity, "AccountError", "Account Exception")
	AccountWrongPassword     = New(http.StatusUnprocessableEntity, "AccountWrongPassword", "Account Exception")
	AccountIsBanned          = New(http.StatusUnprocessableEntity, "AccountIsBanned", "Account Exception")
	AccountUpdateFailed      = New(http.StatusUnprocessableEntity, "AccountUpdateFailed", "Account Exception")
	AccountDuplicateUsername = New(http.StatusUnprocessableEntity, "AccountDuplicateUsername", "Account Exception")
	AccountAbnormalStatus    = New(http.StatusUnprocessableEntity, "AccountAbnormalStatus", "Account Exception")
	AccountNicknameViolation = New(http.StatusUnprocessableEntity, "AccountNicknameViolation", "Account Exception")
	AccountNotBoundRole      = New(http.StatusUnprocessableEntity, "AccountNotBoundRole", "Account Exception")
)

// 验证码

var (
	VerificationCodeErr = New(http.StatusUnprocessableEntity, "verification code err", "verification code Exception")
)
var (
	TokenNotRequest          = New(http.StatusUnprocessableEntity, "TokenNotRequest", "Token Exception")
	TokenFormatErr           = New(http.StatusUnprocessableEntity, "TokenFormatErr", "Token Exception")
	TokenErr                 = New(http.StatusUnprocessableEntity, "TokenErr", "Token Exception")
	TokenInvalidErr          = New(http.StatusUnprocessableEntity, "TokenInvalidErr", "Token Exception")
	TokenExpired             = New(http.StatusUnprocessableEntity, "TokenExpired", "Token Exception")
	TokenVerificationFailed  = New(http.StatusUnprocessableEntity, "TokenVerificationFailed", "Token Exception")
	TokenRefreshErr          = New(http.StatusUnprocessableEntity, "TokenRefreshErr", "Token Exception")
	TokenStorageFailed       = New(http.StatusUnprocessableEntity, "TokenStorageFailed", "Token Exception")
	TokenErrSignatureParam   = New(http.StatusUnprocessableEntity, "TokenErrSignatureParam", "Token Exception")
	TokenWrongTypeOfBusiness = New(http.StatusUnprocessableEntity, "TokenWrongTypeOfBusiness", "Token Exception")
	TokenGenerationFailed    = New(http.StatusUnprocessableEntity, "TokenGenerationFailed", "Token Exception")
)
