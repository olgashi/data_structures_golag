package main
import (
	"fmt"
)

func returnListOfEvenOnly(numList []int) []int {
	if len(numList) == 1 {
		if numList[0] % 2 == 0 {
			return []int{numList[0]}
		} else {
			return []int{}
		}
	}

	evenNumsList := returnListOfEvenOnly(numList[1:])

  if (numList[0] % 2 == 0) {
		evenNumsList = append(evenNumsList, numList[0])
	}

	return evenNumsList
}


func main() {
  fmt.Println(returnListOfEvenOnly([]int{1, 2, 3, 4, 5, 6, 7}))
}