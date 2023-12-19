package puzzles

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

func Solve_part_1() int {
	lines, err := readFile("puzzles/inputs/day12.txt")
	if err != nil {
		return -1
	}

	totalArrangements := 0

	for _, line := range lines {
		textualInfo, numericalInfo, _ := strings.Cut(line, " ")
		textualRecords := []byte(textualInfo)
		numericalRecords := parseNumericalRecords(strings.Split(numericalInfo, ","))
		totalArrangements += getPossibleArrangements(textualRecords, numericalRecords)
	}

	return totalArrangements
}

func Solve_part_2() int {
	lines, err := readFile("puzzles/inputs/day12.txt")
	if err != nil {
		return -1
	}

	totalArrangements := 0

	for _, line := range lines {
		textualInfo, numericalInfo, _ := strings.Cut(line, " ")

		var tempTextual, tempNumerical string
		for i := 0; i < 5; i++ {
			tempTextual = tempTextual + textualInfo + "?"
			tempNumerical = tempNumerical + numericalInfo + ","
		}

		textualInfo = tempTextual[:len(tempTextual)-1]
		numericalInfo = tempNumerical[:len(tempNumerical)-1]

		textualRecords := []byte(textualInfo)
		numericalRecords := parseNumericalRecords(strings.Split(numericalInfo, ","))
		totalArrangements += getPossibleArrangements(textualRecords, numericalRecords)
	}

	return totalArrangements
}

func getPossibleArrangements(textualRecords []byte, numericalRecords []int) int {
	possibleArrangements := 0
	currentStates := map[[4]int]int{{0, 0, 0, 0}: 1}
	newStates := map[[4]int]int{}

	for len(currentStates) > 0 {
		for state, num := range currentStates {
			firstState, secondState, continuousCount, expectedDot := state[0], state[1], state[2], state[3]

			if firstState == len(textualRecords) {
				if secondState == len(numericalRecords) {
					possibleArrangements += num
				}

				continue
			}

			if (textualRecords[firstState] == '#' || textualRecords[firstState] == '?') && secondState < len(numericalRecords) && expectedDot == 0 {
				if textualRecords[firstState] == '?' && continuousCount == 0 {
					newStates[[4]int{firstState + 1, secondState, continuousCount, expectedDot}] += num
				}

				continuousCount++
				if continuousCount == numericalRecords[secondState] {
					secondState++
					expectedDot = 1
					continuousCount = 0
				}

				newStates[[4]int{firstState + 1, secondState, continuousCount, expectedDot}] += num
				continue
			}

			if (textualRecords[firstState] == '.' || textualRecords[firstState] == '?') && continuousCount == 0 {
				expectedDot = 0
				newStates[[4]int{firstState + 1, secondState, continuousCount, expectedDot}] += num
			}
		}

		currentStates, newStates = newStates, currentStates
		maps.Clear(newStates)
	}

	return possibleArrangements
}

func parseNumericalRecords(input []string) []int {
	output := make([]int, len(input))
	for index, record := range input {
		value, _ := strconv.Atoi(record)
		output[index] = value
	}

	return output
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
