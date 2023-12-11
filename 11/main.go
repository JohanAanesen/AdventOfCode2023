package main

import (
	"bufio"
	"fmt"
	"math"
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

	galaxies := make([]coordinate, 0)

	y := 0
	for r.Scan() {
		line := r.Text()
		for i, i2 := range line {
			if i2 == '#' {
				galaxies = append(galaxies, coordinate{x: i, y: y})
			}
		}
		y++
	}

	expanded := expand(galaxies, 2)

	sum := sumOfDistances(expanded)

	fmt.Println("Part1: ", sum)

	expanded2 := expand(galaxies, 1000000)
	sum2 := sumOfDistances(expanded2)
	fmt.Println("Part2: ", sum2)

}
func sumOfDistances(galaxies []coordinate) int {
	var sum int
	for i := range galaxies {
		for j := i + 1; j < len(galaxies); j++ {
			sum += distance(galaxies[i], galaxies[j])
		}
	}
	return sum
}
func distance(a, b coordinate) int {
	return abs(a.y-b.y) + abs(a.x-b.x)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func expand(galaxies []coordinate, factor int) []coordinate {
	columnsWithGalaxies := make(map[int]bool)
	rowsWithGalaxies := make(map[int]bool)

	for _, galaxy := range galaxies {
		columnsWithGalaxies[galaxy.x] = true
		rowsWithGalaxies[galaxy.y] = true
	}

	minCol, maxCol := keyRange(columnsWithGalaxies)
	minRow, maxRow := keyRange(rowsWithGalaxies)

	var columnsWithoutGalaxies []int
	for col := minCol; col <= maxCol; col++ {
		if !columnsWithGalaxies[col] {
			columnsWithoutGalaxies = append(columnsWithoutGalaxies, col)
		}
	}

	var rowsWithoutGalaxies []int
	for row := minRow; row <= maxRow; row++ {
		if !rowsWithGalaxies[row] {
			rowsWithoutGalaxies = append(rowsWithoutGalaxies, row)
		}
	}

	var expanded []coordinate

	for _, g := range galaxies {
		e := g
		for _, row := range rowsWithoutGalaxies {
			if row > g.y {
				break
			}
			e.y += factor - 1
		}
		for _, col := range columnsWithoutGalaxies {
			if col > g.x {
				break
			}
			e.x += factor - 1
		}
		expanded = append(expanded, e)
	}

	return expanded
}

func keyRange(m map[int]bool) (int, int) {
	min, max := math.MaxInt, math.MinInt
	for key := range m {
		if key < min {
			min = key
		}
		if key > max {
			max = key
		}
	}
	return min, max
}
