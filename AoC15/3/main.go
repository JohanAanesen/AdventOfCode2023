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

	visitedHouses := make(map[coordinate]int, 0)

	for r.Scan() {
		line := r.Text()
		input := strings.Split(line, "")

		x, y, x2, y2 := 0, 0, 0, 0
		for i, direction := range input {
			if i%2 == 0 {
				switch direction {
				case "^":
					y++
				case "v":
					y--
				case "<":
					x--
				case ">":
					x++
				}
				visitedHouses[coordinate{x: x, y: y}] = 1
			} else {
				switch direction {
				case "^":
					y2++
				case "v":
					y2--
				case "<":
					x2--
				case ">":
					x2++
				}
				visitedHouses[coordinate{x: x2, y: y2}] = 1
			}

		}
	}

	sum := 0
	for _, val := range visitedHouses {
		sum += val
	}
	fmt.Println("Part1: ", sum)

}
