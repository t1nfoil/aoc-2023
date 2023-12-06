package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

type microrace struct {
	winTime            int
	winDistance        int
	buttonPressTime    []int
	buttonBoatTime     []int
	buttonBoatDistance []int
	buttonPressWinners int
}

func (m *microrace) buttonPressCalc() {
	buttonPressWinners := 0
	for i := 0; i < m.winTime; i++ {
		buttonPressTime := i
		timeLeft := m.winTime - buttonPressTime
		dPmm := (m.winTime - buttonPressTime) * 1
		buttonDistance := i * dPmm

		if buttonDistance > m.winDistance {
			m.buttonPressTime = append(m.buttonPressTime, buttonPressTime)
			m.buttonBoatTime = append(m.buttonBoatTime, timeLeft)
			m.buttonBoatDistance = append(m.buttonBoatDistance, buttonDistance)
			buttonPressWinners++
		}
	}
	m.buttonPressWinners = buttonPressWinners
}

func (m *microrace) displayWinningButtonPresses() {
	for i := 0; i < len(m.buttonPressTime); i++ {
		//	formatString := "A buttonpress of %d milliseconds, beats the win time/distance[%d ms / %d mm] with a time/distance of [ %d ms / %d mm]\n"
		//	fmt.Printf(formatString, m.buttonPressTime[i], m.winTime, m.winDistance, m.buttonBoatTime[i], m.buttonBoatDistance[i])
	}
}

func main() {
	var fileName string

	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	var tl []string
	var dl []string
	digitExtractor := regexp.MustCompile("[0-9]+")
	for i, line := range data {
		if i == 0 {
			tl = digitExtractor.FindAllString(line, -1)
		}
		if i == 1 {
			dl = digitExtractor.FindAllString(line, -1)
		}
	}

	var boats []microrace
	for b := 0; b < len(tl); b++ {
		winTime, _ := strconv.Atoi(tl[b])
		winDistance, _ := strconv.Atoi(dl[b])
		boat := microrace{
			winTime:     winTime,
			winDistance: winDistance,
		}
		boats = append(boats, boat)
	}

	winners := 1
	for b := 0; b < len(boats); b++ {
		boats[b].buttonPressCalc()
		//boats[b].displayWinningButtonPresses()
		winners *= boats[b].buttonPressWinners
	}

	fmt.Println("Winners are: ", winners)
}
