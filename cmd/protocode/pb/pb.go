package pb

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/emicklei/proto"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/imports"
)

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
	FirstChar   string
	GoogleEmpty bool
	UseIO       bool
	UseContext  bool
	Methods     []*Method
}

// Method is a proto method.
type Method struct {
	Package     string
	UpperName   string
	LowerName   string
	FirstChar   string
	GoogleEmpty bool
	UseIO       bool
	UseContext  bool

	Name    string
	Request string
	Reply   string
	Type    MethodType // type: unary or stream
	Comment string
}

func (s *Proto) Execute() error {
	const empty = "google.protobuf.Empty"
	for k, method := range s.Methods {
		s.Methods[k].Package = s.Package
		s.Methods[k].UpperName = s.UpperName
		s.Methods[k].LowerName = s.LowerName
		if (method.Type == unaryType && (method.Request == empty || method.Reply == empty)) ||
			(method.Type == returnsStreamsType && method.Request == empty) {
			s.GoogleEmpty = true
			s.Methods[k].GoogleEmpty = true
		}
		if method.Type == twoWayStreamsType || method.Type == requestStreamsType {
			s.UseIO = true
			s.Methods[k].UseIO = true
		}
		if method.Type == unaryType {
			s.UseContext = true
			s.Methods[k].UseContext = true
		}
	}
	return nil
}

func (s *Proto) ExecuteService(tpl string) ([]byte, error) {
	return TemplateExecute(tpl, s)
}

func (s *Proto) ExecuteMethod(tpl string, method *Method) ([]byte, error) {
	return TemplateExecute(tpl, method)
}

func Output(fileName string, content []byte) error {
	result, err := imports.Process(fileName, content, nil)
	if err != nil {
		lines := strings.Split(string(content), "\n")
		errLine, _ := strconv.Atoi(strings.Split(err.Error(), ":")[1])
		startLine, endLine := errLine-5, errLine+5
		fmt.Println("Format fail:", errLine, err)
		if startLine < 0 {
			startLine = 0
		}
		if endLine > len(lines)-1 {
			endLine = len(lines) - 1
		}
		for i := startLine; i <= endLine; i++ {
			fmt.Println(i, lines[i])
		}
		return fmt.Errorf("cannot format file: %w", err)
	}
	return os.WriteFile(fileName, result, 0600)
}

func GetMethodType(streamsRequest, streamsReturns bool) MethodType {
	if !streamsRequest && !streamsReturns {
		return unaryType
	} else if streamsRequest && streamsReturns {
		return twoWayStreamsType
	} else if streamsRequest {
		return requestStreamsType
	} else if streamsReturns {
		return returnsStreamsType
	}
	return unaryType
}

func ParametersName(name string) string {
	return strings.ReplaceAll(name, ".", "_")
}

func UpperName(name string) string {
	return toUpperCamelCase(strings.Split(name, ".")[0])
}

func LowerName(name string) string {
	return toLowerCamelCase(strings.Split(name, ".")[0])
}

func toUpperCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = cases.Title(language.Und, cases.NoLower).String(s)
	return strings.ReplaceAll(s, " ", "")
}

func toLowerCamelCase(s string) string {
	s = strings.ReplaceAll(s, "_", " ")
	s = LcFirst(cases.Title(language.Und, cases.NoLower).String(s))
	return strings.ReplaceAll(s, " ", "")
}

// LcFirst 首字母小写
func LcFirst(s string) string {
	if s == "" {
		return s
	}

	rs := []rune(s)
	f := rs[0]

	if 'A' <= f && f <= 'Z' {
		return string(unicode.ToLower(f)) + string(rs[1:])
	}
	return s
}

func Translate(path string) []*Proto {
	var (
		pkg string
		res []*Proto
	)
	reader, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer reader.Close()
	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Println(err)
		return nil
	}
	proto.Walk(definition,
		proto.WithOption(func(o *proto.Option) {
			if o.Name == "go_package" {
				pkg = strings.Split(o.Constant.Source, ";")[0]
			}
		}),
		proto.WithService(func(s *proto.Service) {
			cs := &Proto{
				Package:   pkg,
				UpperName: UpperName(s.Name),
				LowerName: LowerName(s.Name),
				FirstChar: strings.ToLower(s.Name[0:1]),
			}
			for _, e := range s.Elements {
				r, ok := e.(*proto.RPC)
				if !ok {
					continue
				}
				cs.Methods = append(cs.Methods, &Method{
					Name:      UpperName(r.Name),
					Request:   ParametersName(r.RequestType),
					Reply:     ParametersName(r.ReturnsType),
					Type:      GetMethodType(r.StreamsRequest, r.StreamsReturns),
					Comment:   strings.TrimSpace(r.Comment.Message()),
					FirstChar: cs.FirstChar,
				})
			}
			res = append(res, cs)
		}),
	)
	return res
}

func TemplateExecute(tpl string, data any) ([]byte, error) {
	buf := new(bytes.Buffer)
	tmpl, err := template.New("tpl").Parse(tpl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	if err := tmpl.Execute(buf, data); err != nil {
		log.Println(err)
		return nil, err
	}
	return buf.Bytes(), nil
}

// GetPathPbFiles 获取目录下指定后缀文件
func GetPathPbFiles(p string) ([]string, error) {
	var files []string
	err := filepath.Walk(p, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".proto") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return files, nil
}
