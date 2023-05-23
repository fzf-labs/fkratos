package errorx

var EnUSMap = map[string]string{
	ParamBindErr.GetReason(): "parameter binding error",
	ParamErr.GetReason():     "parameter error",
	DataSqlErr.GetReason():   "db data exception",
}
