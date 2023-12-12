package main

import "fmt"

func main() {
	input := "cqjxxyzz" //"cqjxjnds"

	iteratePassword(input)
}

func iteratePassword(password string) string {
	tempPw := make([]rune, len(password))
	for i, i2 := range password {
		tempPw[i] = i2
	}
	i := len(password) - 1
	found := false
	for !found {
		if tempPw[i] >= 'z' {
			for j := len(password) - 1; j > 0; j-- {
				tempPw[j]++
				if tempPw[j] > 'z' {
					tempPw[j] = 'a'
				} else {
					break
				}
			}
		} else {
			tempPw[i] = tempPw[i] + 1
		}
		//fmt.Println(string(tempPw))
		if checkIncreasing(string(tempPw)) && checkNoIllegals(string(tempPw)) && checkTwoDoubles(string(tempPw)) {
			found = true
		}
	}

	fmt.Println(string(tempPw))
	return string(tempPw)
}

func checkIncreasing(password string) bool {
	tempChar := password[0]
	counter := 1
	for i := 1; i < len(password); i++ {

		if password[i] == tempChar+1 {
			counter++
			tempChar = password[i]
		} else {
			counter = 1
			tempChar = password[i]
		}

		if counter == 3 {
			return true
		}
	}

	return false
}

func checkNoIllegals(password string) bool {
	for _, character := range password {
		if character == 'i' || character == 'o' || character == 'l' {
			return false
		}
	}
	return true
}

func checkTwoDoubles(password string) bool {
	tempMap := make(map[rune]int, 0)
	dubs := 0
	for i := 0; i < len(password)-1; i++ {
		tempMap[rune(password[i])]++
		if password[i] == password[i+1] {
			dubs++
			i++
		}
	}
	quads := false
	for _, i := range tempMap {
		if i >= 4 {
			quads = true
		}
	}
	return dubs == 2 && !quads
}
