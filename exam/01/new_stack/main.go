package main

import (
	"fmt"
	"new_stack/stack" // Убедитесь, что здесь правильный путь
)

func main() {
	s := stack.NewStack()
	s.Push(10)
	s.Push(20)
	fmt.Println(s.Pop()) // 20
	fmt.Println(s.Pop()) // 10
	fmt.Println(s.Pop()) // nil
}
