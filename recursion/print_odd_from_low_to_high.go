package main

import (
	"fmt"
)

func print_every_other(low, high int) {
	if low > high {
		return
	}

	fmt.Println(low)
	print_every_other(low + 2, high)
}

func main() {
	print_every_other(0, 10)
} 