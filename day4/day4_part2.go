package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type cardType string

type Card struct {
	matches  int
	cardName string
	copies   int
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

		// create the card
		var myCard Card
		myCard.cardName = cardName
		myCard.matches = numMatches

		deck = append(deck, myCard)

		// calculate the points
		var points int
		if numMatches > 0 {
			points = 1 << (numMatches - 1)
		} else {
			points = 0
		}

		rollingSum += points

		fmt.Println(cardName, " matches ", numMatches)
	}

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
	fmt.Println("Total Points is ", rollingSum)
	fmt.Println("Total cards ", totalCards)
}
