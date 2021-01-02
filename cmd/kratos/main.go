package main

import (
	"log"

	"github.com/spf13/cobra"

	"github.com/go-kratos/kratos/v2/cmd/kratos/internal/new"
	"github.com/go-kratos/kratos/v2/cmd/kratos/internal/proto"
)

var rootCmd = &cobra.Command{
	Use:   "kratos",
	Short: "Kratos: An elegant toolkit for Go microservices.",
	Long:  `Kratos: An elegant toolkit for Go microservices.`,
}

func init() {
	rootCmd.AddCommand(new.CmdNew)
	rootCmd.AddCommand(proto.CmdProto)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
