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

// CmdBiz the CmdBiz command.
var CmdBiz = &cobra.Command{
	Use:   "biz",
	Short: "Generate the proto biz implementations",
	Long:  "Generate the proto biz implementations. Example: xxx proto biz api/xxx.proto --target-dir=internal/biz",
	Run:   runBiz,
}
var targetBizDir string

var bizWireTemplate = `
{{- /* delete empty line */ -}}
package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(
	{{ range . }}
	{{- /* delete empty line */ -}}
	{{ . }},
	{{ end }}
)
`

//nolint:lll
var bizServiceTemplate = `
{{- /* delete empty line */ -}}
package biz

import (
	{{- if .UseContext }}
	"context"
	{{- end }}
	{{- if .UseIO }}
	"io"
	{{- end }}

	pb "{{ .Package }}"
	{{- if .GoogleEmpty }}
	"google.golang.org/protobuf/types/known/emptypb"
	{{- end }}
)

type {{ .UpperName }}Repo interface {
}

func New{{ .UpperName }}UseCase(
	logger log.Logger,
	{{ .LowerName }}Repo {{ .UpperName }}Repo,
) *{{ .UpperName }}UseCase {
	l := log.NewHelper(log.With(logger, "module", "biz/{{ .LowerName }}"))
	return &{{ .UpperName }}UseCase{
		log:         l,
		{{ .LowerName }}Repo:{{ .LowerName }}Repo,
	}
}

type {{ .UpperName }}UseCase struct {
	log *log.Helper
	{{ .LowerName }}Repo {{ .UpperName }}Repo
}
`

//nolint:lll
var bizMethodTemplate = `
{{- /* delete empty line */ -}}
package biz

import (
	{{- if .UseContext }}
	"context"
	{{- end }}
	{{- if .UseIO }}
	"io"
	{{- end }}

	pb "{{ .Package }}"
	{{- if .GoogleEmpty }}
	"google.golang.org/protobuf/types/known/emptypb"
	{{- end }}
)
{{- $s1 := "google.protobuf.Empty" }}
{{- if eq .Type 1 }}
// {{ .Name }} {{ .Comment }}
func ({{.FirstChar}} *{{ .UpperName }}UseCase) {{ .Name }}(ctx context.Context, req {{ if eq .Request $s1 }}*emptypb.Empty{{ else }}*pb.{{ .Request }}{{ end }}) ({{ if eq .Reply $s1 }}*emptypb.Empty{{ else }}*pb.{{ .Reply }}{{ end }}, error) {
	{{- if eq .Reply $s1 }}
 	resp :=	&emptypb.Empty{}
	{{- else }}
 	resp :=	&pb.{{ .Reply }}{}
	{{- end }}
	return resp, nil
}

{{- else if eq .Type 2 }}
// {{ .Name }} {{ .Comment }}
func ({{.FirstChar}} *{{ .UpperName }}UseCase) {{ .Name }}(conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
	for {
		req, err := conn.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		
		err = conn.Send(&pb.{{ .Reply }}{})
		if err != nil {
			return err
		}
	}
}

{{- else if eq .Type 3 }}
// {{ .Name }} {{ .Comment }}
func ({{.FirstChar}} *{{ .UpperName }}UseCase) {{ .Name }}(conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
	for {
		req, err := conn.Recv()
		if err == io.EOF {
			return conn.SendAndClose(&pb.{{ .Reply }}{})
		}
		if err != nil {
			return err
		}
	}
}

{{- else if eq .Type 4 }}
// {{ .Name }} {{ .Comment }}
func ({{.FirstChar}} *{{ .UpperName }}UseCase) {{ .Name }}(req {{ if eq .Request $s1 }}*emptypb.Empty
{{ else }}*pb.{{ .Request }}{{ end }}, conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
	for {
		err := conn.Send(&pb.{{ .Reply }}{})
		if err != nil {
			return err
		}
	}
}

{{- end }}
`

func init() {
	CmdBiz.Flags().StringVarP(&targetBizDir, "target-dir", "t", "internal/biz", "generate target directory")
}

func runBiz(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the proto path.")
		return
	}
	pbFiles, err := pb.GetPathPbFiles(args[0])
	if err != nil {
		return
	}
	news := make([]string, 0)
	for _, file := range pbFiles {
		res := pb.Translate(file)
		if _, err := os.Stat(targetBizDir); os.IsNotExist(err) {
			fmt.Printf("Target directory: %s does not exsit\n", targetBizDir)
			return
		}
		for _, s := range res {
			news = append(news, fmt.Sprintf("New%sUseCase", s.UpperName))
			err := s.Execute()
			if err != nil {
				log.Fatal(err)
			}
			toService := filepath.Join(targetBizDir, strings.ToLower(s.UpperName)+".go")
			if _, err := os.Stat(toService); !os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "biz new already exists: %s\n", toService)
			} else {
				b, err := s.ExecuteService(bizServiceTemplate)
				if err != nil {
					return
				}
				if err := pb.Output(toService, b); err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(os.Stderr, "biz new generated successfully: %s\n", toService)
			}
			for _, method := range s.Methods {
				toMethod := filepath.Join(targetBizDir, strings.ToLower(s.UpperName+"_"+method.Name)+".go")
				if _, err := os.Stat(toMethod); !os.IsNotExist(err) {
					fmt.Fprintf(os.Stderr, "biz method already exists: %s\n", toMethod)
					continue
				}
				b, err := s.ExecuteMethod(bizMethodTemplate, method)
				if err != nil {
					fmt.Println(err)
					return
				}
				if err := pb.Output(toMethod, b); err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(os.Stderr, "biz method generated successfully: %s\n", toMethod)
			}
		}
	}
	toBizWire := filepath.Join(targetBizDir, "biz.go")
	execute, err := pb.TemplateExecute(bizWireTemplate, news)
	if err != nil {
		return
	}
	if err := pb.Output(toBizWire, execute); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stderr, "biz wire generated successfully: %s\n", toBizWire)
}
