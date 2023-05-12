package errorx

import (
	"github.com/go-kratos/kratos/v2/errors"
)

// 参数相关
var (
	ParamBindErr        = errors.New(20001, "Param Bind Err", "Param Exception")
	ParamErr            = errors.New(20002, "Param Err", "Param Exception")
	ParamValidationErr  = errors.New(20003, "Param Validation Err", "Param Exception")
	ParamNotJsonRequest = errors.New(20004, "Param Not Json Request", "Param Exception")
	ParamHeaderErr      = errors.New(20005, "Param Header Err", "Param Exception")
)

var (
	DataSqlErr           = errors.New(20100, "Data Sql Err", "Data Exception")
	DataRedisErr         = errors.New(20101, "Data Redis Err", "Data Exception")
	DataRecordNotFound   = errors.New(20102, "Data Record Not Found", "Data Exception")
	DataDuplicateRecords = errors.New(20103, "Data Duplicate Records", "Data Exception")
	DataFormattingError  = errors.New(20104, "Data Formatting Error", "Data Exception")
	DataProcessingError  = errors.New(20105, "Data Processing Error", "Data Exception")
)

var (
	AccountNotExist          = errors.New(21000, "Account Not Exist", "Account Exception")
	AccountIsLocked          = errors.New(21002, "Account Locked", "Account Exception")
	AccountIsLogout          = errors.New(21003, "Account Is Logout", "Account Exception")
	AccountError             = errors.New(21004, "Account Error", "Account Exception")
	AccountWrongPassword     = errors.New(21005, "Account Wrong Password", "Account Exception")
	AccountIsBanned          = errors.New(21006, "Account Is Banned", "Account Exception")
	AccountUpdateFailed      = errors.New(21007, "Account Update Failed", "Account Exception")
	AccountDuplicateUsername = errors.New(21008, "Account Duplicate Username", "Account Exception")
	AccountAbnormalStatus    = errors.New(21009, "Account Abnormal Status", "Account Exception")
	AccountNicknameViolation = errors.New(21010, "Account Nickname Violation", "Account Exception")
	AccountNotBoundRole      = errors.New(21011, "Account Not Bound Role", "Account Exception")
)

var (
	TokenNotRequest          = errors.New(20200, "TokenNotRequest", "Token Exception")
	TokenFormatErr           = errors.New(20201, "TokenFormatErr", "Token Exception")
	TokenErr                 = errors.New(20202, "TokenErr", "Token Exception")
	TokenInvalidErr          = errors.New(20203, "TokenInvalidErr", "Token Exception")
	TokenExpired             = errors.New(20205, "TokenExpired", "Token Exception")
	TokenVerificationFailed  = errors.New(20206, "TokenVerificationFailed", "Token Exception")
	TokenRefreshErr          = errors.New(20207, "TokenRefreshErr", "Token Exception")
	TokenStorageFailed       = errors.New(20208, "TokenStorageFailed", "Token Exception")
	TokenErrSignatureParam   = errors.New(20209, "TokenErrSignatureParam", "Token Exception")
	TokenWrongTypeOfBusiness = errors.New(20210, "TokenWrongTypeOfBusiness", "Token Exception")
	TokenGenerationFailed    = errors.New(21011, "TokenGenerationFailed", "Token Exception")
)
