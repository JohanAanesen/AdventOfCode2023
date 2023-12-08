package main

import (
	"bufio"
	"fmt"
	"os"
)

type node struct {
	left  string
	right string
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	instructions := "LLLLLLLRRRLRRRLRLRLRLRRLLRRRLRLLRRRLLRRLRRLRRLRLRRLRLRRRLRRLRLRRRLRRLRRLRLRRLLRLLRLRRRLRRLRLLLLRRLLLLRLRLRLRRRLRLRLLLRLRRRLRRRLRRRLRLRRLRRRLRLLLRLLRRLRRRLRRLRRLRRLRLRRRLRLRLRLLRRRLRRRLRRLRRRLLLRRLRRLRRRLRLRRRLRRRLRLRRLRRRLRLRRLRLRRLRRRLRLRRLRLLRRRLLRLRRLRRRLLLRLRRLRRRR"
	//instSplit := strings.Split(instructions, "")
	nodeMap := make(map[string]node, 0)

	for r.Scan() {
		line := r.Text()

		nodeS := line[0:3]
		left := line[7:10]
		right := line[12:15]
		//fmt.Println(nodeS, left, right)
		nodeMap[nodeS] = node{left: left, right: right}

	}

	currentNode := "AAA"
	found := false
	iterations := 0
	curIt := 0
	for !found {
		for i, i2 := range instructions {
			if currentNode == "ZZZ" {
				found = true
				curIt = iterations + i
				break
			}

			if i2 == 'L' {
				currentNode = nodeMap[currentNode].left
			} else if i2 == 'R' {
				currentNode = nodeMap[currentNode].right
			}
		}
		iterations += len(instructions)
	}
	fmt.Println("Part 1: ", curIt)

	startingNodes := make([]string, 0)
	for s, _ := range nodeMap {
		if s[2] == 'A' {
			startingNodes = append(startingNodes, s)
		}
	}

	steps := make([]int, len(startingNodes))
	for i, cur := range startingNodes {
		for cur[2] != 'Z' {
			for _, turn := range instructions {
				if turn == 'L' {
					cur = nodeMap[cur].left
				} else {
					cur = nodeMap[cur].right
				}

				steps[i]++

				if cur[2] == 'Z' {
					break
				}
			}
		}
	}

	factors := make([]int, 0)
	for _, s := range steps {
		for _, f := range PrimeFactors(s) {
			if !sliceContains(factors, f) {
				factors = append(factors, f)
			}
		}
	}

	prod := 1
	for _, f := range factors {
		prod *= f
	}

	fmt.Println("Part 2: ", prod)
}

func sliceContains(slice []int, nr int) bool {
	m := make(map[int]int, 0)
	for _, i2 := range slice {
		m[i2]++
	}

	return m[nr] > 0
}

// stolen
func PrimeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}

	// n must be odd at this point. so we can skip one element
	// (note i = i + 2)
	for i := 3; i*i <= n; i = i + 2 {
		// while i divides n, append i and divide n
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}

	// This condition is to handle the case when n is a prime number
	// greater than 2
	if n > 2 {
		pfs = append(pfs, n)
	}

	return
}
