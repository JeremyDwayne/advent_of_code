package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	bag := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		ok := true
		matcher := regexp.MustCompile(`Game (\d+): (.*)`)
		matched := matcher.FindAllStringSubmatch(line, -1)[0]

		id, err := strconv.Atoi(matched[1])
		if err != nil {
			panic(err)
		}
		games := strings.Split(matched[2], "; ")

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
	fmt.Println(sum)
}
