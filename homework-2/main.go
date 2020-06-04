package main

import (
	"fmt"

	"github.com/GoCourse/homework-2/even_number"
	"github.com/GoCourse/homework-2/num_Fibonacci"
	"github.com/GoCourse/homework-2/num_divisible3"
)

func main() {

	fmt.Println("1. Функция, определяет четное число или нечётное.\n")
	even_number.EvenNumber()

	fmt.Println("2. Функция определяет, делится число без остатка на 3 или нет.\n")
	num_divisible3.NumDivisible3()

	fmt.Println("3. Функция, которая последовательно выводит на экран заданное пользователем колличество первых чисел ряда Фибоначчи, начиная от 0.\n")
	num_Fibonacci.Numfibonaci()
}
