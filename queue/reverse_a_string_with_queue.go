package main

import (
	"fmt"
)

func reverseString(str string) string {
	// abcdef
  queue := ""
	for _, char := range str {
    queue = string(char) + queue
	}
	return queue
}

func main() {
  fmt.Println(reverseString("abcdefgh"))
}