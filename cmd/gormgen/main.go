package main

import (
	"flag"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/orm/gen"
	"github.com/spf13/viper"
)

var configFile = flag.String("f", "config.yaml", "the config file")

func main() {
	flag.Parse()
	dsn := GetDsn(*configFile)
	connectDB := gen.ConnectDB("postgres", dsn)
	gen.NewGenerationDB(connectDB, "./internal/data/gorm", gen.WithDataMap(gen.DataTypeMap()), gen.WithDBOpts(gen.ModelOptionRemoveDefault(), gen.ModelOptionUnderline("UL"))).Do()
}

func GetDsn(configFile string) string {
	config := viper.New()
	config.SetConfigFile(configFile)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return conv.String(config.Get("data.gorm.dataSourceName"))
}
