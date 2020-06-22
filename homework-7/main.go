package main

import (
	"fmt"
	server "github.com/GoCourse/homework-6/html"
	"github.com/GoCourse/homework-6/line_img"
	"github.com/GoCourse/homework-6/сhess_board"
	timeSpin "github.com/GoCourse/homework-7/spiner"
)

var task, Spin int
//var timeSpin time.Duration

func main()  {
	for {
		fmt.Println("\nВыберите задание для проверки. Для выхода наюерите 0.")
		fmt.Println("1 Задание 1 Уберите из первого примера (Фибоначчи и спиннер) функцию, вычисляющую числа Фибоначчи. Как теперь заставить спиннер вращаться в течение некоторого времени?")
		fmt.Println("2 Задание 2. Дополните пример из раздела «Пакет img» изображением горизонтальных и вертикальных линий.")
		fmt.Println("3 Задание 3. Дополните функцию hello() http-сервера так, чтобы принять и отобразить на странице один GET-параметр, например name.")
		fmt.Println("5 Задание 5. Напишите программу, генерирующую png-файл с рисунком шахматной доски.\n")

		fmt.Scan(&task)
		switch task {
		case 1:

			//timeSpin.TimeSpin()
			//fmt.Println("Получили: ", statistic.SumArg([]float64{1,2,3,4,5,6}))
		case 2:
			line_img.LineImg()
		case 3:
			server.ServerGo()
		case 5:
			chess_board.ChessBoard()
		case 0:
			return
		}
	}
}