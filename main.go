package main

import (
	"github.com/NanXiao/stack"
	"fmt"
)

func main() {
	s := stack.New()
	s.Push(0)
	s.Push(1)
	s.Pop()
	fmt.Println(s)
}
