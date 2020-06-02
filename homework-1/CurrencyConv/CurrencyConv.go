package CurrencyConv

import (
	"fmt"

	"github.com/GoCourse/mypackage"
)

const dollar = 70.75
const euro = 78.55

// Данная функция переводит сумму в рублях в доллары и в евро.

func CurrencyConv() {

	fmt.Printf("Курсы валют на сегодня $-%.2f руб за доллар и евро €-%.2f руб за евро.\n", dollar, euro)

	rub := mypackage.UserInputFloat("Какую сумму в рублях вы хотите обменять?")

	convdollar := rub / dollar
	conveuro := rub / euro
	fmt.Printf("После обмена вы получите %.2f$ или евро %.2f€.\n", convdollar, conveuro)
}
