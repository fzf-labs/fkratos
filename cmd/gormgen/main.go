package main

import (
	"flag"
	"fmt"
	"github.com/fzf-labs/fpkg/conv"
	"github.com/spf13/viper"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"path/filepath"
	"strings"
	"unicode"
)

var dataMap = map[string]func(detailType string) (dataType string){
	//"int":     func(detailType string) (dataType string) { return "int64" },
	//"tinyint": func(detailType string) (dataType string) { return "int32" },
	"json": func(string) string { return "datatypes.JSON" },
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

var configFile = flag.String("f", "config.yaml", "the config file")
var serverName = flag.String("s", "", "the server name")

func main() {
	flag.Parse()
	db := ConnectDB(GetDsn(*configFile))
	g := gen.NewGenerator(gen.Config{
		OutPath:      fmt.Sprintf("./internal/data/gorm/%sdao", *serverName),
		ModelPkgPath: fmt.Sprintf("./internal/data/gorm/%smodel", *serverName),
	})
	g.UseDB(db)
	g.WithDataTypeMap(dataMap)
	g.WithJSONTagNameStrategy(func(c string) string {
		return LowerCamelCase(c)
	})
	// generate all table from database
	g.ApplyBasic(g.GenerateAllTable()...)
	// apply diy interfaces on structs or table models
	//g.ApplyInterface(func(model.Method) {}, g.GenerateModel("user"))
	g.Execute()
}

func ConnectDB(dsn string) *gorm.DB {
	var db *gorm.DB
	var err error
	db, err = gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(fmt.Errorf("connect db fail: %s", err))
	}
	return db
}

// UpperCamelCase 下划线单词转为大写驼峰单词
func UpperCamelCase(s string) string {
	s = strings.Replace(s, "_", " ", -1)
	causer := cases.Title(language.English)
	s = causer.String(s)
	return strings.Replace(s, " ", "", -1)
}

// LowerCamelCase 下划线单词转为小写驼峰单词
func LowerCamelCase(s string) string {
	s = UpperCamelCase(s)
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}
