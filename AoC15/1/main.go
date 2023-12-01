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

	floor := 0
	found := false

	for r.Scan() {
		line := r.Text()
		instructions := strings.Split(line, "")

		for i, instruction := range instructions {
			if instruction == "(" {
				floor++
			} else if instruction == ")" {

				floor--
				if floor == -1 && !found {
					fmt.Println("Part2: ", i+1)
					found = true
				}
			}
		}

	}

	fmt.Println("Part1: ", floor)

}
