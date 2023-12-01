package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	calWords := make(map[string]rune)
	calWords["one"] = '1'
	calWords["two"] = '2'
	calWords["three"] = '3'
	calWords["four"] = '4'
	calWords["five"] = '5'
	calWords["six"] = '6'
	calWords["seven"] = '7'
	calWords["eight"] = '8'
	calWords["nine"] = '9'

	var rowDigit string
	var totalCount int

	for _, line := range data {
		fmt.Println(line)
		isFirst := true
		var f, l rune
		for rowIndex, rowChar := range line {
			if unicode.IsDigit(rowChar) {
				if isFirst {
					f = rowChar
					isFirst = false
				}
				l = rowChar
			}

			// nope, check agains worded calibration rules
			for k, v := range calWords {

				// compare the rowIndex against our map key, if the len(line) > len(k) + rowIndex that means we're out of bounds on the line check, so skip this key
				if len(line) < rowIndex+len(k) {
					fmt.Println("Skipping ", k, "(", len(k), ") rowIndex is ", rowIndex, " and len(line) is ", len(line))
					continue
				}

				fmt.Println("Checking ", k, "(", len(k), ") rowIndex is ", rowIndex, " and len(line) is ", len(line))

				// length of key won't overflow check, so we want to check if line[rowIndex:len(rowIndex+k)] == k, if it is and isFirst, set f.. flip Flag, otherwise set l

				if strings.EqualFold(line[rowIndex:rowIndex+len(k)], k) {
					fmt.Println("rowIndex:rowIndex+len(k) = ", line[rowIndex:rowIndex+len(k)], " k = ", k)
					if isFirst {
						f = rune(v)
						isFirst = false
					} else {
						l = rune(v)
					}
				}
			}
		}

		rowDigit = string(f) + string(l)
		fmt.Println("line is: ", line, " .. found digits: ", rowDigit)
		rowCount, err := strconv.Atoi(rowDigit)
		if err == nil {
			totalCount += rowCount
		}
		rowCount = 0
		rowDigit = ""
		isFirst = true
	}

	fmt.Println("Calibration Value is ", totalCount)

	// part two

}
