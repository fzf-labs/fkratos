package errorx

var ZhCNMap = map[string]string{
	ParamBindErr.GetReason(): "参数绑定错误",
	ParamErr.GetReason():     "参数错误",
	DataSQLErr.GetReason():   "DB数据异常",
}
