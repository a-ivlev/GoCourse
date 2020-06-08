package main

import (
	"fmt"

	"github.com/GoCourse/homework-3/addresBook"
	"github.com/GoCourse/homework-3/queue"
	"github.com/GoCourse/homework-3/stractAuto"
)

var task int

func main() {
	fmt.Println("Выберите задание для проверки\n")
	fmt.Println("1 Задание 1 и 2\n3 Задание 3\n4 Задание 4\n")
	fmt.Scan(&task)
	if task == 1 {
		fmt.Println("Задание 1 и 2. Описать несколько структур — любой легковой автомобиль и грузовик.\n")
		stractAuto.StractAuto()
	}
	if task == 3 {
		fmt.Println("Задание 3. Реализовать очередь по принципу FIFO(первый пришел первый ушел).\n")
		queue.Push("Этот текст", "Будет находиться в очереди", "До первого обращения к pop")
		n := 1
		for i := 0; i < n; i++ {
			element, err := queue.Pop()
			fmt.Println(element)
			if err == nil {
				n++
			}
		}
	}
	if task == 4 {
		fmt.Println("Задание 4. Внести в телефонный справочник дополнительные данные. Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.\n")
		addresBook.AddresBook()
	}

}
