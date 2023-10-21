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

// CmdService the service command.
var CmdService = &cobra.Command{
	Use:   "service",
	Short: "Generate the proto service implementations",
	Long:  "Generate the proto service implementations. Example: kratos proto service api/xxx.proto --target-dir=internal/service",
	Run:   runService,
}
var targetServiceDir string

func init() {
	CmdService.Flags().StringVarP(&targetServiceDir, "target-dir", "t", "internal/service", "generate target directory")
}

var serviceWireTemplate = `
{{- /* delete empty line */ -}}
package service

import "github.com/google/wire"

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	{{ range . }}
	{{- /* delete empty line */ -}}
	{{ . }},
	{{ end }}
)
`

//nolint:lll
var serviceTemplate = `
{{- /* delete empty line */ -}}
package service

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

func New{{ .UpperName }}Service(
	logger log.Logger,
	{{ .LowerName }}UseCase *biz.{{ .UpperName }}UseCase,
) *{{ .UpperName }}Service {
	l := log.NewHelper(log.With(logger, "module", "service/{{ .LowerName }}"))
	return &{{ .UpperName }}Service{
		log:         l,
		{{ .LowerName }}UseCase:{{ .LowerName }}UseCase,
	}
}

type {{ .UpperName }}Service struct {
	pb.Unimplemented{{ .UpperName }}Server
	log *log.Helper
	{{ .LowerName }}UseCase *biz.{{ .UpperName }}UseCase
}
`

//nolint:lll
var methodTemplate = `
{{- /* delete empty line */ -}}
package service

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
func ({{.FirstChar}} *{{ .UpperName }}Service) {{ .Name }}(ctx context.Context, req {{ if eq .Request $s1 }}*emptypb.Empty{{ else }}*pb.{{ .Request }}{{ end }}) ({{ if eq .Reply $s1 }}*emptypb.Empty{{ else }}*pb.{{ .Reply }}{{ end }}, error) {
	return {{ if eq .Reply $s1 }}&emptypb.Empty{},nil{{ else }}{{.FirstChar}}.{{ .LowerName }}UseCase.{{ .Name }}(ctx, req){{ end }}
}

{{- else if eq .Type 2 }}
func ({{.FirstChar}} *{{ .UpperName }}Service) {{ .Name }}(conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
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
func ({{.FirstChar}} *{{ .UpperName }}Service) {{ .Name }}(conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
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
func ({{.FirstChar}} *{{ .UpperName }}Service) {{ .Name }}(req {{ if eq .Request $s1 }}*emptypb.Empty
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

func runService(_ *cobra.Command, args []string) {
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
		if _, err := os.Stat(targetServiceDir); os.IsNotExist(err) {
			fmt.Printf("Target directory: %s does not exsit\n", targetServiceDir)
			return
		}
		for _, s := range res {
			news = append(news, fmt.Sprintf("New%sService", s.UpperName))
			err := s.Execute()
			if err != nil {
				log.Fatal(err)
			}
			toService := filepath.Join(targetServiceDir, strings.ToLower(s.UpperName)+".go")
			if _, err := os.Stat(toService); !os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "service new already exists: %s\n", toService)
			} else {
				b, err := s.ExecuteService(serviceTemplate)
				if err != nil {
					return
				}
				if err := pb.Output(toService, b); err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(os.Stderr, "service new generated successfully: %s\n", toService)
			}
			for _, method := range s.Methods {
				toMethod := filepath.Join(targetServiceDir, strings.ToLower(s.UpperName+"_"+method.Name)+".go")
				if _, err := os.Stat(toMethod); !os.IsNotExist(err) {
					fmt.Fprintf(os.Stderr, "service method already exists: %s\n", toMethod)
					continue
				}
				b, err := s.ExecuteMethod(methodTemplate, method)
				if err != nil {
					fmt.Println(err)
					return
				}
				if err := pb.Output(toMethod, b); err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(os.Stderr, "service method generated successfully: %s\n", toMethod)
			}
		}
	}
	toServiceWire := filepath.Join(targetServiceDir, "service.go")
	execute, err := pb.TemplateExecute(serviceWireTemplate, news)
	if err != nil {
		return
	}
	if err := pb.Output(toServiceWire, execute); err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(os.Stderr, "service wire generated successfully: %s\n", toServiceWire)
}
