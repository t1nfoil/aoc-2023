package main

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type card int

const (
	JOKER card = 1
	TWO   card = 2
	THREE card = 3
	FOUR  card = 4
	FIVE  card = 5
	SIX   card = 6
	SEVEN card = 7
	EIGHT card = 8
	NINE  card = 9
	TEN   card = 10
	QUEEN card = 12
	KING  card = 13
	ACE   card = 14
)

var convertToString = map[card]string{
	JOKER: "1",
	TWO:   "2",
	THREE: "3",
	FOUR:  "4",
	FIVE:  "5",
	SIX:   "6",
	SEVEN: "7",
	EIGHT: "8",
	NINE:  "9",
	TEN:   "T",
	QUEEN: "Q",
	KING:  "K",
	ACE:   "A",
}

var convertToCard = map[string]card{
	"J": JOKER,
	"2": TWO,
	"3": THREE,
	"4": FOUR,
	"5": FIVE,
	"6": SIX,
	"7": SEVEN,
	"8": EIGHT,
	"9": NINE,
	"T": TEN,
	"Q": QUEEN,
	"K": KING,
	"A": ACE,
}

type handType string

const (
	FIVEOAK   handType = "Five Of A Kind"
	FOUROAK   handType = "Four Of A Kind"
	FULLHOUSE handType = "Full House"
	THREEOAK  handType = "Three Of A Kind"
	TWOPAIR   handType = "Two Pair"
	ONEPAIR   handType = "One Pair"
	HIGHCARD  handType = "High Card"
)

var handOrder = map[handType]int{
	FIVEOAK:   7,
	FOUROAK:   6,
	FULLHOUSE: 5,
	THREEOAK:  4,
	TWOPAIR:   3,
	ONEPAIR:   2,
	HIGHCARD:  1,
}

type hand struct {
	holding           string
	fiveOak           card
	fourOak           card
	threeOak          card
	twoPair           card
	onePair           card
	highCard          card
	secondHighestCard card

	hasJoker      bool
	camelHandType handType

	bid int
}

func (h *hand) determineHand(holding string) {
	h.holding = holding
	//fmt.Println("Analyzing hand ", h.holding)

	//JT55T
	numJokers := strings.Count(h.holding, "J")
	for i := 0; i < len(h.holding); i++ {
		cardValue := convertToCard[string(h.holding[i])]
		cardCount := strings.Count(h.holding, string(h.holding[i]))

		fmt.Println("Cardvalue is ", cardValue, "Card Count is ", cardCount)

		if cardValue == JOKER && h.holding != "JJJJJ" {
			continue
		}

		switch cardCount {
		case 5:
			h.camelHandType = FIVEOAK
			h.fiveOak = cardValue
			return
		case 4:
			h.camelHandType = FOUROAK
			h.fourOak = cardValue
		case 3:
			if h.onePair != 0 {
				h.camelHandType = FULLHOUSE
				h.threeOak = cardValue
				break
			}
			h.camelHandType = THREEOAK
			h.threeOak = cardValue
		case 2:
			if h.camelHandType == THREEOAK {
				h.onePair = cardValue
				h.camelHandType = FULLHOUSE
			}
			if h.camelHandType != FULLHOUSE {
				if h.onePair == 0 {
					h.camelHandType = ONEPAIR
					h.onePair = cardValue
				}
				if cardValue > h.onePair {
					h.camelHandType = TWOPAIR
					h.twoPair = cardValue
				}
				if cardValue < h.onePair {
					h.camelHandType = TWOPAIR
					t := h.onePair
					h.twoPair = t
					h.onePair = cardValue
				}
			}
		case 1:
			if h.highCard == 0 {
				h.highCard = cardValue
				break
			}
			if h.highCard < cardValue {
				h.highCard = cardValue
			}
			if cardValue < h.highCard {
				h.secondHighestCard = cardValue
			}
		}
	}

	if h.camelHandType == "" {
		h.camelHandType = HIGHCARD
	}

	if numJokers > 0 {
		h.hasJoker = true
	}

	switch h.camelHandType {
	case FOUROAK:
		if numJokers == 1 {
			h.camelHandType = FIVEOAK
			break
		}

	case FULLHOUSE, THREEOAK:
		if numJokers == 1 {
			h.camelHandType = FOUROAK
			break
		}
		if numJokers == 2 {
			h.camelHandType = FIVEOAK
			break
		}

	case TWOPAIR:
		if numJokers == 1 {
			h.camelHandType = FULLHOUSE
			break
		}
		if numJokers == 2 {
			h.camelHandType = FOUROAK
			break
		}
		if numJokers == 3 {
			h.camelHandType = FIVEOAK
			break
		}

	case ONEPAIR:
		if numJokers == 1 {
			h.camelHandType = THREEOAK
			break
		}
		if numJokers == 2 {
			h.camelHandType = FOUROAK
			break
		}
		if numJokers == 3 {
			h.camelHandType = FIVEOAK
			break
		}
	case HIGHCARD:
		if numJokers == 1 {
			h.camelHandType = ONEPAIR
			break
		}
		if numJokers == 2 {
			h.camelHandType = THREEOAK
			break
		}
		if numJokers == 3 {
			h.camelHandType = FOUROAK
			break
		}

		if numJokers == 4 {
			h.camelHandType = FIVEOAK
			break
		}
	}

}

