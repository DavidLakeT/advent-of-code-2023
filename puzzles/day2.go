package puzzles

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve_part_1() int {
	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	lines, err := readFile("puzzles/inputs/day2.txt")
	if err != nil {
		return 0
	}

	sum := 0

gameLoop:
	for index, line := range lines {
		_, roundInfo, _ := strings.Cut(line, ":")
		for _, round := range strings.Split(roundInfo, ";") {
			for _, colorInfo := range strings.Split(round, ",") {
				amount, color, _ := strings.Cut(strings.TrimPrefix(colorInfo, " "), " ")
				integerAmount, err := strconv.Atoi(amount)
				if err != nil {
					return 0
				}

				if integerAmount > limits[color] {
					continue gameLoop
				}
			}
		}
		sum += (index + 1)
	}

	return sum
}

func Solve_part_2() int {
	lines, err := readFile("puzzles/inputs/day2.txt")
	if err != nil {
		return 0
	}

	sum := 0

	for _, line := range lines {
		limits := map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		_, roundInfo, _ := strings.Cut(line, ":")
		for _, round := range strings.Split(roundInfo, ";") {
			for _, colorInfo := range strings.Split(round, ",") {
				amount, color, _ := strings.Cut(strings.TrimPrefix(colorInfo, " "), " ")
				integerAmount, err := strconv.Atoi(amount)
				if err != nil {
					return 0
				}

				if integerAmount > limits[color] {
					limits[color] = integerAmount
				}
			}
		}
		power := limits["red"] * limits["green"] * limits["blue"]
		sum += power
	}

	return sum
}

func readFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}
