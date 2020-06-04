package even_number

import (
	"fmt"

	"github.com/GoCourse/mypackage"
)

// EvenNumber() эта функция определяет введено четное число или нечётное.
func EvenNumber() {

	number := mypackage.UserInputInt("Введите число которое хотите проверить чётное оно или нет?")

	if number%2 == 0 {
		fmt.Printf("Число %d чётное\n", number)
		return
	}
	fmt.Printf("Число %d нечётное\n", number)
}
