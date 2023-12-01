package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	totalPaper := 0
	totalRibbon := 0

	for r.Scan() {
		line := r.Text()
		sides := strings.Split(line, "x")

		l, _ := strconv.Atoi(sides[0])
		w, _ := strconv.Atoi(sides[1])
		h, _ := strconv.Atoi(sides[2])

		ls := l * w
		ws := w * h
		hs := h * l

		smallestSide := ls

		if ws < smallestSide {
			smallestSide = ws
		}
		if hs < smallestSide {
			smallestSide = hs
		}

		totalPaper += (ls * 2) + (hs * 2) + (ws * 2) + smallestSide

		lenArr := []int{l, w, h}
		sort.Ints(lenArr)

		totalRibbon += lenArr[0]*2 + lenArr[1]*2 + (l * w * h)

	}

	fmt.Println("Part1: ", totalPaper)
	fmt.Println("Part2: ", totalRibbon)

}
