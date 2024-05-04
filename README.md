**Stack Data Structure Golang**
====================================
Stack is a linear data structure which follows the LIFO(Last In First Out) policy. You can only access or modify the most recently item added.

**Get started**
===================
## Installation

With the help of below Go command, simply install the mystack package
```bash
 import "github.com/afzalansari028/go-stack/mystack"
```
## Example:
```bash
 package main

import (
	"fmt"

	"github.com/afzalansari028/go-stack/mystack"
)

func main() {

	fmt.Println("let's play with stack using golang")

	stack := mystack.Stack{}

	stack.Push(10) // push an elemment into the stack int type
	stack.Push("a")
	stack.Display()
	stack.Pop() // remomve an elememnt from the stack

	stack.Display()                            //display all elements
	fmt.Println("top element::", stack.Peek()) // see the top element

	fmt.Println("isempty::", stack.IsEmpty())
	fmt.Println("size::", stack.Size())
}

```
 Use Go command to run the code
```bash
$ go run main.go
```


