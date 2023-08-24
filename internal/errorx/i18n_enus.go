package errorx

var EnUSMap = map[string]string{
	ParamBindErr.GetReason(): "parameter binding error",
	ParamErr.GetReason():     "parameter error",
	DataSQLErr.GetReason():   "db data exception",
}
