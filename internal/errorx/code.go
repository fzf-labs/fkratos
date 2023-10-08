package errorx

import "net/http"

var (
	InternalServerError = New(http.StatusInternalServerError, "InternalServerError", "服务崩溃了,请稍后再试")
	RequestCanceledErr  = New(http.StatusConflict, "RequestCanceledErr", "请求被取消")
)

// 参数相关
var (
	ParamBindErr        = New(http.StatusBadRequest, "ParamBindErr", "参数绑定错误")
	ParamErr            = New(http.StatusBadRequest, "ParamErr", "参数错误")
	ParamValidationErr  = New(http.StatusBadRequest, "ParamValidationErr", "参数验证错误")
	ParamNotJSONRequest = New(http.StatusBadRequest, "ParamNotJSONRequest", "参数不是JSON请求")
	ParamHeaderErr      = New(http.StatusBadRequest, "ParamHeaderErr", "参数Header错误")
)

// Data数据相关
var (
	DataSQLErr           = New(http.StatusInternalServerError, "DataSQLErr", "数据处理异常(S),请稍后再试")
	DataRedisErr         = New(http.StatusInternalServerError, "DataRedisErr", "数据处理异常(R),请稍后再试")
	DataMQErr            = New(http.StatusInternalServerError, "DataMQErr", "数据处理异常(M),请稍后再试")
	DataRecordNotFound   = New(http.StatusInternalServerError, "DataRecordNotFound", "数据记录未找到")
	DataDuplicateRecords = New(http.StatusInternalServerError, "DataDuplicateRecords", "数据重复记录")
	DataFormattingError  = New(http.StatusInternalServerError, "DataFormattingError", "数据格式化错误")
	DataProcessingError  = New(http.StatusInternalServerError, "DataProcessingError", "数据处理错误")
)

// API第三方请求
var (
	APIInternalErr = New(http.StatusConflict, "APIInternalErr", "请求错误(I),请稍后再试")
	APIThirdErr    = New(http.StatusConflict, "APIThirdErr", "请求错误(T),请稍后再试")
)

// Account账号验证
var (
	AccountNotExist            = New(http.StatusConflict, "AccountNotExist", "帐户不存在")
	AccountIsLocked            = New(http.StatusConflict, "AccountIsLocked", "帐号被锁定")
	AccountIsLogout            = New(http.StatusConflict, "AccountIsLogout", "帐户已注销")
	AccountError               = New(http.StatusConflict, "AccountError", "帐号错误")
	AccountWrongPassword       = New(http.StatusConflict, "AccountWrongPassword", "帐号密码错误")
	AccountIsBanned            = New(http.StatusConflict, "AccountIsBanned", "帐号被封禁")
	AccountUpdateFailed        = New(http.StatusConflict, "AccountUpdateFailed", "帐号更新失败")
	AccountDuplicateUsername   = New(http.StatusConflict, "AccountDuplicateUsername", "重复帐户")
	AccountAbnormalStatus      = New(http.StatusConflict, "AccountAbnormalStatus", "帐户异常状态")
	AccountNicknameViolation   = New(http.StatusConflict, "AccountNicknameViolation", "帐户昵称违规")
	AccountNotBoundRole        = New(http.StatusConflict, "AccountNotBoundRole", "帐号未绑定角色")
	AccountVerificationCodeErr = New(http.StatusConflict, "AccountVerificationCodeErr", "帐户验证码错误")
)

// Token验证
var (
	TokenNotRequest          = New(http.StatusConflict, "TokenNotRequest", "令牌不存在")
	TokenFormatErr           = New(http.StatusConflict, "TokenFormatErr", "令牌格式错误")
	TokenErr                 = New(http.StatusConflict, "TokenErr", "令牌错误")
	TokenInvalidErr          = New(http.StatusConflict, "TokenInvalidErr", "令牌校验失败")
	TokenExpired             = New(http.StatusConflict, "TokenExpired", "令牌过期")
	TokenVerificationFailed  = New(http.StatusConflict, "TokenVerificationFailed", "令牌验证失败")
	TokenRefreshErr          = New(http.StatusConflict, "TokenRefreshErr", "令牌刷新错误")
	TokenStorageFailed       = New(http.StatusConflict, "TokenStorageFailed", "令牌存储失败")
	TokenErrSignatureParam   = New(http.StatusConflict, "TokenErrSignatureParam", "令牌签名错误")
	TokenWrongTypeOfBusiness = New(http.StatusConflict, "TokenWrongTypeOfBusiness", "令牌错误的业务类型")
	TokenGenerationFailed    = New(http.StatusConflict, "TokenGenerationFailed", "令牌生成失败")
)
