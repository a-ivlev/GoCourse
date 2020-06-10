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
	selec int
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
		_, ok := AddresList[Name]
		if ok {
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
		fmt.Printf("Абонент: %v не найден!\n", Name)
		return
	}
}

// PrintListPhone выводит на экран содержимое телефонной книги.
func PrintListPhone()  {
	fmt.Println("Телефонный справочник")
	for name, phone := range AddresList {
		fmt.Printf("Абонент: %v  %v\n", name, phone)
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

// AddresBook позволяет выбрать действие которое мы хотим совершить с телефонным справочником.
func AddresBook() {
	for {
		fmt.Println("Выберите действие:")
		fmt.Println("1 Загрузить телефонную книгу из json файла.")
		fmt.Println("2 Записать телефон.")
		fmt.Println("3 Удалить телефон.")
		fmt.Println("4 Сохранить телефонную книгу в json файл.")
		fmt.Println("5 Печать телефонной книги.")
		fmt.Println("0 Выход из программы.")

		fmt.Scan(&selec)
		switch selec {
		case 1:
			LoadfromJson()
		case 2:
			SavePhone()
		case 3:
			DelPhone()
		case 4:
			SaveJson()
		case 5:
			PrintListPhone()
		case 0:
			return
		}
	}
}


