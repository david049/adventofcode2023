package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	cards       string
	bid         int
	designation int
	numJacks    int
}

func lookUp(char rune) int {
	index := int(char - '0')
	if char == 'T' {
		index = 10
	}
	if char == 'J' {
		index = -1
	}
	if char == 'Q' {
		index = 12
	}
	if char == 'K' {
		index = 13
	}
	if char == 'A' {
		index = 14
	}
	return index
}

func charMap(input string) (map[int]int, int) {
	cMap := make(map[int]int)
	maxCards := -1
	for _, char := range input {
		cMap[lookUp(char)]++
		if cMap[lookUp(char)] > maxCards && lookUp(char) != -1 {
			maxCards = cMap[lookUp(char)]
		}
	}
	return cMap, maxCards
}

func compareCards(card1 string, card2 string) int {

	for i := 0; i < 5; i++ {
		if lookUp(rune(card1[i])) == lookUp(rune(card2[i])) {
			continue
		}
		if lookUp(rune(card1[i])) > lookUp(rune(card2[i])) {
			return 0
		}
		return 1
	}
	return 0
}

func transformWithJ(hand Hand) Hand {
	// what happens to 5
	if hand.numJacks == 5 {
		return Hand{cards: hand.cards, bid: hand.bid, designation: 7, numJacks: hand.numJacks}
	}
	if hand.numJacks == 4 {
		return Hand{cards: hand.cards, bid: hand.bid, designation: 7, numJacks: hand.numJacks}
	}
	if hand.designation == 7 {
		return Hand{cards: hand.cards, bid: hand.bid, designation: hand.designation, numJacks: hand.numJacks}
	}
	// what happens to 4
	if hand.designation == 6 {
		if hand.numJacks == 1 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 7, numJacks: hand.numJacks}
		}
		return Hand{cards: hand.cards, bid: hand.bid, designation: hand.designation, numJacks: hand.numJacks}
	}
	// what happens to full house
	if hand.designation == 5 {
		if hand.numJacks == 2 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 7, numJacks: hand.numJacks}
		}
		return Hand{cards: hand.cards, bid: hand.bid, designation: hand.designation, numJacks: hand.numJacks}
	}
	// what happens to 3
	if hand.designation == 4 {
		if hand.numJacks == 1 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 6, numJacks: hand.numJacks}
		}
		if hand.numJacks == 2 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 7, numJacks: hand.numJacks}
		}
		return Hand{cards: hand.cards, bid: hand.bid, designation: hand.designation, numJacks: hand.numJacks}
	}
	// what happens to 2 pair
	if hand.designation == 3 {
		if hand.numJacks == 1 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 5, numJacks: hand.numJacks}
		}
		if hand.numJacks == 2 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 6, numJacks: hand.numJacks}
		}
		return Hand{cards: hand.cards, bid: hand.bid, designation: hand.designation, numJacks: hand.numJacks}
	}
	// what happens to 1 pair
	if hand.designation == 2 {
		if hand.numJacks == 1 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 4, numJacks: hand.numJacks}
		}
		if hand.numJacks == 2 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 6, numJacks: hand.numJacks}
		}
		if hand.numJacks == 3 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 7, numJacks: hand.numJacks}
		}
		return Hand{cards: hand.cards, bid: hand.bid, designation: hand.designation, numJacks: hand.numJacks}
	}
	// what happens to high
	if hand.designation == 1 {
		if hand.numJacks == 1 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 2, numJacks: hand.numJacks}
		}
		if hand.numJacks == 2 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 4, numJacks: hand.numJacks}
		}
		if hand.numJacks == 3 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 6, numJacks: hand.numJacks}
		}
		if hand.numJacks == 4 {
			return Hand{cards: hand.cards, bid: hand.bid, designation: 7, numJacks: hand.numJacks}
		}
		return Hand{cards: hand.cards, bid: hand.bid, designation: hand.designation, numJacks: hand.numJacks}
	}
	return Hand{}
}

func main() {
	inputstring, _ := os.ReadFile("input.txt")
	lines := strings.Split(string(inputstring), "\n")
	var hands []Hand
	for _, line := range lines {
		cardBid := strings.Fields(line)
		bid, _ := strconv.Atoi(cardBid[1])
		currentHand := Hand{cards: cardBid[0], bid: bid}
		cMap, maxCard := charMap(cardBid[0])
		currentHand.numJacks = cMap[-1]
		if maxCard == 5 {
			currentHand.designation = 7 // 5
		}
		if maxCard == 4 {
			currentHand.designation = 6 // 4
		}
		if maxCard == 3 {
			currentHand.designation = 4 // three
			for key, val := range cMap {
				if val == 2 && key != -1 {
					currentHand.designation = 5 // fullhouse
				}
			}
		}
		if maxCard == 2 {
			var twos []int
			for key, val := range cMap {
				if val == 2 && key != -1 {
					twos = append(twos, key)
				}
			}
			if len(twos) == 2 {
				currentHand.designation = 3 // twopair
			} else {
				currentHand.designation = 2 // onepair
			}
		}
		if maxCard == 1 {
			currentHand.designation = 1 // highcard
		}
		hands = append(hands, currentHand)
	}
	sort.SliceStable(hands, func(i, j int) bool {
		// transform
		hand1 := transformWithJ(hands[i])
		hand2 := transformWithJ(hands[j])
		if hand1.designation > hand2.designation {
			return false
		}
		if hand1.designation < hand2.designation {
			return true
		}
		if hand1.designation == hand2.designation {
			if compareCards(hand1.cards, hand2.cards) == 0 {
				return false
			} else {
				return true
			}
		}
		return true
	})
	sumBids := 0
	rank := 1
	for _, hand := range hands {
		sumBids += rank * hand.bid
		rank++
	}
	fmt.Println(sumBids)
}
