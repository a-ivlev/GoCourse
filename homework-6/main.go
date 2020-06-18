package main

import (
	"fmt"
	"github.com/GoCourse/homework-6/line_img"
	"github.com/GoCourse/homework-6/statistic"
	"github.com/GoCourse/homework-6/сhess_board"
)

var task int

func main()  {
	for {
		fmt.Println("\nВыберите задание для проверки. Для выхода наюерите 0.")
		fmt.Println("1 Задание 1 Дополните код из раздела «Тестирование» функцией подсчета суммы переданных элементов и тестом для этой функции.")
		fmt.Println("2 Задание 2. Дополните пример из раздела «Пакет img» изображением горизонтальных и вертикальных линий.")
		fmt.Println("3 Задание 3. Дополните функцию hello() http-сервера так, чтобы принять и отобразить на странице один GET-параметр, например name.")
		fmt.Println("5 Задание 5. Напишите программу, генерирующую png-файл с рисунком шахматной доски.\n")

		fmt.Scan(&task)
		switch task {
		case 1:
			fmt.Println("Функция подсчёта суммы переданных элементов. Передали 1, 2, 3, 4, 5, 6")
			fmt.Println("Получили: ", statistic.SumArg([]float64{1,2,3,4,5,6}))
		case 2:
			line_img.LineImg()
		case 3:
			//Calculat()
		case 5:
			chess_board.ChessBoard()
		case 0:
			return
		}
	}
}