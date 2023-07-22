package main

import (
	"flag"
	"path/filepath"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/db/gen"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var configFile = flag.String("f", "config.yaml", "the config file")

// 默认Postgres字段类型映射
var defaultPostgresDataMap = map[string]func(columnType gorm.ColumnType) (dataType string){
	"json": func(columnType gorm.ColumnType) string { return "datatypes.JSON" },
	"timestamptz": func(columnType gorm.ColumnType) string {
		nullable, _ := columnType.Nullable()
		if nullable {
			return "sql.NullTime"
		}
		return "time.Time"
	},
}

func main() {
	flag.Parse()
	db := GetDsn(*configFile)
	outPath := "./internal/data/gorm"
	gen.Generation(gen.ConnectDB("postgres", db), defaultPostgresDataMap, outPath)
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
