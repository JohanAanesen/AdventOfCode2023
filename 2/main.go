package main

import (
	"bufio"
	"fmt"
	"os"
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

	gameIdCounter := 0
	part2Counter := 0
	//Game 1: 2 red, 2 green; 6 red, 3 green; 2 red, 1 green, 2 blue; 1 red
	for r.Scan() {
		line := r.Text()
		tempArr := strings.Split(line, ":")
		var game int
		fmt.Sscanf(tempArr[0], "Game %d", &game)

		sets := strings.Split(tempArr[1], ";")
		valid := true
		r, g, b := 0, 0, 0
		for _, set := range sets {
			if !checkSetIsWithinBounds(set) {
				valid = false
			}

			tr, tg, tb := checkMinimumCubes(set)
			if tr > r {
				r = tr
			}
			if tg > g {
				g = tg
			}
			if tb > b {
				b = tb
			}
		}
		if valid {
			gameIdCounter += game
		} else {
			fmt.Println(tempArr[0])
		}
		fmt.Println(r, g, b, r*g*b)
		part2Counter += r * g * b
	}

	fmt.Println("Part1: ", gameIdCounter)
	fmt.Println("Part2: ", part2Counter)

}

func checkSetIsWithinBounds(input string) bool {

	red, green, blue := 12, 13, 14

	inputs := strings.Split(input, ",")

	for _, in := range inputs {
		number, color := 0, ""
		fmt.Sscanf(in, " %d %s", &number, &color)
		if color == "red" && number > red {
			return false
		}
		if color == "green" && number > green {
			return false
		}
		if color == "blue" && number > blue {
			return false
		}
	}
	//fmt.Println(inputs)

	return true
}

func checkMinimumCubes(input string) (int, int, int) {

	red, green, blue := 0, 0, 0

	inputs := strings.Split(input, ",")
	//probably very overkill but its copypaste from part1 so works
	for _, in := range inputs {
		number, color := 0, ""
		fmt.Sscanf(in, " %d %s", &number, &color)
		if color == "red" && number > red {
			red = number
		}
		if color == "green" && number > green {
			green = number
		}
		if color == "blue" && number > blue {
			blue = number
		}
	}

	return red, green, blue
}
