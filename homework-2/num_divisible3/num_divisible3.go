package num_divisible3

import (
	"fmt"

	"github.com/GoCourse/mypackage"
)

// NumDivisible3 функция проверяет делится число на 3 или нет.
func NumDivisible3() {

	number := mypackage.UserInputInt("Введите число которое хотите проверить делится оно на 3 или нет?")

	if number%3 == 0 {
		fmt.Printf("Число %d делиться на 3.\n", number)
		return
	}
	fmt.Printf("Число %d на 3 не делиться.\n", number)
}
