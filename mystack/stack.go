package mystack

import "fmt"

type Stack struct {
	Data []interface{}
}

const default_size = 10

func (stack *Stack) Size() int {
	size := len(stack.Data)
	return size
}

func (stack *Stack) IsEmpty() bool {
	return stack.Size() == 0
}

// add an element into the stack
func (stack *Stack) Push(val interface{}) {
	if stack.Size() == default_size {
		fmt.Println("STACK IS FULL")
		return
	}
	stack.Data = append(stack.Data, val)
}

// remove the element from top
func (stack *Stack) Pop() int {
	if stack.IsEmpty() {
		fmt.Println("STACK IS EMPTY")
		return 0
	}
	rv := len(stack.Data) - 1
	stack.Data = stack.Data[:len(stack.Data)-1]

	return rv
}

func (stack *Stack) Peek() interface{} {
	if stack.Size() == 0 {
		fmt.Println("STACK IS EMPTY")
		return 0
	}
	rv := stack.Data[len(stack.Data)-1]
	return rv
}

func (stack *Stack) Display() {

	for i := len(stack.Data) - 1; i >= 0; i-- {
		fmt.Println(stack.Data[i])
	}
}
