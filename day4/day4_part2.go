package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type cardType string

const (
	ORIGINAL cardType = "ORIGINAL"
	COPY     cardType = "COPY"
)

type Card struct {
	originalOrCopy cardType
	matches        int
	cardName       string
	winning        []int
	mine           []int
	points         int
	copies         int
}

func main() {
	var fileName string

	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	var rollingSum int
	var deck []Card
	for _, card := range data {
		val := strings.Split(card, ":")
		cardName := val[0]
		matches := strings.Split(val[1], "|")

		var winNum []int
		wS := strings.Split(matches[0], " ")
		for _, v := range wS {
			t, err := strconv.Atoi(v)
			if err == nil {
				winNum = append(winNum, t)
			}
		}

		var myNum []int
		mS := strings.Split(matches[1], " ")
		for _, v := range mS {
			t, err := strconv.Atoi(v)
			if err == nil {
				myNum = append(myNum, t)
			}
		}

		numMatches := 0
		for _, w := range winNum {
			for _, m := range myNum {
				if w == m {
					numMatches++
				}
			}
		}

		var points int
		if numMatches == 1 {
			points = 1
		}
		if numMatches == 2 {
			points = 2
		}
		if numMatches == 3 {
			points = 4
		}
		if numMatches == 4 {
			points = 8
		}
		if numMatches == 5 {
			points = 16
		}
		if numMatches == 6 {
			points = 32
		}
		if numMatches == 7 {
			points = 64
		}
		if numMatches == 8 {
			points = 128
		}
		if numMatches == 9 {
			points = 256
		}
		if numMatches == 10 {
			points = 512
		}

		// create the card
		var myCard Card
		myCard.originalOrCopy = ORIGINAL
		myCard.cardName = cardName
		myCard.points = points
		myCard.mine = myNum
		myCard.matches = numMatches
		myCard.winning = winNum

		deck = append(deck, myCard)
		rollingSum += points

		fmt.Println(cardName, " matches ", numMatches)
		points = 0
	}
	fmt.Println("Total Points is ", rollingSum)

	firstRun := true
	for i := 0; i < len(deck); i++ {
		if firstRun {
			currentCard := deck[i]
			fmt.Println("Processing ", deck[i].cardName, " with ", currentCard.matches, " points..")
			for scratchCopies := 1; scratchCopies <= currentCard.matches; scratchCopies++ {
				deck[i+scratchCopies].copies++
				fmt.Println("increased copy for ", deck[i+scratchCopies].cardName, " to include ", deck[i+scratchCopies].copies)
			}
			firstRun = false
		} else {
			currentCard := deck[i]
			fmt.Println("Processing ", deck[i].cardName, " with ", currentCard.matches, " points.. with ", deck[i].copies, " copies")
			for copyLoop := 0; copyLoop <= deck[i].copies; copyLoop++ {
				for scratchCopies := 1; scratchCopies <= currentCard.matches; scratchCopies++ {
					deck[i+scratchCopies].copies++
					fmt.Println("increased copy for ", deck[i+scratchCopies].cardName, " to include ", deck[i+scratchCopies].copies)
				}
			}
		}
	}

	totalCards := 0
	for i := 0; i < len(deck); i++ {
		totalCards++
		totalCards += deck[i].copies
	}
	fmt.Println("Total cards ", totalCards)
}
