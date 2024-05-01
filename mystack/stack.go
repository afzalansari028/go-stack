package mystack

import "fmt"

type Stack struct {
	data []int
}

func (stack *Stack) Size() int {
	size := len(stack.data)
	return size
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

// add an element into the stack
func (stack *Stack) Push(val int) {
	if stack.Size() == 5 {
		fmt.Println("STACK IS FULL")
		return
	}
	stack.data = append(stack.data, val)
}

// remove the element from top
func (stack *Stack) Pop() int {
	if stack.Size() == 0 {
		fmt.Println("STACK IS EMPTY")
		return 0
	}
	rv := len(stack.data) - 1
	stack.data = stack.data[:len(stack.data)-1]

	return rv
}

func (stack *Stack) Peek() int {
	if stack.Size() == 0 {
		fmt.Println("STACK IS EMPTY")
		return 0
	}
	rv := stack.data[len(stack.data)-1]
	return rv
}

func (stack *Stack) Display() {

	for i := len(stack.data) - 1; i >= 0; i-- {
		fmt.Println(stack.data[i])
	}
	fmt.Println("---END---")
}
