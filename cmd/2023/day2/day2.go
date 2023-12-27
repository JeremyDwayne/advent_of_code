package day2

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "day2",
	Short: "day2",
	Long:  "day2",
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
	bag := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0
	matcher := regexp.MustCompile(`Game (\d+): (.*)`)

	for _, line := range strings.Split(s, "\n") {
		ok := true
		matched := matcher.FindAllStringSubmatch(line, -1)

		if len(matched) == 0 {
			continue
		}
		id, err := strconv.Atoi(matched[0][1])
		if err != nil {
			panic(err)
		}
		games := strings.Split(matched[0][2], "; ")

		for _, game := range games {
			cubes := strings.Split(game, ", ")
			for _, cube := range cubes {
				rolls := strings.Split(cube, " ")
				quantity, err := strconv.Atoi(rolls[0])
				if err != nil {
					panic(err)
				}
				color := rolls[1]

				if bag[color] < quantity {
					ok = false
				}
			}
		}
		if ok {
			sum += id
		}
	}
	return sum
}
