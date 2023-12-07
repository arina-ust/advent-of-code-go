package day7

import (
	"advent-of-code-go/util"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

const day = "day7"

var inputFile string

func Solve(easy bool) (name string, res int, err error) {
	name = day
	setInput(easy)
	lines, err := util.ReadStringList(inputFile)
	if err != nil {
		return
	}

	res, err = partOne(lines)

	return

}

func setInput(easy bool) {
	if easy {
		inputFile = day + "/" + util.InputFileEasy
	} else {
		inputFile = day + "/" + util.InputFileFull
	}
}

type hand struct {
	cards        string
	bid          int
	rank         int
	typeStrength int
}

func (h *hand) getStrength() int {
	m := map[rune]int{}
	for _, card := range h.cards {
		m[card] += 1
	}
	if len(m) == 5 { // hand size
		if m['J'] != 0 {
			return 2 // one pair
		}
		return 1 // high card
	}

	for _, v := range m {
		if v == 5 {
			return 7 // five of kind
		} else if v == 4 {
			if m['J'] != 0 {
				return 7 // five of kind
			}
			return 6 // four of kind
		} else if v == 3 {
			if len(m) == 2 {
				if m['J'] != 0 {
					return 7 // five of kind
				}
				return 5 // full house
			}
			if len(m) == 3 {
				if m['J'] != 0 {
					return 4 + 1 + m['J']// four of kind
				}
				return 4 // three of kind
			}
		} else if v == 2 {
			if len(m) == 3 {
				if m['J'] != 0 {
					return 3 + 1 + m['J'] // full house OR four of a kind
				}
				return 3 // two pair
			} else if len(m) == 4 {
				if m['J'] != 0 {
					return 2 + m['J'] // three of a kind OR two pair
				}
				return 2 // one pair
			}
		}
	}
	return 0 // error
}

var cardStrengths = map[rune]int{'A': 13, 'K': 12, 'Q': 11, 'J': 0, 'T': 9, '9': 8,
	'8': 7, '7': 6, '6': 5, '5': 4, '4': 3, '3': 2, '2': 1,}

func (h *hand) compareLess(other *hand) bool {
	if h.typeStrength > other.typeStrength {
		return false
	} else if h.typeStrength < other.typeStrength {
		return true
	}
	// equal strength, compare by card
	for i := 0; i < 5; i++ {
		rh := rune(h.cards[i])
		ro := rune(other.cards[i])
		if cardStrengths[rh] > cardStrengths[ro] {
//			fmt.Printf("rh %s turned out bigger than ro %s with values %v %v\n", string(rh), string(ro), cardStrengths[rh], cardStrengths[ro])
			return false
		} else if cardStrengths[rh] < cardStrengths[ro] {
//			fmt.Printf("rh %s turned out smaller than ro %s with values %v %v\n", string(rh), string(ro), cardStrengths[rh], cardStrengths[ro])
			return true
		}
//		fmt.Printf("rh %s was the same as ro %s with values %v %v\n", string(rh), string(ro), cardStrengths[rh], cardStrengths[ro])
	}
	fmt.Println("shouldn't be here")
	return false
}

func partOne(lines []string) (int, error) {
	hands := make([]*hand, len(lines))
	for i, line := range lines {
		input := strings.Split(line, " ")
		bid, _ := strconv.Atoi(input[1])
		hand := &hand{
			cards: input[0],
			bid:   bid,
		}
		strgth := hand.getStrength()
		hand.typeStrength = strgth
		hands[i] = hand
	}
	
	sort.Slice(hands, func(i, j int) bool {
		return hands[i].compareLess(hands[j])
	})
	
	res := 0
	for j, h := range hands {
		h.rank = j+1
		res += h.rank * h.bid
	}
	
	return res, nil
}