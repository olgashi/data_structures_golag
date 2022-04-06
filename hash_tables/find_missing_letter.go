package main

import (
	"fmt"
)
// time complexity of O(N)
func findMissingLetter(str string) (char string) {
	abc := map[string]int{"a":0, "b":0, "c":0, "d":0, "e":0, "f":0, "g":0, "h":0, "i":0, "j":0, "k":0, "l":0,
	                      "m":0, "n":0, "o":0, "p":0, "q":0, "r":0, "s":0, "t":0, "u":0, "v":0, "w":0, "x":0, "y":0, "z":0}

  for _, char := range str {
		abc[string(char)]++
	}

	for k, v := range abc {
		if v == 0 {
			return k
		}
	}
	return ""
}

func main() {
  fmt.Println(findMissingLetter("the quick borwn box jumps over a lazy dog"))
}

