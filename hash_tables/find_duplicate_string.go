package main

import (
	"fmt"
)
// time complexity of O(N)
func findDuplicateString(strSlice []string) string {
	var stringsMap = make(map[string]bool)

	for _, str := range strSlice {
		if _, ok := stringsMap[str]; ok {
			return str
		}

		stringsMap[str] = true
	}
	return ""
}

func main() {
	fmt.Println(findDuplicateString([]string{"a", "b", "c", "d", "c", "e", "f"})) // "c"
	fmt.Println(findDuplicateString([]string{"a", "b", "c", "d", "e", "f"})) //""
	fmt.Println(findDuplicateString([]string{"a", "b", "c", "d", "e", "f", "a"})) //"a"

}