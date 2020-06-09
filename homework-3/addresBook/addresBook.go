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

const fileName = "./AddresBook.json"

var AddresList = make(map[string]int)

// SavePhone записывает новый контакт в телефонную книгу.
func SavePhone()  {
	for {
		fmt.Printf("Записать новый контакт? Y/N ")
		fmt.Scan(&reply)
		if reply == "y" || reply == "Y" || reply == "д" || reply == "Д" {
			fmt.Println("Введите Имя контакта")
			fmt.Scanln(&Name)
			fmt.Println("Введите номер телефона")
			fmt.Scan(&Phone)
			AddresList[Name] = Phone
			SaveJson()
		}
		if reply == "n" || reply == "N" || reply == "н" || reply == "Н" {
			PrintListPhone()
			return
		}
	}
}

// DelPhone удаляет введённый контакт из адресной книги.
func DelPhone()  {
	for {
			fmt.Println("Введите Имя контакта который хотите удалить")
			fmt.Scanln(&Name)
			/*for name, phone := range AddresList {
				if Name == name {
					fmt.Println("Абонент: ", name, phone)
				}
				fmt.Printf("Абонент: %v не найден!", Name)
				return
			}*/
			fmt.Printf("Вы действительно хотите удалить %v ? Y/N ", Name)
			fmt.Scan(&reply)
			if reply == "y" || reply == "Y" || reply == "д" || reply == "Д" {
				delete(AddresList, Name)
				PrintListPhone()
				SaveJson()
				return
		}
		if reply == "n" || reply == "N" || reply == "н" || reply == "Н" {
			PrintListPhone()
			return
		}
	}
}

// PrintListPhone выводит на экран содержимое телефонной книги.
func PrintListPhone()  {
	fmt.Println("Телефонный справочник")
	for name, phone := range AddresList {
		fmt.Println("Абонент: ", name, phone)
	}
}

// SaveJson записывает данные телефонной книги в файл.
func SaveJson() {
	serialized, err := json.Marshal(AddresList)
	if err != nil {
		fmt.Println(err)
		return
	}

	error := ioutil.WriteFile(fileName, serialized, 0644)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Println("Запись в файл прошла успешно.")
}

// LoadfromJson Загружает AddresBook из json файла.
func LoadfromJson()  {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))

	if  err := json.Unmarshal(content, &AddresList); err != nil {
		log.Fatal(err)
	}
	PrintListPhone()
}

// AddresBook Добавляет новые контакты, вывести на экран и сохранить телефонный справочник.
func AddresBook() {
	// Заготовленный список контактов.
/*	AddresList["Alexey"] = 7900047372
	AddresList["Ivan"] = 7432637482
	AddresList["Bob"] = 79003214546
	AddresList["Bob"] = 79003210046*/


	LoadfromJson()
	DelPhone()
	//PrintListPhone()
	//SavePhone()
	//SaveJson()
}





