package utils

import (
	"aoc/cmd/2023/cmd2023"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{}

func init() {
	Cmd.AddCommand(cmd2023.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
