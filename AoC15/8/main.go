package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	charOfCode := make([]int, 0)
	charOfMem := make([]int, 0)

	charOfMem2 := make([]int, 0)

	for r.Scan() {
		line := r.Text()

		charOfCode = append(charOfCode, len(line))
		tempCounter := 0
		tempCounter2 := 2
		for i := 0; i < len(line); i++ {
			if line[i] == '"' {
				tempCounter2++
				tempCounter2++
			} else if line[i] == '\\' {
				tempCounter2++
				tempCounter2++
			} else {
				tempCounter2++
			}
			fmt.Println(tempCounter2)
		}
		charOfMem = append(charOfMem, tempCounter)
		charOfMem2 = append(charOfMem2, tempCounter2)
	}

	fmt.Println(charOfCode)
	fmt.Println(charOfMem)
	fmt.Println(charOfMem2)
	codeLen, memLen, memLen2 := 0, 0, 0

	for _, v := range charOfCode {
		codeLen += v
	}

	for _, v := range charOfMem {
		memLen += v
	}

	for _, v := range charOfMem2 {
		memLen2 += v
	}
	fmt.Println(memLen2, codeLen)
	fmt.Println("Part2: ", memLen2-codeLen)

}
