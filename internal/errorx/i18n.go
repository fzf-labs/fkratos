package errorx

const (
	ZhCN = "zh-CN" // zh_CN 简体中文-中国
	EnUS = "en-US" // en_US 英文-美国
)

func GetMessage(reason string, lang string) string {
	var msg string
	switch lang {
	case ZhCN:
		msg = ZhCNMap[reason]
	case EnUS:
		msg = EnUSMap[reason]
	default:
		msg = ZhCNMap[reason]
	}
	return msg
}
