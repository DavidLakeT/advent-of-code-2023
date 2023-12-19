package puzzles

import (
	"bufio"
	"os"
)

type Point struct {
	x int
	y int
}

type Location struct {
	initPoint   Point
	targetPoint Point
}

var directionsMap = map[Point]int{{1, 0}: 1, {0, 1}: 2, {-1, 0}: 4, {0, -1}: 8}

func Solve_part_1() int {
	input, err := readFile("puzzles/inputs/day16.txt")
	if err != nil {
		return -1
	}

	initPoint := Point{-1, 0}
	targetPoint := Point{1, 0}

	return traverse(input, initPoint, targetPoint)
}

func Solve_part_2() int {
	input, err := readFile("puzzles/inputs/day16.txt")
	if err != nil {
		return -1
	}

	sum := 0
	for row := range input {
		sum = max(sum, traverse(input, Point{-1, row}, Point{1, 0}))
		sum = max(sum, traverse(input, Point{len(input[0]), row}, Point{-1, 0}))
	}

	for column := range input[0] {
		sum = max(sum, traverse(input, Point{column, -1}, Point{0, 1}))
		sum = max(sum, traverse(input, Point{column, len(input)}, Point{0, -1}))
	}

	return sum
}

func traverse(grid []string, currentPosition, targetDirection Point) int {
	energized := make(map[Point]int, len(grid)*len(grid[0]))
	testTile := Point{currentPosition.x + targetDirection.x, currentPosition.y + targetDirection.y}
	pendingLocations := []Location{{testTile, targetDirection}}

	for len(pendingLocations) > 0 {
		testTile, targetDirection = pendingLocations[0].initPoint, pendingLocations[0].targetPoint
		pendingLocations = pendingLocations[1:]

	traverseLoop:
		for (energized[testTile]&directionsMap[targetDirection]) == 0 && isWithinLimits(grid, testTile) {
			energized[testTile] += directionsMap[targetDirection]

			switch grid[testTile.y][testTile.x] {
			case '-':
				switch targetDirection {
				case Point{0, 1}, Point{0, -1}:
					pendingLocations = append(pendingLocations, Location{testTile, Point{1, 0}})
					pendingLocations = append(pendingLocations, Location{testTile, Point{-1, 0}})
					break traverseLoop
				}
			case '/':
				switch targetDirection {
				case Point{1, 0}:
					targetDirection = Point{0, -1}
				case Point{-1, 0}:
					targetDirection = Point{0, 1}
				case Point{0, 1}:
					targetDirection = Point{-1, 0}
				case Point{0, -1}:
					targetDirection = Point{1, 0}
				}
			case '\\':
				switch targetDirection {
				case Point{1, 0}:
					targetDirection = Point{0, 1}
				case Point{-1, 0}:
					targetDirection = Point{0, -1}
				case Point{0, 1}:
					targetDirection = Point{1, 0}
				case Point{0, -1}:
					targetDirection = Point{-1, 0}
				}
			case '|':
				switch targetDirection {
				case Point{1, 0}, Point{-1, 0}:
					pendingLocations = append(pendingLocations, Location{testTile, Point{0, 1}})
					pendingLocations = append(pendingLocations, Location{testTile, Point{0, -1}})
					break traverseLoop
				}
			}

			testTile = Point{testTile.x + targetDirection.x, testTile.y + targetDirection.y}
		}
	}

	return len(energized)
}

func isWithinLimits(grid []string, pos Point) bool {
	return pos.x >= 0 && pos.x < len(grid[0]) && pos.y >= 0 && pos.y < len(grid)
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
