package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Printf("Введите сумму в рублях на которую вы хотите открыть вклад:\n")
	userInput := "10000"
	fmt.Scan(&userInput)
	deposit, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Введите процент который будет начисляться ежегодно на ваш вклад:\n")
	fmt.Scan(&userInput)
	procent, err := strconv.ParseFloat(userInput, 64)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Введите колличество лет на которые открывается вклад:\n")
	fmt.Scan(&userInput)
	numbeYear, err := strconv.Atoi(userInput)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Посмотрите как будет расти ваш вклад:\n")
	for kolLet := 1; kolLet <= numbeYear; kolLet++ {
		deposit += deposit * procent / 100.00
		fmt.Printf("%d год размер вашего вклада составит %.2f рублей.\n", kolLet, deposit)
	}
}
