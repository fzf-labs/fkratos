package main

import (
	"flag"
	"path/filepath"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/db"
	"github.com/fzf-labs/fpkg/db/gen"
	"github.com/spf13/viper"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()
	dsn := GetDsn(*configFile)
	connectDB := gen.ConnectDB("postgres", dsn)
	db.DumpPostgres(connectDB, dsn, "../../doc/sql")
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
	return conv.String(config.Get("data.gorm.dataSourceName"))
}
