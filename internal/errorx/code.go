package errorx

import (
	"net/http"

	"github.com/fzf-labs/fkratos-contrib/errx"
)

var Manager = errx.NewErrorManager(errx.WithI18n(errx.EnUS, EnUSMap))

var (
	InternalServerError = Manager.New(http.StatusInternalServerError, "InternalServerError", "服务崩溃了,请稍后再试")
	RequestCanceledErr  = Manager.New(http.StatusConflict, "RequestCanceledErr", "请求被取消")
)

// 参数相关
var (
	ParamBindErr        = Manager.New(http.StatusBadRequest, "ParamBindErr", "参数绑定错误")
	ParamErr            = Manager.New(http.StatusBadRequest, "ParamErr", "参数错误")
	ParamValidationErr  = Manager.New(http.StatusBadRequest, "ParamValidationErr", "参数验证错误")
	ParamNotJSONRequest = Manager.New(http.StatusBadRequest, "ParamNotJSONRequest", "参数不是JSON请求")
	ParamHeaderErr      = Manager.New(http.StatusBadRequest, "ParamHeaderErr", "参数Header错误")
)

// Data数据相关
var (
	DataSQLErr           = Manager.New(http.StatusInternalServerError, "DataSQLErr", "数据处理异常(S),请稍后再试")
	DataRedisErr         = Manager.New(http.StatusInternalServerError, "DataRedisErr", "数据处理异常(R),请稍后再试")
	DataMQErr            = Manager.New(http.StatusInternalServerError, "DataMQErr", "数据处理异常(M),请稍后再试")
	DataRecordNotFound   = Manager.New(http.StatusInternalServerError, "DataRecordNotFound", "数据记录未找到")
	DataDuplicateRecords = Manager.New(http.StatusInternalServerError, "DataDuplicateRecords", "数据重复记录")
	DataFormattingError  = Manager.New(http.StatusInternalServerError, "DataFormattingError", "数据格式化错误")
	DataProcessingError  = Manager.New(http.StatusInternalServerError, "DataProcessingError", "数据处理错误")
)

// API第三方请求
var (
	APIInternalErr = Manager.New(http.StatusConflict, "APIInternalErr", "请求错误(I),请稍后再试")
	APIThirdErr    = Manager.New(http.StatusConflict, "APIThirdErr", "请求错误(T),请稍后再试")
)

// Account账号验证
var (
	AccountNotExist            = Manager.New(http.StatusConflict, "AccountNotExist", "帐户不存在")
	AccountIsLocked            = Manager.New(http.StatusConflict, "AccountIsLocked", "帐号被锁定")
	AccountIsLogout            = Manager.New(http.StatusConflict, "AccountIsLogout", "帐户已注销")
	AccountError               = Manager.New(http.StatusConflict, "AccountError", "帐号错误")
	AccountWrongPassword       = Manager.New(http.StatusConflict, "AccountWrongPassword", "帐号密码错误")
	AccountIsBanned            = Manager.New(http.StatusConflict, "AccountIsBanned", "帐号被封禁")
	AccountUpdateFailed        = Manager.New(http.StatusConflict, "AccountUpdateFailed", "帐号更新失败")
	AccountDuplicateUsername   = Manager.New(http.StatusConflict, "AccountDuplicateUsername", "重复帐户")
	AccountAbnormalStatus      = Manager.New(http.StatusConflict, "AccountAbnormalStatus", "帐户异常状态")
	AccountNicknameViolation   = Manager.New(http.StatusConflict, "AccountNicknameViolation", "帐户昵称违规")
	AccountNotBoundRole        = Manager.New(http.StatusConflict, "AccountNotBoundRole", "帐号未绑定角色")
	AccountVerificationCodeErr = Manager.New(http.StatusConflict, "AccountVerificationCodeErr", "帐户验证码错误")
)

// Token验证
var (
	TokenNotRequest          = Manager.New(http.StatusConflict, "TokenNotRequest", "令牌不存在")
	TokenFormatErr           = Manager.New(http.StatusConflict, "TokenFormatErr", "令牌格式错误")
	TokenErr                 = Manager.New(http.StatusConflict, "TokenErr", "令牌错误")
	TokenInvalidErr          = Manager.New(http.StatusConflict, "TokenInvalidErr", "令牌校验失败")
	TokenExpired             = Manager.New(http.StatusConflict, "TokenExpired", "令牌过期")
	TokenVerificationFailed  = Manager.New(http.StatusConflict, "TokenVerificationFailed", "令牌验证失败")
	TokenRefreshErr          = Manager.New(http.StatusConflict, "TokenRefreshErr", "令牌刷新错误")
	TokenStorageFailed       = Manager.New(http.StatusConflict, "TokenStorageFailed", "令牌存储失败")
	TokenErrSignatureParam   = Manager.New(http.StatusConflict, "TokenErrSignatureParam", "令牌签名错误")
	TokenWrongTypeOfBusiness = Manager.New(http.StatusConflict, "TokenWrongTypeOfBusiness", "令牌错误的业务类型")
	TokenGenerationFailed    = Manager.New(http.StatusConflict, "TokenGenerationFailed", "令牌生成失败")
)
