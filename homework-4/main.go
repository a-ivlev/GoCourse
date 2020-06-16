package main

import (
	"fmt"
	"github.com/GoCourse/homework-4/calculator"
	"github.com/GoCourse/homework-4/getspeed"
	"github.com/GoCourse/homework-4/sortList"
	"sort"
)

var task int

func Calculat() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "help" {
			calculator.Help()
			continue
		}

		if input == "exit" {
			break
		}

		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}


func main() {
	for {
		fmt.Println("\nВыберите задание для проверки. Для выхода наюерите 0.")
		fmt.Println("1 Задание 1 Написать свой интерфейс и создать несколько структур, удовлетворяющих ему.")
		fmt.Println("2 Задание 2. Создать тип, описывающий контакт в телефонной книге. Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него интерфейс Sort{}.")
		fmt.Println("3 Задание 3. Дописать функцию, которая будет выводить справку к калькулятору. Она должна вызываться при вводе слова help вместо выражения.\n")
		fmt.Scan(&task)
		switch task {
		case 1:
			Mazda := getspeed.Auto{Marka: "Мазда", Way: 60, Time: 1}
			fmt.Printf("Автомобиль %s двигается со скоростью %v km/h.\n", Mazda.Marka, getspeed.GetSpeed(Mazda))
			man := getspeed.Human{Name: "Иван", Age: 25, Phone: 89102255555, Way: 5, Time: 1}
			speed := getspeed.Status(man).Speed()
			fmt.Printf("Человек %s двигается со скоростью %v km/h.\n", man.Name, speed)
		case 2:

			//sorting by phone
			var AdressBook = sortList.ContactList{
				{ "Алексей",  79003214546},
				{"Иван",  7432637482},
				{"Михаил", 79175552233},
				{"Андрей", 79272777123},
				{"Родион", 7903335566},
				{"Татьяна", 79275784525},
			}
			sort.Sort(AdressBook)
			fmt.Println(AdressBook) // [{Алексей 79003214546} {Андрей 79272777123} {Иван 7432637482} {Михаил 79175552233} {Родион 7903335566} {Татьяна 79275784525}]
		case 3:
			Calculat()
		case 0:
			return
		}
	}
}
