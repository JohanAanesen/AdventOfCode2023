package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "yzbqklnj"
	for i := 0; i < 100000000; i++ {
		tempString := GetMD5Hash(input + strconv.Itoa(i))
		if strings.HasPrefix(tempString, "00000") {
			fmt.Println("Part1: ", i)
			break
		}
	}
	for i := 0; i < 100000000; i++ {
		tempString := GetMD5Hash(input + strconv.Itoa(i))
		if strings.HasPrefix(tempString, "000000") {
			fmt.Println("Part2: ", i)
			break
		}
	}
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
