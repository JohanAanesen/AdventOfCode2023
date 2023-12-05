package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type inputMap struct {
	destIndex   int
	sourceIndex int
	length      int
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	step := 0

	seeds := make([]int, 0)
	seeds2 := make([]int, 0)
	s2s := make([]inputMap, 0)
	s2f := make([]inputMap, 0)
	f2w := make([]inputMap, 0)
	w2l := make([]inputMap, 0)
	l2t := make([]inputMap, 0)
	t2h := make([]inputMap, 0)
	h2l := make([]inputMap, 0)

	for r.Scan() {
		line := r.Text()
		if line == "" {
			step++
			continue
		}

		if step == 0 {
			line = strings.TrimPrefix(line, "seeds: ")
			seedString := strings.Split(line, " ")
			for _, s := range seedString {
				seedint, _ := strconv.Atoi(s)
				seeds = append(seeds, seedint)
			}

			for i := 0; i < len(seedString); i++ {
				seedint, _ := strconv.Atoi(seedString[i])
				lenint, _ := strconv.Atoi(seedString[i+1])

				for j := seedint; j < seedint+lenint; j++ {
					seeds2 = append(seeds2, j)
				}

				i++
			}
		} else if step == 1 {
			if line == "seed-to-soil map:" {
				continue
			}
			readDigits(line, &s2s)
		} else if step == 2 {
			if line == "soil-to-fertilizer map:" {
				continue
			}
			readDigits(line, &s2f)
		} else if step == 3 {
			if line == "fertilizer-to-water map:" {
				continue
			}
			readDigits(line, &f2w)
		} else if step == 4 {
			if line == "water-to-light map:" {
				continue
			}
			readDigits(line, &w2l)
		} else if step == 5 {
			if line == "light-to-temperature map:" {
				continue
			}
			readDigits(line, &l2t)
		} else if step == 6 {
			if line == "temperature-to-humidity map:" {
				continue
			}
			readDigits(line, &t2h)
		} else if step == 7 {
			if line == "humidity-to-location map:" {
				continue
			}
			readDigits(line, &h2l)
		}

	}

	seedLocations := make([]int, 0)
	lowestLocation := 99999999999999
	for _, seed := range seeds {
		s2ss := checkMapThing(seed, &s2s)
		s2fs := checkMapThing(s2ss, &s2f)
		f2ws := checkMapThing(s2fs, &f2w)
		w2ls := checkMapThing(f2ws, &w2l)
		l2ts := checkMapThing(w2ls, &l2t)
		t2hs := checkMapThing(l2ts, &t2h)
		h2ls := checkMapThing(t2hs, &h2l)
		seedLocations = append(seedLocations, h2ls)
		if h2ls < lowestLocation {
			lowestLocation = h2ls
		}
	}
	fmt.Println("Part1: ", lowestLocation)

	seedLocations2 := make([]int, 0)
	lowestLocation2 := 99999999999999
	for i, seed := range seeds2 {
		s2ss := checkMapThing(seed, &s2s)
		s2fs := checkMapThing(s2ss, &s2f)
		f2ws := checkMapThing(s2fs, &f2w)
		w2ls := checkMapThing(f2ws, &w2l)
		l2ts := checkMapThing(w2ls, &l2t)
		t2hs := checkMapThing(l2ts, &t2h)
		h2ls := checkMapThing(t2hs, &h2l)
		seedLocations2 = append(seedLocations2, h2ls)
		if h2ls < lowestLocation2 {
			lowestLocation2 = h2ls
		}

		if i%10000000 == 0 {
			fmt.Println("Progress: ", i, "/", len(seeds2))
		}

	}

	fmt.Println("Part2: ", lowestLocation2)
}

func checkMapThing(input int, storeMap *[]inputMap) int {
	for _, s := range *storeMap {
		if input >= s.sourceIndex && input < s.sourceIndex+s.length {
			tempNr := input - s.sourceIndex
			return s.destIndex + tempNr
		}
	}

	return input
}

func readDigits(line string, storeMap *[]inputMap) {
	temps2s := inputMap{}
	fmt.Sscanf(line, "%d %d %d", &temps2s.destIndex, &temps2s.sourceIndex, &temps2s.length)
	*storeMap = append(*storeMap, temps2s)
}
