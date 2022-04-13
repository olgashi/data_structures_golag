package main

import (
	"fmt"
)

func addUntil100(numList []int) int {
	if len(numList) == 0 {
		return 0
	}
  sumNumsAfterTheFirst := addUntil100(numList[1:])
	if numList[0] + sumNumsAfterTheFirst > 100 {
		return sumNumsAfterTheFirst
	} else {
		return numList[0] + sumNumsAfterTheFirst
	}
}

func main() {
	fmt.Println(addUntil100([]int{1,2,3,4,5,6,77,77,1,2,3,4,6,7,8,22,3,5,1,3,3,44,1,2}))
}