package puzzles

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve_part_1() int {
	lines, err := readFile("puzzles/inputs/day9.txt")
	if err != nil {
		return -1
	}

	sum := 0

	for _, line := range lines {
		numbers := parseLine(line)
		sum += processNumbers(numbers, false)
	}

	return sum
}

func Solve_part_2() int {
	lines, err := readFile("puzzles/inputs/day9.txt")
	if err != nil {
		return -1
	}

	sum := 0

	for _, line := range lines {
		numbers := parseLine(line)
		sum += processNumbers(numbers, true)
	}

	return sum
}

func processNumbers(numbers []int, reverse bool) int {
	if checkZeros(numbers) {
		return 0
	}

	if reverse {
		return numbers[0] - processNumbers(calcDifferences(numbers, reverse), reverse)
	} else {
		return processNumbers(calcDifferences(numbers, reverse), reverse) + numbers[len(numbers)-1]
	}
}

func calcDifferences(numbers []int, reverse bool) []int {
	var differences []int

	for i := 1; i < len(numbers); i++ {
		differences = append(differences, numbers[i]-numbers[i-1])
	}

	return differences
}

func checkZeros(numbers []int) bool {
	for _, number := range numbers {
		if number != 0 {
			return false
		}
	}

	return true
}

func parseLine(line string) []int {
	var numbers []int
	splittedNumbers := strings.Fields(line)

	for _, field := range splittedNumbers {
		number, _ := strconv.Atoi(field)
		numbers = append(numbers, number)
	}

	return numbers
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
