package service

import (
	"bytes"
	"html/template"
)

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

func New{{ .UpperName }}Service(logger log.Logger, {{ .LowerName }}UseCase *biz.{{ .UpperName }}UseCase) *{{ .UpperName }}Service {
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

{{- $s1 := "google.protobuf.Empty" }}
{{ range .Methods }}
{{- if eq .Type 1 }}
func (s *{{ .UpperName }}Service) {{ .Name }}(ctx context.Context, req {{ if eq .Request $s1 }}*emptypb.Empty{{ else }}*pb.{{ .Request }}{{ end }}) ({{ if eq .Reply $s1 }}*emptypb.Empty{{ else }}*pb.{{ .Reply }}{{ end }}, error) {
	return {{ if eq .Reply $s1 }}&emptypb.Empty{},nil{{ else }}s.{{ .LowerName }}UseCase.{{ .Name }}(ctx, req){{ end }}
}

{{- else if eq .Type 2 }}
func (s *{{ .UpperName }}Service) {{ .Name }}(conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
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
func (s *{{ .UpperName }}Service) {{ .Name }}(conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
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
func (s *{{ .UpperName }}Service) {{ .Name }}(req {{ if eq .Request $s1 }}*emptypb.Empty
{{ else }}*pb.{{ .Request }}{{ end }}, conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
	for {
		err := conn.Send(&pb.{{ .Reply }}{})
		if err != nil {
			return err
		}
	}
}

{{- end }}
{{- end }}
`

type MethodType uint8

const (
	unaryType          MethodType = 1
	twoWayStreamsType  MethodType = 2
	requestStreamsType MethodType = 3
	returnsStreamsType MethodType = 4
)

// Proto is a proto service.
type Proto struct {
	Package     string
	UpperName   string
	LowerName   string
	Methods     []*Method
	GoogleEmpty bool

	UseIO      bool
	UseContext bool
}

// Method is a proto method.
type Method struct {
	UpperName string
	LowerName string
	Name      string
	Request   string
	Reply     string

	// type: unary or stream
	Type MethodType
}

func (s *Proto) execute() ([]byte, error) {
	const empty = "google.protobuf.Empty"
	buf := new(bytes.Buffer)
	for _, method := range s.Methods {
		if (method.Type == unaryType && (method.Request == empty || method.Reply == empty)) ||
			(method.Type == returnsStreamsType && method.Request == empty) {
			s.GoogleEmpty = true
		}
		if method.Type == twoWayStreamsType || method.Type == requestStreamsType {
			s.UseIO = true
		}
		if method.Type == unaryType {
			s.UseContext = true
		}
	}
	tmpl, err := template.New("service").Parse(serviceTemplate)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}