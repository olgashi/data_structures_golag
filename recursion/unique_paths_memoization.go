package main

import (
	"fmt"
)

func uniquePaths(rows, cols int, pathM(map[string]int)) int {
	if rows == 1 || cols == 1 {
		return 1
	}

  _, ok := pathM[string(rows) + "-" + string(cols)]
	if !ok {
		pathM[string(rows) + "-" + string(cols)] = uniquePaths(rows - 1, cols, pathM) + uniquePaths(rows, cols - 1, pathM)
	}

	return pathM[string(rows) + "-" + string(cols)]
}

func main() {
	pathMap := make(map[string]int)
	fmt.Println(uniquePaths(3, 7, pathMap))
}