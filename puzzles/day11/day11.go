package puzzles

import (
	"bufio"
	"os"
	"strings"
)

type Galaxy struct {
	row    int
	column int
}

func Solve_part_1() int {
	matrix, err := readFile("puzzles/inputs/day11.txt")
	if err != nil {
		return -1
	}

	processMatrix(&matrix)

	var galaxies []*Galaxy
	searchGalaxies(&matrix, &galaxies)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		galaxy := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			secondGalaxy := galaxies[j]
			distance := 0

			min := minGalaxy(galaxy, secondGalaxy, "row")
			max := maxGalaxy(galaxy, secondGalaxy, "row")
			for m := min.row + 1; m <= max.row; m++ {
				if matrix[m][min.column] == "" {
					distance += 2
					continue
				}
				distance++
			}

			min = minGalaxy(galaxy, secondGalaxy, "column")
			max = maxGalaxy(galaxy, secondGalaxy, "column")
			for m := min.column + 1; m <= max.column; m++ {
				if matrix[min.row][m] == "" {
					distance += 2
					continue
				}
				distance++
			}

			sum += distance
		}
	}

	return sum
}

func Solve_part_2() int {
	matrix, err := readFile("puzzles/inputs/day11.txt")
	if err != nil {
		return -1
	}

	processMatrix(&matrix)

	var galaxies []*Galaxy
	searchGalaxies(&matrix, &galaxies)

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		galaxy := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			secondGalaxy := galaxies[j]
			distance := 0

			min := minGalaxy(galaxy, secondGalaxy, "row")
			max := maxGalaxy(galaxy, secondGalaxy, "row")
			for m := min.row + 1; m <= max.row; m++ {
				if matrix[m][min.column] == "" {
					distance += 1000000
					continue
				}
				distance++
			}

			min = minGalaxy(galaxy, secondGalaxy, "column")
			max = maxGalaxy(galaxy, secondGalaxy, "column")
			for m := min.column + 1; m <= max.column; m++ {
				if matrix[min.row][m] == "" {
					distance += 1000000
					continue
				}
				distance++
			}

			sum += distance
		}
	}

	return sum
}

func searchGalaxies(matrix *[][]string, galaxies *[]*Galaxy) {
	for rowIndex, row := range *matrix {
		for columnIndex, value := range row {
			if value == "#" {
				*galaxies = append(*galaxies, &Galaxy{row: rowIndex, column: columnIndex})
			}
		}
	}
}

func isEmptyRow(matrix *[][]string, row int) bool {
	for _, value := range (*matrix)[row] {
		if value == "#" {
			return false
		}
	}
	return true
}

func isEmptyColumn(matrix *[][]string, column int) bool {
	for _, row := range *matrix {
		if row[column] == "#" {
			return false
		}
	}
	return true
}

func appendEmptyRow(matrix *[][]string, index int) {
	(*matrix)[index] = make([]string, len((*matrix)[index]))
}

func appendEmptyColumn(matrix *[][]string, index int) {
	for i := range *matrix {
		(*matrix)[i][index] = ""
	}
}

func processMatrix(matrix *[][]string) {
	i := 0
	for i < len(*matrix) {
		if isEmptyRow(matrix, i) {
			appendEmptyRow(matrix, i)
		}
		i++
	}

	i = 0
	for i < len((*matrix)[0]) {
		if isEmptyColumn(matrix, i) {
			appendEmptyColumn(matrix, i)
		}
		i++
	}
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

func maxGalaxy(a, b *Galaxy, vector string) *Galaxy {
	if vector == "row" {
		if a.row > b.row {
			return a
		}
		return b
	}

	if vector == "column" {
		if a.column > b.column {
			return a
		}
		return b
	}

	return nil
}

func minGalaxy(a, b *Galaxy, vector string) *Galaxy {
	if vector == "row" {
		if a.row < b.row {
			return a
		}
		return b
	}

	if vector == "column" {
		if a.column < b.column {
			return a
		}
		return b
	}

	return nil
}
