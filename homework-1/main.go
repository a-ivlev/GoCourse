package main

import (
	"fmt"

	"github.com/GoCourse/homework-1/AreaTringle"
	"github.com/GoCourse/homework-1/BankDeposit"
	"github.com/GoCourse/homework-1/CurrencyConv"
)

func main() {

	fmt.Println("1. Программа для конвертации рублей в доллары и в евро.\n Программа запрашивает сумму в рублях и выводит сумму в долларах и в евро.\n")
	CurrencyConv.CurrencyConv()

	fmt.Println("\n2. Программа для расчёта площади, периметра и гипотенузы прямоугольного треугольника.\n Пользователь вводит с клавиатуры длину катетов (a) и (b) прямоугольного треугольника.\n")
	AreaTringle.AreaTringle()

	fmt.Println("\n3. Программа для расчета суммы вклада и процентов по нему через колличество лет заданных пользователем.\n Программа запрашивает сумму вклада в рублях, процент под который размещается вкалад и колличество лет на которые открывается вклад.\n")
	BankDeposit.BankDeposit()
}
