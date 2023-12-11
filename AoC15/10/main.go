package main

import (
	"fmt"
	"strings"
)

type route struct {
	from, to string
}

func main() {
	input := "3113322113"
	inputSplit := strings.Split(input, "")

	for i := 0; i < 50; i++ {

		newS := lookSay(inputSplit)

		inputSplit = strings.Split(newS, "")
		fmt.Println(i, len(newS))
	}

}

func lookSay(s []string) string {
	if len(s) < 1 {
		return ""
	}
	temp := s[0]

	length := 0
	for i := 0; i < len(s); i++ {
		if s[i] == temp {
			length++
		} else {
			break
		}
	}
	return fmt.Sprint(length) + temp + lookSay(s[length:])
}
