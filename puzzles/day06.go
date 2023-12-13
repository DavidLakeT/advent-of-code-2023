package puzzles

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Solve_part_1() int {
	lines, err := readFile("puzzles/inputs/day6.txt")
	if err != nil {
		return -1
	}

	times := strings.Fields(lines[0][12:])
	distances := strings.Fields(lines[1][12:])

	sum := 1

	for i := 0; i < len(times); i++ {
		recordBeats := 0

		timeNumber, _ := strconv.Atoi(times[i])
		recordDistanceNumber, _ := strconv.Atoi(distances[i])
		for j := 1; j < timeNumber; j++ {
			remainingTime := timeNumber - j
			traveledDistance := j * remainingTime

			if traveledDistance > recordDistanceNumber {
				recordBeats++
			}
		}

		sum *= recordBeats
	}

	return sum
}

func Solve_part_2() int {
	lines, err := readFile("puzzles/inputs/day6.txt")
	if err != nil {
		return -1
	}

	time, _ := strconv.Atoi(strings.ReplaceAll(lines[0][12:], " ", ""))
	distance, _ := strconv.Atoi(strings.ReplaceAll(lines[1][12:], " ", ""))

	recordBeats := 0

	for i := 0; i < time; i++ {
		remainingTime := time - i
		traveledDistance := i * remainingTime

		if traveledDistance > distance {
			recordBeats++
		}
	}

	return recordBeats
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
