package main

import "fmt"

type race struct {
	time     int
	distance int
}

func main() {
	//Time:        53     91     67     68
	//Distance:   250   1330   1081   1025
	races := make([]race, 0)

	races = append(races, race{time: 53, distance: 250})
	races = append(races, race{time: 91, distance: 1330})
	races = append(races, race{time: 67, distance: 1081})
	races = append(races, race{time: 68, distance: 1025})

	race2 := race{time: 53916768, distance: 250133010811025}
	//race2 := race{time: 71530, distance: 940200}

	part1 := 1

	for _, r := range races {
		part1 *= calculatePossibleWins(r.time, r.distance)
	}

	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", calculatePossibleWins(race2.time, race2.distance))
}

func calculatePossibleWins(time int, distance int) int {
	firstW := time - 1
	lastW := 1
	for i := 1; i <= firstW; i++ {
		dist := (time - i) * i
		if dist > distance {
			firstW = i
			break
		}
	}

	for i := time - 1; i >= lastW; i-- {
		dist := (time - i) * i
		if dist > distance {
			lastW = i
			break
		}
	}

	return lastW - firstW + 1
}
