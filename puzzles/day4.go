package puzzles

import (
	"bufio"
	"math"
	"os"
	"strings"
)

func Solve_part_1() int {
	lines, err := readFile("puzzles/inputs/day4.txt")
	if err != nil {
		return -1
	}

	sum := 0

	for _, line := range lines {
		winningNumbersMap := make(map[string]bool)
		matches := 0

		cardInfo := line[10:]
		winningNumbersInfo, purchasedNumbersInfo, _ := strings.Cut(cardInfo, "|")

		winningNumbers := strings.Fields(winningNumbersInfo)
		for _, number := range winningNumbers {
			winningNumbersMap[number] = true
		}

		purchasedNumbers := strings.Fields(purchasedNumbersInfo)
		for _, number := range purchasedNumbers {
			if winningNumbersMap[number] {
				matches += 1
			}
		}

		if matches > 0 {
			sum += int(math.Pow(2, float64(matches-1)))
		}
	}

	return sum
}

func Solve_part_2() int {
	lines, err := readFile("puzzles/inputs/day4.txt")
	if err != nil {
		return -1
	}

	cardAmounts := make(map[int]int)
	sum := 0

	for index, line := range lines {
		cardAmounts[index] = cardAmounts[index] + 1
		winningNumbersMap := make(map[string]bool)

		cardInfo := line[10:]
		winningNumbersInfo, purchasedNumbersInfo, _ := strings.Cut(cardInfo, "|")

		winningNumbers := strings.Fields(winningNumbersInfo)
		for _, number := range winningNumbers {
			winningNumbersMap[number] = true
		}

		purchasedNumbers := strings.Fields(purchasedNumbersInfo)
		matches := 0
		for _, number := range purchasedNumbers {
			if winningNumbersMap[number] {
				matches += 1
			}
		}

		for i := 1; i <= cardAmounts[index]; i++ {
			for j := 1; j <= matches; j++ {
				cardAmounts[index+j] = cardAmounts[index+j] + 1
			}
		}

		sum += cardAmounts[index]
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
