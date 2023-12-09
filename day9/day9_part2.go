package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func differ(n []int) []int {

	var j []int
	for i := len(n) - 1; i > 0; i-- {
		if i > 0 {
			d := n[i] - n[i-1]
			j = append(j, d)
		}
	}
	for i := 0; i < len(j)/2; i++ {
		k := len(j) - i - 1
		j[i], j[k] = j[k], j[i]
	}
	return j
}

func main() {
	var fileName string

	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	total := 0
	for i, line := range data {

		s := strings.Split(line, " ")
		var n []int
		for i = 0; i < len(s); i++ {
			j, _ := strconv.Atoi(s[i])
			n = append(n, j)
		}

		invertedPyramid := make([][]int, 0)
		invertedPyramid = append(invertedPyramid, n)

		for {
			k := differ(n)

			zero := 0
			for i := 0; i < len(k); i++ {
				zero += k[i]
			}

			if zero != 0 {
				invertedPyramid = append(invertedPyramid, k)
				n = k
			} else {
				break
			}
		}

		var subIndex []int
		for i, v := range invertedPyramid {
			subIndex = append(subIndex, v[0])
			if i+1 == len(invertedPyramid) {
				startVal := subIndex[len(subIndex)-1]

				for x := len(subIndex) - 2; x >= 0; x-- {
					if x == 0 {
						total += subIndex[x] - startVal
					}
					startVal = subIndex[x] - startVal
				}
				subIndex = []int{}
			}
		}

	}
	fmt.Println("total is ", total)
}
