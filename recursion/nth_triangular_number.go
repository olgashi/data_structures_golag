package main
import (
	"fmt"
)

func returnNthTriangularNumber(n int, index int, nextNum int) int {
	if (index >= n) {
		return nextNum
	}
	index++

	return returnNthTriangularNumber(n, index, nextNum+index)
}


func main() {
  fmt.Println(returnNthTriangularNumber(3, 0, 0)) // 6
  fmt.Println(returnNthTriangularNumber(7, 0, 0)) // 28
}