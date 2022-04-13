package main
import (
	"fmt"
)

func numberOfShirtestPaths(rows, cols int) int {
  if cols == 1 || rows == 1 {
		return 1
	}
	return  numberOfShirtestPaths(rows - 1, cols) + numberOfShirtestPaths(rows, cols - 1)
}


func main() {
  fmt.Println(numberOfShirtestPaths(2, 6))
}