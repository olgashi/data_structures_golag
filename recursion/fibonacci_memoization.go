package main

import (
	"fmt"
)

func fib(num int, memo(map[int]int)) int {
	if num <= 1 {
		return num
	}
	_, ok := memo[num]

	if !ok {
		memo[num] = fib(num - 2, memo) + fib(num - 1, memo)
	}

	return memo[num]
}

func main() {
	mem := make(map[int]int)
  fmt.Println(fib(6, mem)) // 8
}

// 0 1 1 2 3 5 8 13