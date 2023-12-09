package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	part1 := 0
	part2 := 0

	for r.Scan() {
		line := r.Text()
		nrStrings := strings.Split(line, " ")
		nrs := make([]int, 0)
		for _, nrString := range nrStrings {
			tempNr, _ := strconv.Atoi(nrString)
			nrs = append(nrs, tempNr)
		}
		fmt.Println(nrs)
		temp := findNext(nrs)
		part1 += temp

		temp2 := findPrevious(nrs)
		fmt.Println(temp2)
		part2 += temp2
	}

	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)
}

func findNext(nrs []int) int {
	difs := make([]int, 0)
	for i := 0; i < len(nrs)-1; i++ {
		difs = append(difs, nrs[i+1]-nrs[i])
	}
	fmt.Println(difs)
	if checkAllZero(difs) {
		return nrs[len(nrs)-1] + difs[len(difs)-1]
	}

	return findNext(difs) + nrs[len(nrs)-1]
}

func findPrevious(nrs []int) int {
	difs := make([]int, 0)
	for i := 0; i < len(nrs)-1; i++ {
		difs = append(difs, nrs[i+1]-nrs[i])
	}
	fmt.Println(difs)
	if checkAllZero(difs) {
		return nrs[0] - difs[0]
	}

	return nrs[0] - findPrevious(difs)
}

func checkAllZero(nrs []int) bool {
	for _, nr := range nrs {
		if nr != 0 {
			return false
		}
	}
	return true
}
