package main

import (
	"fmt"
)

func sum(low, high int) int {
	if low > high {
		return 0
	}

	return low + sum(low + 1, high)
}

func main() {
	fmt.Println(sum(1, 10))
}