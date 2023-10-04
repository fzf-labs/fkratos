package main

import (
	"fkratos/cmd/proto/biz"
	"fkratos/cmd/proto/data"
	"fkratos/cmd/proto/service"
	"log"

	"github.com/spf13/cobra"
)

// CmdProto represents the proto command.
var CmdProto = &cobra.Command{
	Use:   "proto",
	Short: "Generate the proto files",
	Long:  "Generate the proto files.",
}

func init() {
	CmdProto.AddCommand(service.CmdService)
	CmdProto.AddCommand(biz.CmdService)
	CmdProto.AddCommand(data.CmdService)
}

func main() {
	if err := CmdProto.Execute(); err != nil {
		log.Fatal(err)
	}
}
