package main

import (
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var fileName string
	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	digitMatch := regexp.MustCompile("[0-9]+")

	// checkRed := 12
	// checkGreen := 13
	// checkBlue := 14

	power := 0

	for _, line := range data {
		//fmt.Println(line)
		splitString := strings.Split(line, ":")

		if len(splitString) > 1 {

			// id, err := strconv.Atoi(digitMatch.FindString(splitString[0]))
			// if err != nil {
			// 	fmt.Println("Couldn't find ID for line[", i, "]: ", line)
			// 	return
			// }

			// extract the match sets

			var minRed, minGreen, minBlue int
			matchSets := strings.Split(splitString[1], ";")
			for _, set := range matchSets {
				setSplit := strings.Split(set, ",")
				for _, cube := range setSplit {
					if strings.Contains(cube, "red") {
						r, err := strconv.Atoi(digitMatch.FindString(cube))
						if err != nil {
							fmt.Println("We couldn't get the red cube value for ", cube)
						}
						if minRed < r {
							minRed = r
						}
					}
					if strings.Contains(cube, "green") {
						g, err := strconv.Atoi(digitMatch.FindString(cube))
						if err != nil {
							fmt.Println("We couldn't get the green cube value for ", cube)
						}
						if minGreen < g {
							minGreen = g
						}
					}
					if strings.Contains(cube, "blue") {
						b, err := strconv.Atoi(digitMatch.FindString(cube))
						if err != nil {
							fmt.Println("We couldn't get the cube value for ", cube)
						}
						if minBlue < b {
							minBlue = b
						}
					}
				}
			}

			power += minRed * minGreen * minBlue

		}
	}

	fmt.Println(power)
}
