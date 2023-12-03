package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

var checkMapDigits = make([][]int, 0)
var checkMapDots = make([][]int, 0)

type debugLine struct {
	leftCheck  string
	rightCheck string
	backLine   string
	line       string
	bottomLine string
	partNumber int
	lineNumber int
}

func main() {
	var fileName string
	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	numberExtractor := regexp.MustCompile("[0-9]+")
	starExtractor := regexp.MustCompile(`[\*]`)
	rollingSum := 0
	for i := 1; i <= len(data); i++ {

		if i < len(data) {
			topLine := data[i-1]
			starLine := data[i]

			var bottomLine string
			if i+1 >= len(data) {
				fmt.Printf("wtf!\n")
				bottomLine = ""
			} else {
				fmt.Printf("okokok\n")
				bottomLine = data[i+1]
			}

			fmt.Println("t[", i, "] ", topLine)
			fmt.Println("s[", i+1, "] ", starLine)
			fmt.Println("b[", i+2, "] ", bottomLine)
			topLineNumber := i
			starLineNumber := i + 1
			bottomLineNumber := i + 2

			ti := numberExtractor.FindAllIndex([]byte(topLine), -1)    // top index
			ci := numberExtractor.FindAllIndex([]byte(starLine), -1)   // current index
			bi := numberExtractor.FindAllIndex([]byte(bottomLine), -1) // bottom index
			si := starExtractor.FindAllIndex([]byte(starLine), -1)     // star index

			for _, starIndex := range si {
				fmt.Printf("line %d, star found at position: %d\n", starLineNumber, starIndex[0])

				// do check on topline
				var topNumberOne, topNumberTwo int = 0, 0

				for _, topIndex := range ti {
					if starIndex[0] >= topIndex[0]-1 && starIndex[0] <= topIndex[1] {
						fmt.Printf("starIndex/topIndex adjacency [line %d] topIndex[0]-1:%d and [line %d] starIndex[0]:%d ...", topLineNumber, topIndex[0]-1, starLineNumber, starIndex[0])
						if topNumberOne == 0 {
							topNumberOne, err = strconv.Atoi(topLine[topIndex[0]:topIndex[1]])
							if err != nil {
								fmt.Printf("unable to extract top number one: %s\n", err)
							} else {
								fmt.Printf("number that is adjacent with star is %d\n", topNumberOne)
							}
						} else {
							topNumberTwo, err = strconv.Atoi(topLine[topIndex[0]:topIndex[1]])
							if err != nil {
								fmt.Printf("unable to extract adjacent number two: %s\n", err)
							}
							fmt.Printf("we have found another adjacent top number.. %d\n", topNumberTwo)
						}
					}
				}

				// do check on the starline
				var starNumberOne, starNumberTwo int = 0, 0
				for _, currentIndex := range ci {
					if starIndex[0] >= currentIndex[0]-1 && starIndex[0] <= currentIndex[1] {
						fmt.Printf("starIndex/currentIndex adjacency [line %d] currentIndex[0]-1:%d and [line %d] starIndex[0]:%d ...", starLineNumber, currentIndex[0]-1, starLineNumber, starIndex[0])
						if starNumberOne == 0 {
							starNumberOne, err = strconv.Atoi(starLine[currentIndex[0]:currentIndex[1]])
							if err != nil {
								fmt.Printf("unable to extract star number: %s\n", err)
							} else {
								fmt.Printf("number that is adjacent with star is %d\n", starNumberOne)
							}

						} else {
							starNumberTwo, err = strconv.Atoi(starLine[currentIndex[0]:currentIndex[1]])
							if err != nil {
								fmt.Printf("unable to extract adjacent number: %s\n", err)
							}
							fmt.Printf("we have found another adjacent star number.. %d\n", starNumberTwo)
						}
					}
				}

				var bottonNumberOne, bottonNumberTwo int = 0, 0
				for _, bottomIndex := range bi {
					if starIndex[0] >= bottomIndex[0]-1 && starIndex[0] <= bottomIndex[1] {
						fmt.Printf("starIndex/bottomIndex adjacency [line %d] bottomIndex[0]-1:%d and [line %d] starIndex[0]:%d ...", bottomLineNumber, bottomIndex[0]-1, starLineNumber, starIndex[0])
						if bottonNumberOne == 0 {
							bottonNumberOne, err = strconv.Atoi(bottomLine[bottomIndex[0]:bottomIndex[1]])
							if err != nil {
								fmt.Printf("unable to extract bottom number: %s\n", err)
							} else {
								fmt.Printf("Number that is adjacent with star is %d\n", bottonNumberOne)
							}
						} else {
							bottonNumberTwo, err = strconv.Atoi(bottomLine[bottomIndex[0]:bottomIndex[1]])
							if err != nil {
								fmt.Printf("unable to extract adjacent number: %s\n", err)
							}
							fmt.Printf("we have found another adjacent bottom number.. %d\n", bottonNumberTwo)
						}
					}
				}

				fmt.Println("topNumberOne ", topNumberOne)
				fmt.Println("topNumberTwo ", topNumberTwo)
				fmt.Println("starNumberOne ", starNumberOne)
				fmt.Println("starNumberTwo ", starNumberTwo)
				fmt.Println("bottonNumberOne ", bottonNumberOne)
				fmt.Println("bottonNumberTwo ", bottonNumberTwo)

				var loopCheck [6]int = [6]int{
					topNumberOne,
					topNumberTwo,
					starNumberOne,
					starNumberTwo,
					bottonNumberOne,
					bottonNumberTwo,
				}

				numberOfAdjacencies := 0
				var n1, n2 int
				for _, v := range loopCheck {
					if v != 0 {
						if n2 == 0 && n1 != 0 {
							fmt.Printf("setting n2 to %d\n", n2)
							n2 = v
						}
						if n1 == 0 {
							fmt.Printf("setting n1 to %d\n", n1)
							n1 = v
						}
						fmt.Println("v is ", v)
						numberOfAdjacencies++
					}
				}

				if numberOfAdjacencies > 2 {
					fmt.Println("This starNumber exceeds adjacency rules..")
				}

				if numberOfAdjacencies < 2 {
					fmt.Println("This starNumber does not meet adjacency rules..")
				}

				if numberOfAdjacencies == 2 {
					rollingSum += n1 * n2
					fmt.Println("This starNumber has exactly two adjacencies.. product is ", n1*n2)
					fmt.Println("Rolling Sum is", rollingSum)

				}

			}
			fmt.Println()
			//time.Sleep(10 * time.Second)

		}
	}
	fmt.Println("answer is ", rollingSum)
}
