package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
)

func findLineAdjacencies(li [][]int, checkLine, starLine string, starIndex int) (int, int) {
	var numberOne, numberTwo int = 0, 0
	var err error
	for _, index := range li {
		if starIndex >= index[0]-1 && starIndex <= index[1] {
			if numberOne == 0 {
				numberOne, err = strconv.Atoi(checkLine[index[0]:index[1]])
				if err != nil {
					fmt.Printf("unable to extract number: %s\n", err)
				} else {
					fmt.Printf("Number that is adjacent with is %d\n", numberOne)
				}
			} else {
				numberTwo, err = strconv.Atoi(checkLine[index[0]:index[1]])
				if err != nil {
					fmt.Printf("unable to extract adjacent number: %s\n", err)
				}
				fmt.Printf("we have found another adjacent number.. %d\n", numberTwo)
			}
		}
	}
	return numberOne, numberTwo
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

			ti := numberExtractor.FindAllIndex([]byte(topLine), -1)    // top index
			ci := numberExtractor.FindAllIndex([]byte(starLine), -1)   // current index
			bi := numberExtractor.FindAllIndex([]byte(bottomLine), -1) // bottom index
			si := starExtractor.FindAllIndex([]byte(starLine), -1)     // star index

			for _, starIndex := range si {
				fmt.Printf("line %d, star found at position: %d\n", i+1, starIndex[0])

				// do checks
				topNumberOne, topNumberTwo := findLineAdjacencies(ti, topLine, starLine, starIndex[0])	
				starNumberOne, starNumberTwo := findLineAdjacencies(ci, starLine, starLine, starIndex[0])
				bottonNumberOne, bottonNumberTwo := findLineAdjacencies(bi, bottomLine, starLine, starIndex[0])

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
