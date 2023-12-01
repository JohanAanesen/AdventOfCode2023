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

	lights := make([][]int, 1000)
	lights2 := make([][]int, 1000)

	for i := 0; i < 1000; i++ {
		lights[i] = make([]int, 1000)
		lights2[i] = make([]int, 1000)
	}

	for r.Scan() {
		line := r.Text()
		if strings.HasPrefix(line, "turn on ") {
			var x1, y1, x2, y2 int

			fmt.Sscanf(line, "turn on %d,%d through %d,%d\n", &x1, &y1, &x2, &y2)

			for i := x1; i <= x2; i++ {
				for y := y1; y <= y2; y++ {
					lights[i][y] = 1
					lights2[i][y] += 1
				}
			}
		} else if strings.HasPrefix(line, "turn off ") {
			var x1, y1, x2, y2 int

			fmt.Sscanf(line, "turn off %d,%d through %d,%d\n", &x1, &y1, &x2, &y2)
			for i := x1; i <= x2; i++ {
				for y := y1; y <= y2; y++ {
					lights[i][y] = 0
					if lights2[i][y] > 0 {
						lights2[i][y] -= 1
					}
				}
			}
		} else if strings.HasPrefix(line, "toggle ") {
			var x1, y1, x2, y2 int

			fmt.Sscanf(line, "toggle %d,%d through %d,%d\n", &x1, &y1, &x2, &y2)

			for i := x1; i <= x2; i++ {
				for j := y1; j <= y2; j++ {
					if lights[i][j] == 1 {
						lights[i][j] = 0
					} else {
						lights[i][j] = 1
					}
					lights2[i][j] += 2
				}
			}
		}

	}
	sum := 0
	sum2 := 0
	for i := 0; i < 1000; i++ {
		for y := 0; y < 1000; y++ {
			sum += lights[i][y]
			sum2 += lights2[i][y]
			if lights2[i][y] < 0 {
				fmt.Println(lights2[i][y])
			}
		}
	}

	fmt.Println("Part1: ", sum)
	fmt.Println("Part2: ", sum2)

}
