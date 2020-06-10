package main

import (
	"fmt"

	"github.com/GoCourse/homework-3/addresBook"
	"github.com/GoCourse/homework-3/queue"
	"github.com/GoCourse/homework-3/stractAuto"
)

var task int

func main() {
	for {
		fmt.Println("Выберите задание для проверки. Для выхода наюерите 0.")
		fmt.Println("1 Задание 1 и 2. Описать несколько структур — любой легковой автомобиль и грузовик.")
		fmt.Println("3 Задание 3. Реализовать очередь по принципу FIFO (первый пришел первый ушел).")
		fmt.Println("4 Задание 4. Внести в телефонный справочник дополнительные данные. Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.\n")
		fmt.Scan(&task)
		switch task {
		case 1:
			stractAuto.StractAuto()
		case 3:
			queue.Queue()
		case 4:
			addresBook.AddresBook()
		case 0:
			return
		}
	}
}