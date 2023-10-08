package biz

import (
	"bytes"
	"html/template"
)

//nolint:lll
var bizTemplate = `
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

func New{{ .UpperName }}UseCase(logger log.Logger,{{ .LowerName }}Repo {{ .UpperName }}Repo) *{{ .UpperName }}UseCase {
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

{{- $s1 := "google.protobuf.Empty" }}
{{ range .Methods }}
{{- if eq .Type 1 }}
func (s *{{ .UpperName }}UseCase) {{ .Name }}(ctx context.Context, req {{ if eq .Request $s1 }}*emptypb.Empty{{ else }}*pb.{{ .Request }}{{ end }}) ({{ if eq .Reply $s1 }}*emptypb.Empty{{ else }}*pb.{{ .Reply }}{{ end }}, error) {
	{{- if eq .Reply $s1 }}
 	resp :=	&emptypb.Empty{}
	{{- else }}
 	resp :=	&pb.{{ .Reply }}{}
	{{- end }}
	return resp, nil
}

{{- else if eq .Type 2 }}
func (s *{{ .UpperName }}UseCase) {{ .Name }}(conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
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
func (s *{{ .UpperName }}UseCase) {{ .Name }}(conn pb.{{ .UpperName }}_{{ .UpperName }}Server) error {
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
func (s *{{ .UpperName }}UseCase) {{ .Name }}(req {{ if eq .Request $s1 }}*emptypb.Empty
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
	tmpl, err := template.New("biz").Parse(bizTemplate)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
