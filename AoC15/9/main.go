package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type route struct {
	from, to string
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	routes := make(map[route]int, 0)
	cities := make([]string, 0)

	for r.Scan() {
		line := r.Text()

		tempRoute := route{}
		len := 0

		fmt.Sscanf(line, "%s to %s = %d", &tempRoute.from, &tempRoute.to, &len)
		tempRoute2 := route{from: tempRoute.to, to: tempRoute.from}
		cities = append(cities, tempRoute.from)
		cities = append(cities, tempRoute.to)
		routes[tempRoute] = len
		routes[tempRoute2] = len

	}

	cities = removeDuplicateStr(cities)
	shortest := math.MaxInt
	longest := 0

	for _, v := range permutations(cities) {
		temp := 0
		for i := 0; i < len(v)-1; i++ {
			temp += routes[route{v[i], v[i+1]}]
		}
		if temp < shortest {
			shortest = temp
		}

		if temp > longest {
			longest = temp
		}

		//fmt.Println(v, temp)
	}

	fmt.Println("Part 1: ", shortest)
	fmt.Println("Part 2: ", longest)
}

func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func removeDuplicateStr(strSlice []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}
	for _, item := range strSlice {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
