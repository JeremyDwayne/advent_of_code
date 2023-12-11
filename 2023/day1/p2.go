package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	digits := map[string]int{
		"one": 1, "two": 2, "three": 3, "four": 4,
		"five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9,
	}
	sum := 0
	for scanner.Scan() {
		first := ""
		last := ""
		for _, char := range scanner.Text() {
			fmt.Println(string(char))
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
		fmt.Printf("%s + %s = %s\n", first, last, first+last)
	}
	fmt.Println("--------------------")
	fmt.Println("Sum: ", sum)
}
