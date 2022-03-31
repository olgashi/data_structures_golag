package main

import (
	"fmt"
)


func findGreatestInt(numList []int) (greatest int) {
	for _, num := range numList {
		if num > greatest {
			greatest = num
		}
	}
	return greatest
}

func main() {
	fmt.Println(findGreatestInt([]int{443, 200, 1, 3233, 5, 122, 323, 323, 2, 1, 1000}))
}