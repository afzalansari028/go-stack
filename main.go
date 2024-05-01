package main

import "github.com/afzalansari028/mystack/mystack"

func main() {
	stack := mystack.Stack{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	stack.Display()
}
