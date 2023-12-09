package puzzles

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
	numKinds
)

var hands = []*hand{}

type hand struct {
	cards string
	bid   int
	kind  int
}

func Solve_part_1() int {
	lines, err := readFile("puzzles/inputs/day7.txt")
	if err != nil {
		return -1
	}

	cardOrder := []rune{'A', 'K', 'Q', 'J', 'T', '9', '8', '7', '6', '5', '4', '3', '2'}

	parseInput(lines, false)
	sortHands(cardOrder)

	sum := 0

	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	return sum
}

func Solve_part_2() int {
	lines, err := readFile("puzzles/inputs/day7.txt")
	if err != nil {
		return -1
	}

	cardOrder := []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

	parseInput(lines, true)
	sortHands(cardOrder)

	sum := 0

	for i, hand := range hands {
		sum += (i + 1) * hand.bid
	}

	return sum
}

func sortHands(cardOrder []rune) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].kind == hands[j].kind {
			for k := range hands[i].cards {
				if hands[i].cards[k] == hands[j].cards[k] {
					continue
				}
				return indexOf(cardOrder, rune(hands[i].cards[k])) > indexOf(cardOrder, rune(hands[j].cards[k]))
			}
		}
		return hands[i].kind < hands[j].kind
	})
}

func getKind(cards string) int {
	matches := getHandMatches(cards)

	if len(matches) == 1 {
		return fiveOfAKind
	}

	if len(matches) == 5 {
		return highCard
	}

	if len(matches) == 2 {
		for _, character := range matches {
			if character == 4 {
				return fourOfAKind
			}
		}

		return fullHouse
	}

	if len(matches) == 3 {
		for _, character := range matches {
			if character == 3 {
				return threeOfAKind
			}
		}
	}

	pairs := []int{1, 2}
	count := 0

	for _, n := range matches {
		if !contains(pairs, n) {
			continue
		}
		if n == 2 {
			count++
		}
	}

	if count == 2 {
		return twoPair
	}

	return onePair
}

func getHandMatches(cards string) map[rune]int {
	matches := make(map[rune]int)

	for _, r := range cards {
		_, ok := matches[r]
		if !ok {
			matches[r] = strings.Count(cards, string(r))
		}
	}

	return matches
}

func matchJoker(hand *hand) {
	matches := getHandMatches(hand.cards)

	switch hand.kind {
	case fourOfAKind:
		hand.kind = fiveOfAKind
	case fullHouse:
		if matches['J'] == 2 || matches['J'] == 3 {
			hand.kind = fiveOfAKind
		}
	case threeOfAKind:
		if matches['J'] == 1 || matches['J'] == 3 {
			hand.kind = fourOfAKind
		}
	case twoPair:
		if matches['J'] == 1 {
			hand.kind = fullHouse
		} else if matches['J'] == 2 {
			hand.kind = fourOfAKind
		}
	case onePair:
		if matches['J'] == 1 || matches['J'] == 2 {
			hand.kind = threeOfAKind
		}
	case highCard:
		hand.kind = onePair
	}
}

func containsChar(input string, targetChar byte) bool {
	for i := 0; i < len(input); i++ {
		if input[i] == targetChar {
			return true
		}
	}
	return false
}

func indexOf(slice []rune, value rune) int {
	for i, v := range slice {
		if v == value {
			return i
		}
	}
	return -1
}

func contains[T comparable](elems []T, value T) bool {
	for _, elem := range elems {
		if value == elem {
			return true
		}
	}
	return false
}

func parseInput(lines []string, joker bool) {
	for _, line := range lines {
		cards, bid, _ := strings.Cut(line, " ")
		bidNumber, _ := strconv.Atoi(bid)

		newHand := &hand{
			cards: cards,
			bid:   bidNumber,
			kind:  getKind(cards),
		}

		if joker {
			if containsChar(newHand.cards, 'J') {
				matchJoker(newHand)
			}
		}
		hands = append(hands, newHand)
	}
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
