package puzzles

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func Solve_part_1() int {
	input, err := readFile("puzzles/inputs/day5.txt")
	if err != nil {
		return 0
	}

	sections := strings.Split(input, "\n\n")
	seeds := strings.Fields(sections[0][7:])
	totalMin := int(^uint(0) >> 1)

	for _, seed := range seeds {
		seedNumber, _ := strconv.Atoi(seed)

	sectionLoop:
		for _, section := range sections[1:] {
			lines := strings.Split(section, "\n")[1:]

			for _, line := range lines {
				lineInfo := strings.Fields(line)

				endPosition, _ := strconv.Atoi(lineInfo[0])
				startPosition, _ := strconv.Atoi(lineInfo[1])
				rangeLength, _ := strconv.Atoi(lineInfo[2])

				if seedNumber >= startPosition && seedNumber < (startPosition+rangeLength) {
					seedNumber = endPosition + (seedNumber - startPosition)
					continue sectionLoop
				}
			}
		}

		if seedNumber < totalMin {
			totalMin = seedNumber
		}
	}

	return totalMin
}

func Solve_part_2() int {
	input, err := readFile("puzzles/inputs/day5.txt")
	if err != nil {
		return 0
	}

	sections := strings.Split(input, "\n\n")
	seeds := strings.Fields(sections[0][7:])
	totalMin := math.MaxInt64

	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < len(seeds); i += 2 {
		seed, _ := strconv.Atoi(seeds[i])
		iterations, _ := strconv.Atoi(seeds[i+1])

		wg.Add(1)
		go func() {
			defer wg.Done()
			convertSeeds(
				iterations,
				seed,
				sections[1:],
				&totalMin,
				&mutex,
			)
		}()
	}

	wg.Wait()

	return totalMin
}

func convertSeeds(iterations int, seedNumber int, sections []string, totalMin *int, mutex *sync.Mutex) {
	localMinimum := math.MaxInt64
	for i := 0; i < iterations; i++ {
	sectionLoop:
		for _, section := range sections[1:] {
			lines := strings.Split(section, "\n")[1:]

			for _, line := range lines {
				lineInfo := strings.Fields(line)

				endPosition, _ := strconv.Atoi(lineInfo[0])
				startPosition, _ := strconv.Atoi(lineInfo[1])
				rangeLength, _ := strconv.Atoi(lineInfo[2])

				if seedNumber >= startPosition && seedNumber < (startPosition+rangeLength) {
					seedNumber = endPosition + (seedNumber - startPosition)
					continue sectionLoop
				}
			}
		}
	}

	mutex.Lock()
	defer mutex.Unlock()
	if localMinimum < *totalMin {
		*totalMin = localMinimum
	}
}

func readFile(filePath string) (string, error) {
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return "", err
	}

	return string(content), nil
}
