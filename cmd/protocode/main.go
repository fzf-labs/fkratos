//nolint:all
package main

import (
	"fkratos/cmd/protocode/gen"
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
	CmdProto.AddCommand(gen.CmdService)
	CmdProto.AddCommand(gen.CmdBiz)
	CmdProto.AddCommand(gen.CmdData)
}

func main() {
	if err := CmdProto.Execute(); err != nil {
		log.Fatal(err)
	}
}
