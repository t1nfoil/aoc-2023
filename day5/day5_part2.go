package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type rangeMapping struct {
	mapName  string
	dstStart int
	dstEnd   int
	srcStart int
	srcEnd   int
}

type seedLocation struct {
	locationNumber int
	description    string
}

func (r *rangeMapping) getDst(item int) int {
	if item >= r.srcStart && item <= r.srcEnd {
		return r.dstStart + (item - r.srcStart)
	}
	return item
}

func main() {
	var fileName string

	flag.StringVar(&fileName, "inputfile", "input.txt", "the name of the input file")
	flag.Parse()

	data, err := openFile(fileName)
	if err != nil {
		panic(err)
	}

	// get seeds
	t := strings.Split(data[0], ":")
	seedRanges := strings.Split(t[1], " ")
	seedRanges = seedRanges[1:]

	// load and populate the rangeMaps
	seedMap := make(map[string][]rangeMapping)
	mapName := ""
	fmt.Println("Populating rangeMaps")
	for i := 1; i < len(data); i++ {
		if strings.Contains(data[i], "map:") {
			m := strings.Split(data[i], " ")
			mapName = m[0]
			continue
		}
		if data[i] == "" {
			continue
		}

		mapEntry := strings.Split(data[i], " ")
		dst, _ := strconv.Atoi(mapEntry[0])
		src, _ := strconv.Atoi(mapEntry[1])
		rng, _ := strconv.Atoi(mapEntry[2])

		r := rangeMapping{
			mapName:  mapName,
			dstStart: dst,
			dstEnd:   dst + rng - 1,
			srcStart: src,
			srcEnd:   src + rng - 1,
		}

		if sm, ok := seedMap[mapName]; ok {
			sm = append(sm, r)
			seedMap[mapName] = sm
			continue
		} else {
			var sm []rangeMapping
			sm = append(sm, r)
			seedMap[mapName] = sm
		}
	}

	for r := 0; r < len(seedRanges); r += 2 {

		start, _ := strconv.Atoi(seedRanges[r])
		end, _ := strconv.Atoi(seedRanges[r+1])

		fmt.Println(start, " ", end)
		return
		spanMod := (end - start) % 6
		spanLength := (end - start) / 6

		var spans []string
		//spanIndex := 0
		for spanner := 0; spanner <= 6; spanner += 2 {
			spans[spanner] = strconv.Itoa(start)
			start = start + (spanner * spanLength)
			if spanner == 6 {
				start += spanMod
			}
			spans[spanner+1] = strconv.Itoa(start)
		}

		for _, s := range spans {
			fmt.Println(s)
		}

		//crunchNumbers(seedRanges[r:r+1], seedMap)

	}

	// lowestLocation := 0
	// lowestDescription := ""

	// fmt.Println("Iterating Seed Ranges")
	// for r := 0; r < len(seedRanges)-1; r += 2 {
	// 	seedStart, _ := strconv.Atoi(seedRanges[r])
	// 	seedIterations, _ := strconv.Atoi(seedRanges[r+1])
	// 	fmt.Printf("\nstarting at %d-%d [seed range is +%d]\n", seedStart, seedStart+seedIterations-1, seedIterations)

	// 	//var tmpSeedLoc = []seedLocation{}

	// 	for seedNumber := seedStart; seedNumber < seedStart+seedIterations; seedNumber++ {
	// 		if seedNumber%1000000 == 0 {
	// 			fmt.Printf(".")
	// 		}
	// 		//fmt.Println("Processing seed ", seedNumber, "r is ", r)
	// 		mapOrder := []string{
	// 			"seed-to-soil",
	// 			"soil-to-fertilizer",
	// 			"fertilizer-to-water",
	// 			"water-to-light",
	// 			"light-to-temperature",
	// 			"temperature-to-humidity",
	// 			"humidity-to-location",
	// 		}

	// 		originalSeed := seedNumber
	// 		loopSeed := seedNumber

	// 		for _, conversionName := range mapOrder {
	// 			mapping := seedMap[conversionName]

	// 			var forwardLookups []int
	// 			for i := 0; i < len(mapping); i++ {
	// 				// fmt.Printf("Doing a lookup for map[%s] range-index is [%d] [%d to %d], loopSeed is %d, return is %d\n",
	// 				// 	conversionName,
	// 				// 	i,
	// 				// 	mapping[i].srcStart,
	// 				// 	mapping[i].srcEnd,
	// 				// 	loopSeed,
	// 				// 	mapping[i].getDst(loopSeed))
	// 				forwardLookups = append(forwardLookups, mapping[i].getDst(loopSeed))

	// 			}

	// 			//fmt.Printf("seedNumber is %d, loopSeed is %d", seedNumber, loopSeed)
	// 			//spew.Dump(forwardLookups)

	// 			locationNumber := loopSeed
	// 			for k := 0; k < len(forwardLookups); k++ {
	// 				if forwardLookups[k] != loopSeed {
	// 					locationNumber = forwardLookups[k]
	// 				}
	// 			}
	// 			//fmt.Println(conversionName, " seed number ", loopSeed, "maps to ", locationNumber)
	// 			loopSeed = locationNumber

	// 			if conversionName == "humidity-to-location" {
	// 				if lowestLocation == 0 {
	// 					lowestLocation = locationNumber
	// 					lowestDescription = fmt.Sprintf("seed number %d maps to location %d", originalSeed, locationNumber)
	// 					fmt.Printf("\n%s, lowestLocation is %d\n", lowestDescription, lowestLocation)
	// 				}
	// 				if lowestLocation > locationNumber {
	// 					lowestLocation = locationNumber
	// 					lowestDescription = fmt.Sprintf("seed number %d maps to location %d", originalSeed, locationNumber)
	// 					fmt.Printf("\n%s, lowestLocation is %d\n", lowestDescription, lowestLocation)
	// 				}
	// 				//fmt.Println(conversionName, "location -> seed number ", loopSeed, "maps to location ", locationNumber)
	// 				//tmpSeedLoc = append(tmpSeedLoc, seedLocation{description: fmt.Sprintf("seed number %d maps to location %d", originalSeed, locationNumber), locationNumber: locationNumber})
	// 			}

	// 			// if len(tmpSeedLoc) > 5000000 {
	// 			// 	fmt.Printf("!")
	// 			// 	lowestFoundValue := tmpSeedLoc[0]
	// 			// 	tmpSeedLoc = []seedLocation{}
	// 			// 	tmpSeedLoc = append(tmpSeedLoc, lowestFoundValue)
	// 			// }
	// 		}
	// 	}
	// 	// fmt.Printf("Sorting data..\n")
	// 	// sort.Slice(tmpSeedLoc[:], func(i, j int) bool {
	// 	// 	return tmpSeedLoc[i].locationNumber > tmpSeedLoc[j].locationNumber
	// 	// })
	// 	// lowestFoundValue := tmpSeedLoc[0]
	// 	// tmpSeedLoc = []seedLocation{}
	// 	// seedLocations = append(seedLocations, lowestFoundValue)
	// }
	// // sort.Slice(seedLocations[:], func(i, j int) bool {
	// // 	return seedLocations[i].locationNumber > seedLocations[j].locationNumber
	// // })
	// // fmt.Println()
	// // for i, s := range seedLocations {
	// // 	if i == len(seedLocations)-1 {
	// // 		fmt.Printf("%s <-- answer\n", s.description)
	// // 	} else {
	// // 		fmt.Printf("%s\n", s.description)
	// // 	}
	// // }
	// fmt.Printf("final answer: %s, lowestLocation is %d\n", lowestDescription, lowestLocation)
}

