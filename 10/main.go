package main

import (
	"bufio"
	"fmt"
	"os"
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

	inputMap := make(map[coordinate]rune, 0)
	distMap := make(map[coordinate]int, 0)
	startingPoint := coordinate{}
	y := 0
	for r.Scan() {
		line := r.Text()
		for i, i2 := range line {
			inputMap[coordinate{x: i, y: y}] = i2
			if i2 == 'S' {
				startingPoint.x = i
				startingPoint.y = y
			}
		}
		y++
	}

	steps := 1
	direction := "down"
	currentCoord := coordinate{x: startingPoint.x, y: startingPoint.y + 1}
	inputMap[startingPoint] = '7'
	distMap[currentCoord] = 1
	for i := 0; i < 100000; i++ {
		if currentCoord.x == startingPoint.x && currentCoord.y == startingPoint.y {
			fmt.Println("Part1: ", steps/2)
			break
		}
		currentStep := inputMap[currentCoord]
		if currentStep == '|' {
			if direction == "down" {
				currentCoord.y++
			} else {
				currentCoord.y--
			}
		} else if currentStep == '-' {
			if direction == "left" {
				currentCoord.x--
			} else if direction == "right" {
				currentCoord.x++
			}
		} else if currentStep == 'L' {
			if direction == "left" {
				direction = "up"
				currentCoord.y--
			} else if direction == "down" {
				direction = "right"
				currentCoord.x++
			}
		} else if currentStep == 'J' {
			if direction == "right" {
				direction = "up"
				currentCoord.y--
			} else if direction == "down" {
				direction = "left"
				currentCoord.x--
			}
		} else if currentStep == '7' {
			if direction == "right" {
				direction = "down"
				currentCoord.y++
			} else if direction == "up" {
				direction = "left"
				currentCoord.x--
			}
		} else if currentStep == 'F' {
			if direction == "left" {
				direction = "down"
				currentCoord.y++
			} else if direction == "up" {
				direction = "right"
				currentCoord.x++
			}
		}
		distMap[currentCoord] = steps
		steps++

	}
	insideCounter := 0
	for i := 0; i < 140; i++ {
		temps := ""
		isInside := false
		for j := 0; j < 140; j++ {
			if distMap[coordinate{x: j, y: i}] != 0 {
				temps += "x"
				if inputMap[coordinate{x: j, y: i}] == 'L' || inputMap[coordinate{x: j, y: i}] == '|' || inputMap[coordinate{x: j, y: i}] == 'J' {
					isInside = !isInside
				}
			} else {
				if isInside {
					temps += "1"
					insideCounter++
				} else {
					temps += "0"
				}

			}
		}
		fmt.Println(temps)
	}
	fmt.Println("Part 2: ", insideCounter)
}
