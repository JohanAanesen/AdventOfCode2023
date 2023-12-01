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

	calibrationNumbers := make([]int, 0)

	for r.Scan() {
		line := r.Text()
		tempArr := strings.Split(line, "")

		//lowest
		lowestNrNr := "99999"
		lowestNrIndex := 9999
		//highest
		highestNrNr := "-1"
		highestNrIndex := -1
		lowestStringIndex, lowestStringNumber, highestStringIndex, highestStringNumber := checkStringNumber(line)

		for i := 0; i < len(tempArr); i++ {
			if _, err := strconv.Atoi(tempArr[i]); err == nil {
				lowestNrNr = tempArr[i]
				lowestNrIndex = i
				break
			}
		}

		//highest
		for i := len(tempArr) - 1; i >= 0; i-- {
			if _, err := strconv.Atoi(tempArr[i]); err == nil {
				highestNrNr = tempArr[i]
				highestNrIndex = i
				break
			}
		}

		if lowestStringIndex < lowestNrIndex {
			lowestNrNr = lowestStringNumber
		}
		if highestStringIndex > highestNrIndex {
			highestNrNr = highestStringNumber
		}

		returnStr := lowestNrNr + highestNrNr

		value, _ := strconv.Atoi(returnStr)
		fmt.Println(line, value)

		calibrationNumbers = append(calibrationNumbers, value)

	}

	sum := 0

	for _, number := range calibrationNumbers {
		sum += number
	}

	fmt.Println("Part2: ", sum)

}

func checkStringNumber(input string) (int, string, int, string) {

	numberStrings := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	lowestIndex := 999
	lowestNumber := "999999"
	highestIndex := -1
	highestNumber := "-1"

	for i, numberString := range numberStrings {
		if strings.Index(input, numberString) < lowestIndex && strings.Index(input, numberString) != -1 {
			lowestIndex = strings.Index(input, numberString)
			lowestNumber = strconv.Itoa(i + 1)
		}

		if strings.LastIndex(input, numberString) > highestIndex && strings.LastIndex(input, numberString) != -1 {
			highestIndex = strings.LastIndex(input, numberString)
			highestNumber = strconv.Itoa(i + 1)
		}
	}

	return lowestIndex, lowestNumber, highestIndex, highestNumber
}
