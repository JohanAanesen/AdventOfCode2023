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

	sum := 0

	cards := make([]string, 0)
	foundCards := make(map[int]int, 0)

	for r.Scan() {
		line := r.Text()

		nrMap := make(map[int]int, 0)
		correctTickets := 0
		tempSplit := strings.Split(line, ":")
		cards = append(cards, tempSplit[1])
		numberSplit := strings.Split(tempSplit[1], "|")
		lotteryNumbers := strings.Split(numberSplit[0], " ")
		ticketNumbers := strings.Split(numberSplit[1], " ")

		for _, number := range lotteryNumbers {
			nr, err := strconv.Atoi(number)
			if err == nil {
				nrMap[nr] = 1
			}

		}

		for _, number := range ticketNumbers {
			nr, err := strconv.Atoi(number)
			if err == nil {
				if nrMap[nr] == 1 {
					correctTickets++
				}
			}

		}
		if correctTickets > 0 {
			tempSum := 1
			for i := 1; i < correctTickets; i++ {
				tempSum *= 2
			}
			sum += tempSum
		}

	}

	fmt.Println("Part1: ", sum)

	part2Sum := 0
	for i, _ := range cards {
		part2Sum += findCards(i, cards, &foundCards)
	}

	fmt.Println("Part2: ", part2Sum)

}

func findCards(index int, cards []string, foundCards *map[int]int) int {
	numberSplit := strings.Split(cards[index], "|")
	lotteryNumbers := strings.Split(numberSplit[0], " ")
	ticketNumbers := strings.Split(numberSplit[1], " ")

	correctTickets := findNrWinnings(lotteryNumbers, ticketNumbers)

	if correctTickets > 0 {
		tempSum := 1
		for i := 1; i <= correctTickets; i++ {
			if (*foundCards)[index+i] == 0 {
				tempSum += findCards(index+i, cards, foundCards)
			} else {
				tempSum += (*foundCards)[index+i]
			}

		}
		(*foundCards)[index] = tempSum
		return tempSum
	} else {
		(*foundCards)[index] = 1
		return 1
	}
}

func findNrWinnings(lottery []string, ticket []string) int {
	nrMap := make(map[int]int, 0)
	correctTickets := 0
	for _, number := range lottery {
		nr, err := strconv.Atoi(number)
		if err == nil {
			nrMap[nr] = 1
		}

	}

	for _, number := range ticket {
		nr, err := strconv.Atoi(number)
		if err == nil {
			if nrMap[nr] == 1 {
				correctTickets++
			}
		}

	}
	return correctTickets
}
