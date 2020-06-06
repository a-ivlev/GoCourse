package main

import (
	"fmt"
	"github.com/GoCourse/homework-3/queue"
)

func main() {

	queue.Push("Этот текст", "Будет находиться в очереди", "До первого обращения к pop")
	n := 1
	for i :=0; i < n; i++{
		element, err := queue.Pop()
		fmt.Println(element)
		if err == nil {
			n++
		}
	}

}
