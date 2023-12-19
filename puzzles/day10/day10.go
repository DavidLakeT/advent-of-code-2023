package puzzles

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

const (
	north int = iota
	east
	south
	west
)

type Position struct {
	rowPosition int
	colPosition int
}

func Solve_part_1() int {
	matrix, err := readFile("puzzles/inputs/day10.txt")
	if err != nil {
		return -1
	}

	startingPosition, err := searchStartingPoint(&matrix)
	if err != nil {
		return -1
	}

	sum := 1

	for i := 0; i < 3; i++ {
		if isValidDirection(&matrix, startingPosition, i) {
			direction := i
			newPosition := getNewPositionByDirection(startingPosition, direction)
			entryValue := matrix[newPosition.rowPosition][newPosition.colPosition]

			for entryValue != "S" {
				newPosition = getNewPositionByDirection(newPosition, getNewDirectionByEntry(entryValue, direction))
				direction = getNewDirectionByEntry(entryValue, direction)
				entryValue = matrix[newPosition.rowPosition][newPosition.colPosition]
				sum += 1
			}

			return sum / 2
		}
	}

	return -1
}

func getNewPositionByDirection(currentPosition *Position, direction int) *Position {
	newPosition := &Position{
		rowPosition: currentPosition.rowPosition,
		colPosition: currentPosition.colPosition,
	}

	switch direction {
	case north:
		newPosition.rowPosition -= 1
	case east:
		newPosition.colPosition += 1
	case south:
		newPosition.rowPosition += 1
	case west:
		newPosition.colPosition -= 1
	}

	return newPosition
}

func getNewDirectionByEntry(entryValue string, currentDirection int) int {
	switch entryValue {
	case "|":
		if currentDirection == north {
			return north
		}

		if currentDirection == south {
			return south
		}
	case "-":
		if currentDirection == east {
			return east
		}

		if currentDirection == west {
			return west
		}
	case "F":
		if currentDirection == north {
			return east
		}

		if currentDirection == west {
			return south
		}
	case "7":
		if currentDirection == north {
			return west
		}

		if currentDirection == east {
			return south
		}
	case "L":
		if currentDirection == south {
			return east
		}

		if currentDirection == west {
			return north
		}
	case "J":
		if currentDirection == south {
			return west
		}

		if currentDirection == east {
			return north
		}
	}

	return -1
}

func isValidDirection(matrix *[][]string, position *Position, direction int) bool {
	newPosition := getNewPositionByDirection(position, direction)
	entryValue := (*matrix)[newPosition.rowPosition][newPosition.colPosition]

	switch direction {
	case north:
		return entryValue == "|" || entryValue == "7" || entryValue == "F"
	case east:
		return entryValue == "-" || entryValue == "J" || entryValue == "7"
	case south:
		return entryValue == "|" || entryValue == "J" || entryValue == "L"
	case west:
		return entryValue == "-" || entryValue == "F" || entryValue == "L"
	}

	return false
}

func searchStartingPoint(matrix *[][]string) (*Position, error) {
	for i, row := range *matrix {
		for j, column := range row {
			if column == "S" {
				position := &Position{
					rowPosition: i,
					colPosition: j,
				}

				return position, nil
			}
		}
	}

	return nil, errors.New("not found")
}

func readFile(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var matrix [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		chars := strings.Split(scanner.Text(), "")

		matrix = append(matrix, chars)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}
