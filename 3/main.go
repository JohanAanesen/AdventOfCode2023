package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	partGrid := make(map[coordinate]int, 0)
	numberGrid := make(map[coordinate]int, 0)
	numberGrid2 := make(map[coordinate]int, 0)
	numbersStored := make(map[coordinate]string, 0)

	gearsLocations := make([]coordinate, 0)
	y := 0
	for r.Scan() {
		line := r.Text()
		input := strings.Split(line, "")
		for x, s := range input {
			_, err := strconv.Atoi(s)
			if err != nil {
				if s != "." {
					partGrid[coordinate{x: x, y: y}] = 1
					partGrid[coordinate{x: x, y: y + 1}] = 1
					partGrid[coordinate{x: x, y: y - 1}] = 1
					partGrid[coordinate{x: x + 1, y: y}] = 1
					partGrid[coordinate{x: x - 1, y: y}] = 1
					partGrid[coordinate{x: x + 1, y: y + 1}] = 1
					partGrid[coordinate{x: x + 1, y: y - 1}] = 1
					partGrid[coordinate{x: x - 1, y: y + 1}] = 1
					partGrid[coordinate{x: x - 1, y: y - 1}] = 1

					if s == "*" {
						gearsLocations = append(gearsLocations, coordinate{x: x, y: y})
					}
				}
			} else {
				numberGrid[coordinate{x: x, y: y}] = 1
				numberGrid2[coordinate{x: x, y: y}] = 1
				numbersStored[coordinate{x: x, y: y}] = s
			}
		}
		y++
	}

	validNumbers := make([]int, 0)
	sum := 0

	for i := 0; i < 140; i++ {
		for j := 0; j < 140; j++ {
			if numberGrid[coordinate{x: i, y: j}] == 1 && partGrid[coordinate{x: i, y: j}] == 1 {
				num := numbersStored[coordinate{x: i, y: j}]
				if numbersStored[coordinate{x: i - 1, y: j}] != "" {
					num = numbersStored[coordinate{x: i - 1, y: j}] + num
					numberGrid[coordinate{x: i - 1, y: j}] = 0
					if numbersStored[coordinate{x: i - 2, y: j}] != "" {
						num = numbersStored[coordinate{x: i - 2, y: j}] + num
						numberGrid[coordinate{x: i - 2, y: j}] = 0
					}
				}

				if numbersStored[coordinate{x: i + 1, y: j}] != "" {
					num = num + numbersStored[coordinate{x: i + 1, y: j}]
					numberGrid[coordinate{x: i + 1, y: j}] = 0
					if numbersStored[coordinate{x: i + 2, y: j}] != "" {
						num = num + numbersStored[coordinate{x: i + 2, y: j}]
						numberGrid[coordinate{x: i + 2, y: j}] = 0
					}
				}
				n2s, _ := strconv.Atoi(num)
				validNumbers = append(validNumbers, n2s)
				sum += n2s
			}
		}
	}

	sum2 := 0

	for _, gear := range gearsLocations {

		//finn 2 tall som er "adjacent"
		//finn tall rundt, null ut tallet fra søket
		//finn nytt tall
		//gang de sammen
		//legg til i ny sum
		aroundCords := []coordinate{
			{-1, -1}, {0, -1}, {1, -1},
			{-1, 0}, {0, 0}, {1, 0},
			{-1, 1}, {0, 1}, {1, 1},
		}

		x, y := gear.x, gear.y
		numbersFound := make([]int, 0)
		for _, cord := range aroundCords {
			if numberGrid2[coordinate{x: x + cord.x, y: y + cord.y}] == 1 {
				newX, newY := x+cord.x, y+cord.y
				//found number
				//traverse to find full number
				num := numbersStored[coordinate{x: newX, y: newY}]
				if numbersStored[coordinate{x: newX - 1, y: newY}] != "" {
					num = numbersStored[coordinate{x: newX - 1, y: newY}] + num
					numberGrid2[coordinate{x: newX - 1, y: newY}] = 0
					if numbersStored[coordinate{x: newX - 2, y: newY}] != "" {
						num = numbersStored[coordinate{x: newX - 2, y: newY}] + num
						numberGrid2[coordinate{x: newX - 2, y: newY}] = 0
					}
				}

				if numbersStored[coordinate{x: newX + 1, y: newY}] != "" {
					num = num + numbersStored[coordinate{x: newX + 1, y: newY}]
					numberGrid2[coordinate{x: newX + 1, y: newY}] = 0
					if numbersStored[coordinate{x: newX + 2, y: newY}] != "" {
						num = num + numbersStored[coordinate{x: newX + 2, y: newY}]
						numberGrid2[coordinate{x: newX + 2, y: newY}] = 0
					}
				}
				n2s, _ := strconv.Atoi(num)
				numbersFound = append(numbersFound, n2s)
			}
		}
		if len(numbersFound) == 2 { //kun gir med 2 tall skal telles opp
			sum2 += numbersFound[0] * numbersFound[1]
		}

	}

	fmt.Println("Part1: ", sum)
	fmt.Println("Part2: ", sum2)

}
