package main

import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 1; x< 11; x++ {
			naturals <- x
		}
		close(naturals)
	}()


	// возведение в квадрат
	go func() {
		for a := range naturals {
			squares <- a * a
		}
		close(squares)
	}()

	// печать
	for {
		res, ok := <-squares
		if !ok {
			break // канал закрыт и пуст
		}
		fmt.Println(res)
	}
}