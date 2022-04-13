package main
import (
	"fmt"
)

func totalNumberOfCharsInListOfWords(wordsList []string) int {
	if len(wordsList) == 1 {
		return len(wordsList[0])
	}

	return len(wordsList[0]) + totalNumberOfCharsInListOfWords(wordsList[1:])
}


func main() {
  fmt.Println(totalNumberOfCharsInListOfWords([]string{"ab", "c", "def", "ghij"}))
}