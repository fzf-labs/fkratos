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
	gen.NewGeneration(connectDB, "./internal/data/gorm", gen.WithDataMap(gen.DefaultPostgresDataMap), gen.WithOpts(gen.ModelOptionPgDefaultString(), gen.ModelOptionUnderline("UL"))).Do()
}

func GetDsn(configFile string) string {
	config := viper.New()
	config.SetConfigFile(configFile)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return conv.String(config.Get("data.gorm.dataSourceName"))
}