func (h *hand) printHandInfo() {
	var handNotes string

	if h.camelHandType == HIGHCARD {
		handNotes = fmt.Sprintf(", High Card: %s", convertToString[h.highCard])
	}

	if h.camelHandType == ONEPAIR {
		handNotes = fmt.Sprintf(", High Pair: %s, High Card: %s", convertToString[h.onePair], convertToString[h.highCard])
	}

	if h.camelHandType == TWOPAIR {
		handNotes = fmt.Sprintf(", High Pair: %s, Low Pair: %s, High Card: %s", convertToString[h.twoPair], convertToString[h.onePair], convertToString[h.highCard])
	}

	if h.camelHandType == THREEOAK {
		handNotes = fmt.Sprintf(", Three Of A Kind: %s, High Card: %s", convertToString[h.threeOak], convertToString[h.highCard])
	}

	if h.camelHandType == FULLHOUSE {
		handNotes = fmt.Sprintf(", Three Of A Kind: %s, High Pair: %s", convertToString[h.threeOak], convertToString[h.onePair])
	}

	if h.camelHandType == FOUROAK {
		handNotes = fmt.Sprintf(", Four Of A Kind: %s, High Card: %s", convertToString[h.fourOak], convertToString[h.highCard])
	}

	if h.camelHandType == FIVEOAK {
		handNotes = fmt.Sprintf(", Five Of A Kind: %s, High Card: %s", convertToString[h.fiveOak], convertToString[h.highCard])
	}

	fmt.Printf("Hand [%s] bid[%d] is a %s%s\n", h.holding, h.bid, h.camelHandType, handNotes)
}

func main() {
	var fileName string

	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	var hands []hand
	for _, line := range data {
		t := strings.Split(line, " ")

		//fmt.Printf("Analyzing hand #%d [%s]\n", i, t[0])

		var h hand
		bid, _ := strconv.Atoi(t[1])
		h.bid = bid
		h.determineHand(t[0])
		hands = append(hands, h)
		//		h.printHandInfo()
	}

	// Rough quick sort
	sort.Slice(hands[:], func(i, j int) bool {
		return handOrder[hands[i].camelHandType] < handOrder[hands[j].camelHandType]
	})

	// for _, h := range hands {

	// 	h.printHandInfo()
	// }

	SortOrder := []handType{
		HIGHCARD,
		ONEPAIR,
		TWOPAIR,
		THREEOAK,
		FULLHOUSE,
		FOUROAK,
		FIVEOAK,
	}

	var sortedHands []hand
	for i := 0; i < len(SortOrder); i++ {
		var tmp []hand

		for h := 0; h < len(hands); h++ {
			if hands[h].camelHandType == SortOrder[i] {
				tmp = append(tmp, hands[h])
			}
		}

		sort.SliceStable(tmp[:], func(i, j int) bool {
			for x := 0; x < len(tmp[i].holding); x++ {
				if convertToCard[string(tmp[i].holding[x])] == convertToCard[string(tmp[j].holding[x])] {
					continue
				}
				return convertToCard[string(tmp[i].holding[x])] < convertToCard[string(tmp[j].holding[x])]
			}
			return false
		})

		sortedHands = append(sortedHands, tmp...)
		tmp = []hand{}
	}

	totalWinnings := 0
	for i, h := range sortedHands {
		rank := h.bid * (i + 1)
		totalWinnings += rank
		h.printHandInfo()
	}

	fmt.Printf("Total Winnings is %d\n", totalWinnings)

}
