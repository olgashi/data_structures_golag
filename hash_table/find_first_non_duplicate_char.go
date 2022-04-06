package main

import (
	"fmt"
)

func findFirstNonDuplicate(str string) string {
  charCounts := make(map[string]int)

	for _, char := range str {
		charCounts[string(char)]++
	}

	for _, char := range str {
		if val, ok := charCounts[string(char)]; ok && val == 1 {
      return string(char)
		}
	}
	return ""
}


func main() {
  fmt.Println(findFirstNonDuplicate("minimum")) //n
}