var mapOrder = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

func crunchNumbers(seedRanges []string, seedMap map[string][]rangeMapping) {
	lowestLocation := 0
	lowestDescription := ""

	fmt.Println("Iterating Seed Ranges")
	for r := 0; r < len(seedRanges)-1; r += 2 {
		seedStart, _ := strconv.Atoi(seedRanges[r])
		seedIterations, _ := strconv.Atoi(seedRanges[r+1])
		fmt.Printf("\nstarting at %d-%d [seed range is +%d]\n", seedStart, seedStart+seedIterations-1, seedIterations)

		for seedNumber := seedStart; seedNumber < seedStart+seedIterations; seedNumber++ {
			if seedNumber%1000000 == 0 {
				fmt.Printf(".")
			}

			originalSeed := seedNumber
			loopSeed := seedNumber

			for _, conversionName := range mapOrder {
				mapping := seedMap[conversionName]

				var forwardLookups []int
				for i := 0; i < len(mapping); i++ {
					forwardLookups = append(forwardLookups, mapping[i].getDst(loopSeed))

				}

				locationNumber := loopSeed
				for k := 0; k < len(forwardLookups); k++ {
					if forwardLookups[k] != loopSeed {
						locationNumber = forwardLookups[k]
					}
				}

				loopSeed = locationNumber

				if conversionName == "humidity-to-location" {
					if lowestLocation == 0 {
						lowestLocation = locationNumber
						lowestDescription = fmt.Sprintf("seed number %d maps to location %d", originalSeed, locationNumber)
						fmt.Printf("\n%s, lowestLocation is %d\n", lowestDescription, lowestLocation)
					}
					if lowestLocation > locationNumber {
						lowestLocation = locationNumber
						lowestDescription = fmt.Sprintf("seed number %d maps to location %d", originalSeed, locationNumber)
						fmt.Printf("\n%s, lowestLocation is %d\n", lowestDescription, lowestLocation)
					}
				}

			}
		}
	}
	fmt.Printf("final answer: %s, lowestLocation is %d\n", lowestDescription, lowestLocation)
}
