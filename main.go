package main

import (
	"fmt"
	"math/rand"

	"github.com/turbo-d/go-dsa/stack"
)

func main() {
	var s stack.SliceStack

	printStack(s)

	for i := 0; i < 5; i++ {
		random := rand.Intn(100)
		fmt.Println()
		fmt.Println("Push: ", random)
		s.Push(random)
		printStack(s)
	}
}

func printStack(s stack.SliceStack) {
	fmt.Println()
	fmt.Println("Stack:")
	fmt.Println(s)
	fmt.Println("size: ", s.Size())
}
