//nolint:all
package gen

import (
	"fkratos/cmd/protocode/pb"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/fzf-labs/fpkg/conv"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// CmdData the service command.
var CmdData = &cobra.Command{
	Use:   "data",
	Short: "Generate the proto data implementations",
	Long:  "Generate the proto data implementations. Example: xxx proto data api/xxx.proto --target-dir=internal/data",
	Run:   runData,
}

var targetConfigFile string
var targetTables string
var targetDataDir string

var dataServiceTemplate = `
{{- /* delete empty line */ -}}
package data

import (
	"context"
)

var _ biz.{{ .UpperName }}Repo = (*{{ .UpperName }}Repo)(nil)

func New{{ .UpperName }}Repo(
	logger log.Logger,
	data *Data,
) biz.{{ .UpperName }}Repo {
	l := log.NewHelper(log.With(logger, "module", "data/{{ .LowerName }}"))
	return &{{ .UpperName }}Repo{
		log:         l,
		data:     data,
	}
}

type {{ .UpperName }}Repo struct {
	log *log.Helper
	data *Data
}
`

func init() {
	CmdData.Flags().StringVarP(&targetConfigFile, "config", "f", "configs/config.yaml", "configs file")
	CmdData.Flags().StringVarP(&targetDataDir, "dir", "d", "internal/data", "generate target directory")
	CmdData.Flags().StringVarP(&targetTables, "tables", "t", "", "tables")
}

type DB struct {
	UpperName string
	LowerName string
}

func runData(_ *cobra.Command, args []string) {
	if targetConfigFile == "" {
		fmt.Fprintln(os.Stderr, "Please specify the configs file")
		return
	}
	if targetDataDir == "" {
		fmt.Fprintln(os.Stderr, "Please specify the target directory")
		return
	}
	dsn := GetDsn(targetConfigFile)
	db := ConnectDB("postgres", dsn)
	tables := make([]string, 0)
	if targetTables != "" {
		tables = strings.Split(targetTables, ",")
	} else {
		tables, _ = db.Migrator().GetTables()
	}
	if _, err := os.Stat(targetDataDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exsit\n", targetDataDir)
		return
	}
	for _, table := range tables {
		tmp := DB{
			UpperName: UpperName(db, table),
			LowerName: LowerName(db, table),
		}
		toService := filepath.Join(targetDataDir, strings.ToLower(tmp.UpperName)+".go")
		if _, err := os.Stat(toService); !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "data new already exists: %s\n", toService)
		} else {
			b, err := pb.TemplateExecute(dataServiceTemplate, tmp)
			if err != nil {
				return
			}
			if err := pb.Output(toService, b); err != nil {
				log.Fatal(err)
			}
			fmt.Fprintf(os.Stderr, "data new generated successfully: %s\n", toService)
		}
	}
}

func GetDsn(configFile string) string {
	config := viper.New()
	config.SetConfigFile(configFile)
	if err := config.ReadInConfig(); err != nil {
		panic(err)
	}
	return conv.String(config.Get("data.gorm.dataSourceName"))
}

// ConnectDB 数据库连接
func ConnectDB(dbType, dsn string) *gorm.DB {
	var db *gorm.DB
	var err error
	switch dbType {
	case "mysql":
		db, err = gorm.Open(mysql.Open(dsn))
		if err != nil {
			panic(fmt.Errorf("connect mysql db fail: %s", err))
		}
	case "postgres":
		db, err = gorm.Open(postgres.Open(dsn))
		if err != nil {
			panic(fmt.Errorf("connect postgres db fail: %s", err))
		}
	default:
		panic(fmt.Errorf(" db type err"))
	}
	return db
}

// UpperName 大写
func UpperName(db *gorm.DB, s string) string {
	return db.NamingStrategy.SchemaName(s)
}

// LowerName 小写
func LowerName(db *gorm.DB, s string) string {
	str := UpperName(db, s)
	if str == "" {
		return str
	}
	words := []string{"API", "ASCII", "CPU", "CSS", "DNS", "EOF", "GUID", "HTML", "HTTP", "HTTPS", "ID", "IP", "JSON", "LHS", "QPS", "RAM", "RHS", "RPC", "SLA", "SMTP", "SSH", "TLS", "ttl", "UID", "UI", "UUID", "URI", "URL", "UTF8", "VM", "XML", "XSRF", "XSS"}
	// 如果第一个单词命中  则不处理
	for _, v := range words {
		if strings.HasPrefix(str, v) {
			return str
		}
	}
	rs := []rune(str)
	f := rs[0]
	if 'A' <= f && f <= 'Z' {
		str = string(unicode.ToLower(f)) + string(rs[1:])
	}
	return str
}
