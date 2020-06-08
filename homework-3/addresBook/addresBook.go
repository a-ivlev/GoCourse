package addresBook

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	Name string
	Phone int
	reply string

)

// AddresBook функция позволяет добавить новые контакты, вывести на экран и сохранить телефонный справочник.
func AddresBook() {
	AddresList := make(map[string][]int)

	AddresList["Alexey"] = []int{7900047372}
	AddresList["Ivan"] = []int{7432637482}
	AddresList["Bob"] = append(AddresList["Bob"], 79003214546)
	AddresList["Bob"] = append(AddresList["Bob"], 79003210046)

	fmt.Println("Телефонный справочник")
	for name, phone := range AddresList {
		fmt.Println("Абонент: ", name)
		for i, phone := range phone {
			fmt.Printf("\t %v) %v \n", i+1, phone)
		}
	}
	for {
		fmt.Printf("Записать ещё контакт? Y/N ")
		fmt.Scan(&reply)
		if reply == "y" || reply == "Y" || reply == "д" || reply == "Д" {
			fmt.Println("Введите Имя контакта")
			fmt.Scanln(&Name)
			fmt.Println("Введите номер телефона")
			fmt.Scan(&Phone)
			AddresList[Name] = append(AddresList[Name], Phone)
		}

		if reply == "n" || reply == "N" || reply == "н" || reply == "Н" {
			serialized, err := json.Marshal(AddresList)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(serialized)
			fmt.Println(string(serialized))
			message := serialized
			error := ioutil.WriteFile("~/AddresBook", message, 0777)
			if error != nil {
				log.Fatal(error)
			}

			fmt.Println("Запись в файл прошла успешно.")
			return
		}
	}
}




