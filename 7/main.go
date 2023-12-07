package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type poker struct {
	hand        string
	bid         int
	rating      int
	handRatings []int
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	const Order = "23456789TJQKA"
	const JokerOrder = "J23456789TQKA"
	orderMap := make(map[string]int)
	orderSplit := strings.Split(Order, "")
	for i := 0; i < len(orderSplit); i++ {
		orderMap[orderSplit[i]] = i
	}
	jokerOrderMap := make(map[string]int)
	jokerOrderSplit := strings.Split(JokerOrder, "")
	for i := 0; i < len(jokerOrderSplit); i++ {
		jokerOrderMap[jokerOrderSplit[i]] = i
	}
	hands := make([]poker, 0)

	hands2 := make([]poker, 0)

	for r.Scan() {
		line := r.Text()

		tempsplit := strings.Split(line, " ")
		bid, _ := strconv.Atoi(tempsplit[1])
		tempMap := make(map[string]int)
		handSplit := strings.Split(tempsplit[0], "")
		handrating := make([]int, 0)
		handrating2 := make([]int, 0)
		joker := 0
		for _, s := range handSplit {
			if s == "J" {
				joker++
			}
			tempMap[s]++
			handrating = append(handrating, orderMap[s])
			handrating2 = append(handrating2, jokerOrderMap[s])
		}

		hands = append(hands, poker{hand: tempsplit[0], bid: bid, rating: getRanking(tempMap), handRatings: handrating})
		hands2 = append(hands2, poker{hand: tempsplit[0], bid: bid, rating: getRankingJoker(tempMap, joker), handRatings: handrating2})

	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].rating == hands[j].rating {

			for k := 0; k < 5; k++ {
				if hands[i].handRatings[k] != hands[j].handRatings[k] {
					return hands[i].handRatings[k] < hands[j].handRatings[k]
				}
			}

			return true
		}

		return hands[i].rating < hands[j].rating

	})

	sort.Slice(hands2, func(i, j int) bool {
		if hands2[i].rating == hands2[j].rating {

			for k := 0; k < 5; k++ {
				if hands2[i].handRatings[k] != hands2[j].handRatings[k] {
					return hands2[i].handRatings[k] < hands2[j].handRatings[k]
				}
			}
		}

		return hands2[i].rating < hands2[j].rating
	})

	sum := 0
	for i, hand := range hands {
		sum += hand.bid * (i + 1)
	}
	fmt.Println("Part1: ", sum)

	sum2 := 0
	for i, hand := range hands2 {
		sum2 += hand.bid * (i + 1)
		fmt.Println(hand, " Rank: ", i+1, " Rating: ", hand.rating)
	}
	fmt.Println("Part2: ", sum2)
}

func getRanking(input map[string]int) int {
	oak := ofAKind(input)
	if oak == 5 {
		return 6
	} else if oak == 4 {
		return 5
	} else if oak == 3 {
		if isHouse(input) {
			return 4
		}
		return 3
	} else if oak == 2 {
		if isTwoPair(input) {
			return 2
		}
		return 1
	} else {
		return 0
	}
}

func getRankingJoker(input map[string]int, joker int) int {
	oak := ofAKind(input)
	oak += joker
	if oak == 5 {
		return 6
	} else if oak == 4 {
		return 5
	} else if oak == 3 {
		if isHouseJoker(input, joker) {
			return 4
		}
		return 3
	} else if oak == 2 {
		if isTwoPairJoker(input, joker) {
			return 2
		}
		return 1
	} else {
		return 0
	}
}
func ofAKind(input map[string]int) int {
	kinds := 0
	for k, i := range input {
		if k != "J" {
			if i > kinds {
				kinds = i
			}
		}

	}
	return kinds
}
func isHouse(input map[string]int) bool {
	pair, trips := false, false
	for _, i := range input {
		if i == 2 {
			pair = true
		}
		if i == 3 {
			trips = true
		}
	}

	return pair && trips
}

func isTwoPair(input map[string]int) bool {
	pair, pair2 := false, false
	for _, i := range input {
		if i == 2 && !pair {
			pair = true
			continue
		}
		if i == 2 && pair {
			pair2 = true
		}
	}
	return pair && pair2
}

func isHouseJoker(input map[string]int, joker int) bool {
	fmt.Println(input, joker)
	for i1, i2 := range input {
		for j1, j2 := range input {
			if i1 == j1 {
				continue
			}
			if i1 == "J" {
				continue
			}
			if j1 == "J" {
				continue
			}
			if i2+joker == 3 && j2 == 2 {
				return true
			}
			if i2 == 2 && j2+joker == 3 {
				return true
			}
			if i2 == 3 && j2 == 2 {
				return true
			}
			if j2 == 3 && i2 == 2 {
				return true
			}
		}
	}

	return false
}

func isTwoPairJoker(input map[string]int, joker int) bool {
	for i1, i2 := range input {
		for j1, j2 := range input {
			if i1 == j1 {
				continue
			}
			if i2 == 2 && j2+joker == 2 {
				return true
			}
			if i2+joker == 2 && j2 == 2 {
				return true
			}
			if i2 == 2 && j2 == 2 {
				return true
			}
		}
	}

	return false
}
