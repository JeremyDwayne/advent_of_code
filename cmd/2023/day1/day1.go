package day1

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day1",
	Short: "day1",
	Long:  "day1",
	Run: func(cmd *cobra.Command, _ []string) {
		execute(cmd.Parent().Name(), cmd.Name())
	},
}

func execute(parent, command string) {
	b, err := os.ReadFile(fmt.Sprintf(`cmd/%s/%s/input.txt`, parent, command))
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.Infof("score part1: %d", part1(string(b)))
}

func part1(s string) int {
	sum := 0
	for _, line := range strings.Split(s, "\n") {
		first := ""
		last := ""
		for _, char := range line {
			if _, err := strconv.Atoi(string(char)); err == nil {
				if first == "" {
					first = string(char)
				}
				last = string(char)
			}
		}
		if value, err := strconv.Atoi(first + last); err == nil {
			sum += value
		}
	}
	return sum
}
