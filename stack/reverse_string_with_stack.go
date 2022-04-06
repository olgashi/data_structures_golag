package main

import (
	"fmt"
)


func reverseString(str string) (reversedString string) {
  stack := []string{}

	for _, el := range str {
		stack = append(stack, string(el))
	}
  
	for index := len(stack) - 1; index >= 0; index-- {
    reversedString += string(stack[index])
	}

	return
}

func main() {
  fmt.Println(reverseString("abcdefgh"))
  
}