package bankDeposit

import (
	"fmt"

	"github.com/GoCourse/mypackage"
)

func BankDeposit() {

	deposit := mypackage.UserInputFloat("Введите сумму в рублях на которую вы хотите открыть вклад:\n")
	procent := mypackage.UserInputFloat("Введите процент который будет начисляться ежегодно на ваш вклад:\n")
	numbeYear := mypackage.UserInputInt("Введите колличество лет на которые открывается вклад:\n")

	fmt.Printf("\nПосмотрите как будет расти ваш вклад:\n")
	for kolLet := 1; kolLet <= numbeYear; kolLet++ {
		deposit += deposit * procent / 100.00
		fmt.Printf("%d год размер вашего вклада составит %.2f рублей.\n", kolLet, deposit)
	}
}
