package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/util/fileutil"
	"github.com/imroc/req/v3"
	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()
	config := GetConfig(*configFile)
	name := conv.String(config.Get("name"))
	accessToken := conv.String(config.Get("apiDoc.apiFox.accessToken"))
	projectID := conv.String(config.Get("apiDoc.apiFox.projectID"))
	if name == "" || accessToken == "" || projectID == "" {
		slog.Error("Missing configuration")
		return
	}
	// 查询指定目录下的swagger的json文件
	dir := fmt.Sprintf("../../api/%s", name)
	files, err := fileutil.ReadAllFileToSli(dir)
	if err != nil {
		slog.Error(err.Error())
		return
	}
	apiFox := NewAPIFox(accessToken)
	for _, v := range files {
		if strings.Contains(v.Name, ".swagger.json") {
			// 读取文件内容
			toString, err := fileutil.ReadFileToString(v.File)
			if err != nil {
				slog.Error(err.Error())
				return
			}
			err = apiFox.SyncHTTP(projectID, toString)
			if err != nil {
				slog.Error(err.Error())
				return
			}
		}
	}
	slog.Info("apifox import successful")
}

func GetConfig(configFile string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(configFile)
	// 尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return config
}

type APIFox struct {
	AccessToken string `json:"accessToken"`
}

func NewAPIFox(accessToken string) *APIFox {
	return &APIFox{AccessToken: accessToken}
}

type APIFoxHTTPParam struct {
	// 导入数据格式，目前只支持`openapi`，表示 Swagger 或 OpenAPI 格式
	ImportFormat string `json:"importFormat"`
	// 要导入的数据，Swagger（OpenAPI） 格式 json 字符串，支持 OpenAPI 3、Swagger 1、2、3 数据格式
	Data string `json:"data"`
	// 导入到目标目录的ID，不传表示导入到根目录
	APIFolderID float64 `json:"apiFolderId,omitempty"`
	// 覆盖模式，匹配到相同接口时的覆盖模式，不传表示忽略
	APIOverwriteMode string `json:"apiOverwriteMode,omitempty"`
	// 是否在接口路径加上basePath，建议不传，即为 false，推荐将 BasePath 放到环境里的”前置 URL“里
	ImportBasePath bool `json:"importBasePath,omitempty"`
	// 覆盖模式，匹配到相同数据模型时的覆盖模式，不传表示忽略
	SchemaOverwriteMode string `json:"schemaOverwriteMode,omitempty"`
	// 是否同步更新接口所在目录
	SyncAPIFolder bool `json:"syncApiFolder,omitempty"`
}

// SyncHTTP 同步Http文档
func (a *APIFox) SyncHTTP(projectId, data string) error {
	url := fmt.Sprintf("https://api.apifox.cn/api/v1/projects/%s/import-data", projectId)
	headers := map[string]string{
		"X-Apifox-Version": "2022-11-16",
		"Authorization":    fmt.Sprintf("Bearer %s", a.AccessToken),
	}
	body := APIFoxHTTPParam{
		ImportFormat:        "openapi",
		Data:                data,
		APIOverwriteMode:    "methodAndPath",
		SchemaOverwriteMode: "name",
	}
	_, err := req.R().SetHeaders(headers).SetBody(body).Post(url)
	if err != nil {
		return err
	}
	return nil
}
