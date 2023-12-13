package puzzles

import (
	"bufio"
	"math"
	"os"
	"regexp"
	"strconv"
)

func Solve_part_1() int {
	lines, err := readFile("puzzles/inputs/day3.txt")
	if err != nil {
		return -1
	}

	numberRegex := regexp.MustCompile(`\d+`)
	symbolRegex := regexp.MustCompile(`[^\w\s.]`)

	sum := 0

	for index, line := range lines {
		numberMatches := numberRegex.FindAllStringIndex(line, -1)

		for _, match := range numberMatches {
			startIndex, endIndex := match[0], match[1]

			for i := -1; i <= 1; i++ {
				if (index+i) < 0 || (index+i) >= (len(lines)) {
					continue
				}

				adjacentText := lines[index+i]
				if len(symbolRegex.FindAllString(adjacentText[max(startIndex-1, 0):min(endIndex+1, len(line))], -1)) != 0 {
					number, err := strconv.Atoi(line[startIndex:endIndex])
					if err != nil {
						return -1
					}

					sum += number
				}
			}
		}
	}

	return sum
}

func Solve_part_2() int {
	lines, err := readFile("puzzles/inputs/day3.txt")
	if err != nil {
		return -1
	}

	numberRegex := regexp.MustCompile(`\d+`)

	sum := 0

	for index, line := range lines {
		for i, char := range line {
			if char == '*' {
				var adjacentNumbers []int
				for j := -1; j <= 1; j++ {
					if (index+j) < 0 || (index+j) >= (len(lines)) {
						continue
					}

					adjacentText := lines[index+j]
					numberMatches := numberRegex.FindAllStringIndex(adjacentText, -1)
					if len(numberMatches) != 0 {
						for _, match := range numberMatches {
							if math.Abs(float64(match[0]-i)) <= 1 || math.Abs(float64(match[1]-i-1)) <= 1 {
								number, err := strconv.Atoi(adjacentText[match[0]:match[1]])
								if err != nil {
									return -1
								}

								adjacentNumbers = append(adjacentNumbers, number)
							}
						}
					}
				}

				if len(adjacentNumbers) == 2 {
					sum += (adjacentNumbers[0] * adjacentNumbers[1])
				}
			}
		}
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
