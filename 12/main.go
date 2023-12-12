package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type spring struct {
	input, rule string
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var records []string
	var groups [][]int

	for r.Scan() {
		line := r.Text()
		parts := strings.Split(line, " ")
		records = append(records, parts[0])
		var group []int
		for _, num := range strings.Split(parts[1], ",") {
			num, _ := strconv.Atoi(num)
			group = append(group, num)
		}
		groups = append(groups, group)
	}
	res := 0
	for i := range records {
		res += solve(records[i], groups[i])
	}

	fmt.Println("Part1: ", res)

	for i, record := range records {
		records[i] = unfoldRecord(record)
	}
	for i, group := range groups {
		groups[i] = unfoldGroup(group)
	}
	res2 := 0
	for i := range records {
		res2 += solve(records[i], groups[i])
	}

	fmt.Println("Part2: ", res2)
}

func solve(record string, group []int) int {
	var cache [][]int
	for i := 0; i < len(record); i++ {
		cache = append(cache, make([]int, len(group)+1))
		for j := 0; j < len(group)+1; j++ {
			cache[i][j] = -1
		}
	}

	return findSolutions(0, 0, record, group, cache)
}

func findSolutions(i, j int, record string, group []int, cache [][]int) int {
	if i >= len(record) {
		if j < len(group) {
			return 0
		}
		return 1
	}
	if cache[i][j] != -1 {
		return cache[i][j]
	}
	res := 0
	if record[i] == '.' {
		res = findSolutions(i+1, j, record, group, cache)
	} else {
		if record[i] == '?' {
			res += findSolutions(i+1, j, record, group, cache)
		}
		if j < len(group) {
			count := 0
			for k := i; k < len(record); k++ {
				if count > group[j] || record[k] == '.' || count == group[j] && record[k] == '?' {
					break
				}
				count += 1
			}

			if count == group[j] {
				if i+count < len(record) && record[i+count] != '#' {
					res += findSolutions(i+count+1, j+1, record, group, cache)
				} else {
					res += findSolutions(i+count, j+1, record, group, cache)
				}
			}
		}
	}

	cache[i][j] = res
	return res
}

func unfoldRecord(record string) string {
	var res strings.Builder
	for i := 0; i < len(record)*5; i++ {
		if i != 0 && i%len(record) == 0 {
			res.WriteByte('?')
		}
		res.WriteByte(record[i%len(record)])
	}

	return res.String()
}

func unfoldGroup(group []int) []int {
	var res []int
	for i := 0; i < len(group)*5; i++ {
		res = append(res, group[i%len(group)])
	}

	return res
}
