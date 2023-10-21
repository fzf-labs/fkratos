//nolint:all
package gen

import (
	"fkratos/cmd/proto/pb"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

// CmdData the service command.
var CmdData = &cobra.Command{
	Use:   "data",
	Short: "Generate the proto data implementations",
	Long:  "Generate the proto data implementations. Example: xxx proto data api/xxx.proto --target-dir=internal/data",
	Run:   runData,
}

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
	CmdData.Flags().StringVarP(&targetDataDir, "target-dir", "t", "internal/data", "generate target directory")
}

func runData(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the proto file. Example: xxx proto service api/xxx.proto")
		return
	}
	pbFiles, err := pb.GetPathPbFiles(args[0])
	if err != nil {
		return
	}
	for _, file := range pbFiles {
		res := pb.Translate(file)
		if _, err := os.Stat(targetDataDir); os.IsNotExist(err) {
			fmt.Printf("Target directory: %s does not exsit\n", targetDataDir)
			return
		}
		for _, s := range res {
			err := s.Execute()
			if err != nil {
				log.Fatal(err)
			}
			toService := filepath.Join(targetDataDir, strings.ToLower(s.UpperName)+".go")
			if _, err := os.Stat(toService); !os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "data new already exists: %s\n", toService)
			} else {
				b, err := s.ExecuteService(dataServiceTemplate)
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
}
