package biz

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"unicode"

	"github.com/emicklei/proto"
	"github.com/spf13/cobra"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"golang.org/x/tools/imports"
)

// CmdService the service command.
var CmdService = &cobra.Command{
	Use:   "biz",
	Short: "Generate the proto biz implementations",
	Long:  "Generate the proto biz implementations. Example: xxx proto biz api/xxx.proto --target-dir=internal/biz",
	Run:   run,
}
var targetDir string

func init() {
	CmdService.Flags().StringVarP(&targetDir, "target-dir", "t", "internal/biz", "generate target directory")
}

func run(_ *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "Please specify the proto file. Example: xxx proto biz api/xxx.proto")
		return
	}
	reader, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	parser := proto.NewParser(reader)
	definition, err := parser.Parse()
	if err != nil {
		log.Fatal(err)
	}

	var (
		pkg string
		res []*Proto
	)
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
			}
			for _, e := range s.Elements {
				r, ok := e.(*proto.RPC)
				if !ok {
					continue
				}
				cs.Methods = append(cs.Methods, &Method{
					UpperName: UpperName(s.Name),
					LowerName: LowerName(s.Name),
					Name:      UpperName(r.Name),
					Request:   parametersName(r.RequestType),
					Reply:     parametersName(r.ReturnsType), Type: getMethodType(r.StreamsRequest, r.StreamsReturns),
				})
			}
			res = append(res, cs)
		}),
	)
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		fmt.Printf("Target directory: %s does not exsit\n", targetDir)
		return
	}
	for _, s := range res {
		to := filepath.Join(targetDir, strings.ToLower(s.UpperName)+".go")
		if _, err := os.Stat(to); !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "%s already exists: %s\n", s.UpperName, to)
			continue
		}
		b, err := s.execute()
		if err != nil {
			log.Fatal(err)
		}
		if err := output(to, b); err != nil {
			log.Fatal(err)
		}
		fmt.Println(to)
	}
}

func output(fileName string, content []byte) error {
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
	return os.WriteFile(fileName, result, 0o644)
}

func getMethodType(streamsRequest, streamsReturns bool) MethodType {
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

func parametersName(name string) string {
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
