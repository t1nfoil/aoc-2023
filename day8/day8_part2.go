package main

import (
	"flag"
	"fmt"
	"strings"
)

var instruction string

type node struct {
	name  string
	lName string
	l     *node
	rName string
	r     *node
}

func main() {
	var fileName string

	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	// build root nodes
	var navNodes []node
	for i, line := range data {
		if i == 0 {
			instruction = line
			continue
		}

		if line != "" {
			lsplit := strings.Split(line, "=")
			lsplit[0] = strings.ReplaceAll(lsplit[0], " ", "")
			lsplit[1] = strings.ReplaceAll(lsplit[1], " ", "")
			//			fmt.Printf("%s\n%x\n", lsplit[1], lsplit[1])
			rlsplit := strings.Split(lsplit[1], ",")
			lName := strings.TrimPrefix(rlsplit[0], "(")
			//			fmt.Printf("lName: %s\n%x\n", lName, lName)
			rName := strings.TrimSuffix(rlsplit[1], ")")
			//			fmt.Printf("rName: %s\n%x\n", rName, rName)

			n := node{
				name:  lsplit[0],
				lName: lName,
				rName: rName,
			}

			navNodes = append(navNodes, n)

		}
	}

	for i, n := range navNodes {
		for il, leaf := range navNodes {
			if strings.EqualFold(n.lName, leaf.name) {
				navNodes[i].l = &navNodes[il]
			}
		}
		for ir, leaf := range navNodes {
			if strings.EqualFold(n.rName, leaf.name) {
				navNodes[i].r = &navNodes[ir]
			}
		}
	}

	var activeNodes []node
	var aSteps []int
	for i, n := range navNodes {
		if string(n.name[2]) == "A" {
			activeNodes = append(activeNodes, navNodes[i])
			aSteps = append(aSteps, 0)
		}
	}

	ip := 0
	var rcmVals []int
	for {
		if ip < len(instruction) {
			if string(instruction[ip]) == "L" {
				for i := 0; i < len(activeNodes); i++ {
					activeNodes[i] = *activeNodes[i].l
					aSteps[i]++
				}
				ip++

			} else {
				for i := 0; i < len(activeNodes); i++ {
					activeNodes[i] = *activeNodes[i].r
					aSteps[i]++
				}
				ip++
			}
			for i := 0; i < len(activeNodes); i++ {

				if string(activeNodes[i].name[2]) == "Z" {
					rcmVals = append(rcmVals, aSteps[i])
					fmt.Printf("Check %s, ", activeNodes[i].name)
					fmt.Printf("Steps: %d\n", aSteps[i])
				}
			}
		} else {
			ip = 0
			if len(rcmVals) == len(activeNodes) {
				goto FINISH
			}
		}
	}
FINISH:
	pythonString := "python3 ./check.py --rcmlist "
	rcmString := ""
	for _, v := range rcmVals {
		rcmString += fmt.Sprintf("%d ", v)
	}

	fmt.Println("run -> ", pythonString, rcmString)
}
