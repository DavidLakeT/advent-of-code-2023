package puzzles

import (
	"bytes"
	"os"
)

const (
	north int = iota
	east
	south
	west
)

func Solve_part_1() int {
	lines, err := readFile("puzzles/inputs/day14.txt")
	if err != nil {
		return -1
	}

	tiltPlatform(&lines, north)

	return getTotalLoad(&lines)
}

func Solve_part_2() int {
	lines, err := readFile("puzzles/inputs/day14.txt")
	if err != nil {
		return -1
	}

	reachedPatterns := map[string]int{}
	len := 1000000000
	for i := 0; i < len; i++ {
		tiltPlatform(&lines, north)
		tiltPlatform(&lines, west)
		tiltPlatform(&lines, south)
		tiltPlatform(&lines, east)

		pattern := bytes.Join(lines, []byte{})
		if _, exists := reachedPatterns[string(pattern)]; exists {
			i = len - (len-i)%(i-reachedPatterns[string(pattern)])
		}

		reachedPatterns[string(pattern)] = i
	}

	return getTotalLoad(&lines)
}

func tiltPlatform(p *[][]byte, direction int) {
	switch direction {
	case north:
		for x := 0; x < len((*p)[0]); x++ {
			available := 0
			for y := 0; y < len(*p); y++ {
				switch (*p)[y][x] {
				case '#':
					available = y + 1
				case 'O':
					if available < y {
						(*p)[available][x] = 'O'
						(*p)[y][x] = '.'
					}
					available++
				}
			}
		}
	case west:
		for y := 0; y < len(*p); y++ {
			space := 0
			for x := 0; x < len((*p)[y]); x++ {
				switch (*p)[y][x] {
				case '#':
					space = x + 1
				case 'O':
					if space < x {
						(*p)[y][space] = 'O'
						(*p)[y][x] = '.'
					}
					space++
				}
			}
		}
	case south:
		for x := 0; x < len((*p)[0]); x++ {
			space := len(*p) - 1
			for y := len(*p) - 1; y >= 0; y-- {
				switch (*p)[y][x] {
				case '#':
					space = y - 1
				case 'O':
					if space > y {
						(*p)[space][x] = 'O'
						(*p)[y][x] = '.'
					}
					space--
				}
			}
		}
	case east:
		for y := 0; y < len(*p); y++ {
			available := len((*p)[y]) - 1
			for x := len((*p)[y]) - 1; x >= 0; x-- {
				switch (*p)[y][x] {
				case '#':
					available = x - 1
				case 'O':
					if available > x {
						(*p)[y][available] = 'O'
						(*p)[y][x] = '.'
					}
					available--
				}
			}
		}
	}
}

func getTotalLoad(platform *[][]byte) int {
	sum := 0
	for x := 0; x < len(*platform); x++ {
		for y := 0; y < len((*platform)[x]); y++ {
			if (*platform)[y][x] == 'O' {
				sum += len((*platform)) - int(y)
			}
		}
	}

	return sum
}

func readFile(filepath string) ([][]byte, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return bytes.Fields([]byte(file)), nil
}
