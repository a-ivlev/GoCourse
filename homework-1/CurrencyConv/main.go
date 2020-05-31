package main

import (
	"fmt"
	"log"
	"strconv"
)

const dollar = 70.75
const euro = 78.55

func main() {
	fmt.Printf("Курсы валют на сегодня $-%.2f руб за доллар и евро €-%.2f руб за евро.\n", dollar, euro)
	fmt.Println("Какую сумму в рублях вы хотите обменять?")
	summa := "10000"
	fmt.Scanln(&summa)
	rub, err := strconv.ParseFloat(summa, 64)
	if err != nil {
		log.Fatalln(err)
	}
	convdollar := rub / dollar
	conveuro := rub / euro
	fmt.Printf("После обмена вы получите %.2f$ или евро %.2f€.\n", convdollar, conveuro)
}
