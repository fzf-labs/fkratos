package data

import (
	"bytes"
	"html/template"
)

//nolint:lll
var dataTemplate = `
{{- /* delete empty line */ -}}
package data

import (
	"context"
)

var _ biz.{{ .UpperName }}Repo = (*{{ .UpperName }}Repo)(nil)

func New{{ .UpperName }}Repo(logger log.Logger,data *Data) biz.{{ .UpperName }}Repo {
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
	tmpl, err := template.New("data").Parse(dataTemplate)
	if err != nil {
		return nil, err
	}
	if err := tmpl.Execute(buf, s); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
