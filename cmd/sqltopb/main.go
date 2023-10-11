package main

import (
	"flag"
	"fmt"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/fzf-labs/fpkg/orm/gen"
	"github.com/spf13/viper"
)

var configFile = flag.String("f", "config.yaml", "the config file")
var service = flag.String("s", "service", "the service name")

func main() {
	flag.Parse()
	dsn := GetDsn(*configFile)
	connectDB := gen.ConnectDB("postgres", dsn)
	outPutPath := fmt.Sprintf("../../api/%s/v1/", *service)
	packageStr := fmt.Sprintf("api.%s.v1", *service)
	goPackageStr := fmt.Sprintf("fkratos/api/%s/v1;v1", *service)
	gen.NewGenerationPb(connectDB, outPutPath, packageStr, goPackageStr, gen.WithPbOpts(gen.ModelOptionRemoveDefault(), gen.ModelOptionUnderline("UL"))).Do()
}

func GetDsn(configFile string) string {
	config := viper.New()
	config.SetConfigFile(configFile)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return conv.String(config.Get("data.gorm.dataSourceName"))
}
