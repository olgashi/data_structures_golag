package main

import (
	"fmt"
)

func golomb(num int, golombM(map[int]int)) int {
	if num == 1 {
    return 1
	}
	_, ok := golombM[num]
	if !ok {
		golombM[num] = golomb(num - golomb(golomb(num - 1, golombM), golombM), golombM)
	}

	return 1 + golombM[num]
}

func main() {
	golombMap := make(map[int]int)
	fmt.Println(golomb(10, golombMap))
}