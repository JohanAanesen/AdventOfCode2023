package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type coordinate struct {
	x int
	y int
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	niceStrings := 0
	niceStrings2 := 0

	for r.Scan() {
		line := r.Text()

		if checkVowels(line) && checkDoubles(line) && checkIllegals(line) {
			niceStrings++
		}
		if checkPairTwice(line) && checkRepeatsBetween(line) {
			niceStrings2++
		}

	}

	fmt.Println("Part1: ", niceStrings)
	fmt.Println("Part2: ", niceStrings2)

}

func checkVowels(line string) bool {
	vowels := []string{"a", "e", "i", "o", "u"}
	vowelCount := 0

	for _, vowel := range vowels {
		vowelCount += strings.Count(line, vowel)
	}

	if vowelCount >= 3 {
		return true
	}
	return false
}

func checkDoubles(line string) bool {
	list := strings.Split(line, "")
	for i := 0; i < len(list)-1; i++ {
		if list[i] == list[i+1] {
			return true
		}
	}
	return false
}

func checkIllegals(line string) bool {
	illegals := []string{"ab", "cd", "pq", "xy"}

	for _, illegal := range illegals {
		if strings.Contains(line, illegal) {
			return false
		}
	}

	return true
}

func checkPairTwice(line string) bool {
	list := strings.Split(line, "")
	for i := 0; i < len(list)-2; i++ {
		tempString := list[i] + list[i+1]
		if strings.Contains(strings.Join(list[i+2:], ""), tempString) {
			return true
		}
	}
	return false
}

func checkRepeatsBetween(line string) bool {
	list := strings.Split(line, "")
	for i := 0; i < len(list)-2; i++ {
		if list[i] == list[i+2] {
			return true
		}
	}
	return false
}
