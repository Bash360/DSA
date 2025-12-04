package main

import (
	"fair-money/stack"
	"fmt"
)

func main() {
	tasks := stack.New[string]("go to bank", "wash clothes", "cook")
	tasks.Push("go to the hospital")
	tasks.Pop()

	fmt.Println(tasks.Size())
	fmt.Println(tasks.Peek())
	fmt.Println(tasks)

}
