package num_Fibonacci

import (
	"fmt"

	"github.com/GoCourse/mypackage"
)

// Numfibonaci функция выводит заданное пользователем колличество первых чисел ряда чисел Фибоначчи.
func Numfibonaci() {

	n := mypackage.UserInputInt("Какое колличество первых чисел ряда Фибоначчи нужно вывести на экран?")

	// Сдесь я использовал float64 т.к. другие типы давали неточный результат.
	var fib1, fib2 float64 = 0, 1
	for i := 0; i < n; i++ {
		// Оставил чтобы показать ход мыслей.
		//fib1 = fib2
		//fib2 = fib_sum
		//fib_sum = fib1 + fib2
		// Можно сократить до такой записи.
		fmt.Printf("%d %.f\n", i+1, fib1)
		fib1, fib2 = fib2, fib1+fib2

	}

}
