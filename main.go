package main

import "github.com/afzalansari028/go-stack/mystack"

func main() {
	stack := mystack.Stack{}
	stack.Push("a")
	stack.Push("b")
	stack.Push("c")

	stack.Display()
}
