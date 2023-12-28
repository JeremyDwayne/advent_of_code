package day3

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day3",
	Short: "day3",
	Long:  "day3",
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
	score := 0
	numStart := -1
	numEnd := -1
	lines := strings.Split(s, "\n")
	for x, line := range lines {
		for y, c := range line {
			if c >= '0' && c <= '9' {
				if numStart == -1 {
					numStart = y
				}
				numEnd = y
			} else {
				numEnd = y - 1
				hasAdjacent(lines, x, numStart, numEnd)
			}
		}
	}
	return score
}

func hasAdjacent(lines []string, x, numStart, numEnd int) bool {
	for i := numStart; i <= numEnd; i++ {
		if x > 0 {
			if isSymbol(lines[x+1][i]) {
				return true
			}
		}
		if x < len(lines)-1 {
			if isSymbol(lines[x-1][i]) {
				return true
			}
		}
	}

	if numStart > 0 {
		if isSymbol(lines[x][numStart-1]) {
			return true
		}
		if x > 0 {
			if isSymbol(lines[x][numStart+1]) {
				return true
			}
		}
		if x < len(lines)-1 {
			if isSymbol(lines[x][numStart+1]) {
				return true
			}
		}
	}
	return false
}

func isSymbol(c byte) bool {
	if c >= '0' && c <= '9' {
		return false
	}
	return c != '.'
}
