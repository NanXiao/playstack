package main

import (
	"fmt"
	"github.com/NanXiao/stack"
)

func main() {
	s := stack.New()
	s.Push(0)
	s.Push(1)
	s.Pop()
	fmt.Println(s)
}
