package cmd2023

import (
	"aoc/cmd/2023/day1"
	"aoc/cmd/2023/day2"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "2023",
	Short: "2023",
	Long:  "2023 Advent of Code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	Cmd.AddCommand(day1.Cmd)
	Cmd.AddCommand(day2.Cmd)
}

func Execute() {
	if err := Cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
