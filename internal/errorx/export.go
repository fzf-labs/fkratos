package errorx

import (
	"encoding/json"
	"strconv"

	"github.com/fzf-labs/fpkg/util/fileutil"
	"golang.org/x/exp/slog"
)

func Export() {
	list := exportData()
	exportJSON(list)
	exportMarkdown(list)
}

// 数据生成
func exportData() []map[string]string {
	list := make([]map[string]string, 0)
	for _, v := range reasons {
		m := make(map[string]string)
		m["http_code"] = strconv.Itoa(int(errs[v].GetCode()))
		m["reason"] = errs[v].GetReason()
		m["message"] = errs[v].GetMessage()
		for _, lang := range GetLanguages() {
			m[lang] = GetMessage(errs[v].GetReason(), lang)
		}
		list = append(list, m)
	}
	return list
}

// 导出json
func exportJSON(list []map[string]string) {
	marshal, err := json.Marshal(list)
	if err != nil {
		return
	}
	err = fileutil.WriteContentCover("../../doc/errcode/code.json", string(marshal))
	if err != nil {
		return
	}
	slog.Info("错误码JSON导出成功")
}

// 导出markdown
func exportMarkdown(list []map[string]string) {
	first := `|http_code|reason|message|`
	second := `|--|--|--|`
	for _, v := range GetLanguages() {
		first += v + `|`
		second += `--|`
	}
	str := NewLine(first) + NewLine(second)
	if len(list) > 0 {
		for _, m := range list {
			tmpStr := `|` + m["http_code"] + `|` + m["reason"] + `|` + m["message"] + `|`
			for _, v := range GetLanguages() {
				tmpStr += m[v] + `|`
			}
			str += NewLine(tmpStr)
		}
	}
	err := fileutil.WriteContentCover("../../doc/errcode/code.md", str)
	if err != nil {
		return
	}
	slog.Info("错误码MARKDOWN导出成功")
}

func NewLine(str string) string {
	return str + "\n"
}
