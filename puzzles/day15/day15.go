package puzzles

import (
	"os"
	"strconv"
	"strings"
)

func Solve_part_1() int {
	input, err := readFile("puzzles/inputs/day15.txt")
	if err != nil {
		return -1
	}

	sum := 0
	for _, step := range input {
		sum += getHashValue(step)
	}

	return sum
}

func Solve_part_2() int {
	input, err := readFile("puzzles/inputs/day15.txt")
	if err != nil {
		return -1
	}

	boxes := make([][]string, 256)
	focalLengths := make(map[string]int)

	for _, step := range input {
		if strings.Contains(step, "-") {
			label := step[:len(step)-1]
			index := getHashValue(label)

			for i, value := range boxes[index] {
				if value == label {
					boxes[index] = append(boxes[index][:i], boxes[index][i+1:]...)
					break
				}
			}

			continue
		}

		lensLabel, length, _ := strings.Cut(step, "=")
		lengthValue, _ := strconv.Atoi(length)

		index := getHashValue(lensLabel)
		contains := false
		for _, value := range boxes[index] {
			if value == lensLabel {
				contains = true
				break
			}
		}

		if !contains {
			boxes[index] = append(boxes[index], lensLabel)
		}

		focalLengths[lensLabel] = lengthValue
	}

	sum := 0

	for boxIndex, box := range boxes {
		for boxSlotIndex, boxSlot := range box {
			sum += (boxIndex + 1) * (boxSlotIndex + 1) * focalLengths[boxSlot]
		}
	}

	return sum
}

func getHashValue(step string) int {
	currentValue := 0

	for _, character := range step {
		currentValue += int(character)
		currentValue *= 17
		currentValue %= 256
	}

	return currentValue
}

func readFile(filepath string) ([]string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(file), ","), nil
}
