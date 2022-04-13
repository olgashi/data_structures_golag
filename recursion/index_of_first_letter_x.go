package main
import (
	"fmt"
)
// lowercase only

func indexOfFirstLetterX(inputString string, index int) int {
	if len(inputString) == 1 {
		if string(inputString[0]) == "x" {
			return index
		} else {
			return -1
		}
	}

  if string(inputString[0]) == "x" {
		return index
	} else {
		index++
	}
	return indexOfFirstLetterX(inputString[1:], index)
}

func main() {
  fmt.Println(indexOfFirstLetterX("abcdefghijklmnopqrstuvwxyz", 0)) // 6
}