package main

import (
	"flag"
	"fmt"
	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/db/gen"
	"github.com/spf13/viper"
	"path/filepath"
)

var configFile = flag.String("f", "config.yaml", "the config file")
var serverName = flag.String("s", "", "the server name")

// 默认Postgres字段类型映射
var defaultPostgresDataMap = map[string]func(detailType string) (dataType string){
	"json": func(string) string { return "datatypes.JSON" },
}

func main() {
	flag.Parse()
	db := GetDsn(*configFile)
	outPath := fmt.Sprintf("./internal/data/gorm/%sdao", *serverName)
	modelPkgPath := fmt.Sprintf("./internal/data/gorm/%smodel", *serverName)
	gen.Generation(gen.ConnectDB("postgres", db), defaultPostgresDataMap, outPath, modelPkgPath)
}

func GetDsn(configFile string) string {
	config := viper.New()
	config.AddConfigPath(filepath.Dir(configFile)) //设置读取的文件路径
	config.SetConfigName("config")                 //设置读取的文件名
	config.SetConfigType("yaml")                   //设置文件的类型
	//尝试进行配置读取
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return conv.String(config.Get("data.database.source"))
}
