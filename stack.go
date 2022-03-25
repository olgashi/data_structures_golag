package main
import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(val int) {
  s.items = append(s.items, val)
}

func (s *Stack) Pop() int {
	stackLength := len(s.items)
	removedItem := s.items[stackLength - 1]
  s.items = s.items[:stackLength - 1]
	return removedItem
}

func main() {
	stack := Stack{}

	fmt.Println(stack.items)
	stack.Push(6)
	stack.Push(2)
	stack.Push(10)

	fmt.Println(stack.items)
	
	fmt.Println(stack.Pop())
	fmt.Println(stack.items)
}