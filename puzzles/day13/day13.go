package puzzles

import (
	"bytes"
	"os"
	"strings"
)

func Solve_part_1() int {
	patterns, err := readFile("puzzles/inputs/day13.txt")
	if err != nil {
		return -1
	}

	return processPatterns(patterns, false)
}

func Solve_part_2() int {
	patterns, err := readFile("puzzles/inputs/day13.txt")
	if err != nil {
		return -1
	}

	return processPatterns(patterns, true)
}

func processPatterns(patterns []string, smudge bool) int {
	sum := 0
	for _, pattern := range patterns {
		lines := strings.Split(pattern, "\n")
		sum += searchVerticalReflection(lines, smudge) + searchHorizontalReflection(lines, smudge)*100
	}

	return sum
}

func searchVerticalReflection(lines []string, smudge bool) int {
	width := len(lines[0])
	height := len(lines)

	for i := 0; i < width-1; i++ {
		diff := 0

		for j := 0; j < height; j++ {
			for offset := 0; ; offset++ {
				leftColumnIndex := i - offset
				rightColumnIndex := i + offset + 1

				if leftColumnIndex < 0 || rightColumnIndex >= width {
					break
				}

				if lines[j][leftColumnIndex] != lines[j][rightColumnIndex] {
					diff++
				}
			}
		}

		if (smudge && diff == 1) || (!smudge && diff == 0) {
			return i + 1
		}
	}

	return 0
}

func searchHorizontalReflection(lines []string, smudge bool) int {
	width := len(lines[0])
	height := len(lines)

	for i := 0; i < height-1; i++ {
		diff := 0

		for j := 0; j < width; j++ {
			for offset := 0; ; offset++ {
				aboveRowIndex := i - offset
				belowRowIndex := i + offset + 1

				if aboveRowIndex < 0 || belowRowIndex >= height {
					break
				}

				if lines[aboveRowIndex][j] != lines[belowRowIndex][j] {
					diff++
				}
			}
		}

		if (smudge && diff == 1) || (!smudge && diff == 0) {
			return i + 1
		}
	}

	return 0
}

func readFile(filepath string) ([]string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(bytes.TrimSpace(file)), "\n\n"), nil
}
