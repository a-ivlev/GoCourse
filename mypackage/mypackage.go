package mypackage

import (
	"fmt"
	"log"
	"strconv"
)

var (
	msg      string
	useInput string
	//lenLn string
)

// Функция чтения чисел с плавающей точкой введёных пользователем с клавиатуры. Нужно допилить работу с ошибками.
// Данная функция получает на вход строку которую нужно вывести перед тем как она считает число которое введёт пользователь.
func UserInputFloat(msg string) (useInputFloat float64) {

	fmt.Println(msg)

	fmt.Scan(&useInput)
	useInputFloat, err := strconv.ParseFloat(useInput, 64)
	if err != nil {
		//return errors.New("Вы ввели не число, введите число!!!")
		log.Fatalln(err)
	}
	return useInputFloat
}

// Функция чтения целых чисел введёных пользователем с клавиатуры. Нужно допилить работу с ошибками.
// Данная функция получает на вход строку которую нужно вывести перед тем как она считает число которое введёт пользователь.
func UserInputInt(msg string) (useInputUint int) {

	fmt.Println(msg)
	fmt.Scan(&useInput)
	useInputInt, err := strconv.Atoi(useInput)
	if err != nil {
		//return errors.New("Вы ввели не число, введите число!!!")
		log.Fatalln(err)
	}
	return useInputInt
}
