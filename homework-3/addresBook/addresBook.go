package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	//Name string
	//Phone int
	task int
	//reply string
	//AddresBook map[string][]int
)

//AddresBook = make(map[string][]int)

/*func WriteContact() {
	//AddresBook := make(map[string][]int)

	AddresBook["Alexey"] = []int{7900047372}
	AddresBook["Ivan"] = []int{7432637482}
	AddresBook["Bob"] = append(AddresBook["Bob"], 79003214546)
	AddresBook["Bob"] = append(AddresBook["Bob"], 79003210046)

//			fmt.Println("Введите Имя контакта")
//			fmt.Scanln(&Name)
//			fmt.Println("Введите номер телефона")
//			fmt.Scan(&Phone)
//			AddresBook[Name] = append(AddresBook[Name], Phone)
			/*for {
			fmt.Printf("Записать ещё один номер к контакту %s ? Y/N ", Name)
			fmt.Scan(&reply)
			if reply == "y" ||  reply == "Y" || reply == "д" || reply == "Д" {
				fmt.Scan(&Phone)
				AddresBook[Name] = append(AddresBook[Name], Phone)
			}
			//if reply == "n" ||  reply == "N" || reply == "д" || reply == "Д" {
				//break
			for name, phone := range AddresBook{
				fmt.Println("Абонент: ", name)
				for i, phone := range phone {
					fmt.Printf("\t %v) %v \n", i+1, phone)
					//return
					}
				}

			}
}*/
	


/*func PrintContact(){
	AddresBook := make(map[string][]int)
	for name, phone := range AddresBook{
		fmt.Println("Абонент: ", name)
		for i, phone := range phone {
			fmt.Printf("\t %v) %v \n", i+1, phone)
		}
	}

}*/

func main() {
	AddresBook := make(map[string][]int)

for{
	fmt.Println("Если хотите записать новый контакт или отредактировать старый нажмите 1\nНажмите 2 чтобы распечатать список контактов\n")
	fmt.Println("Для выхода из программы наберите 0")
	fmt.Scan(&task)
	if task == 1 {
		AddresBook["Alexey"] = []int{7900047372}
		AddresBook["Ivan"] = []int{7432637482}
		AddresBook["Bob"] = append(AddresBook["Bob"], 79003214546)
		AddresBook["Bob"] = append(AddresBook["Bob"], 79003210046)
	}

	if task == 2 {
		for name, phone := range AddresBook {
			fmt.Println("Абонент: ", name)
			for i, phone := range phone {
				fmt.Printf("\t %v) %v \n", i+1, phone)
			}
		}
	}
	if task == 3{
		serialized, err := json.Marshal(AddresBook)
		if err != nil{
			fmt.Println(err)
			return
		}
		fmt.Println(serialized)
		fmt.Println(string(serialized))
		//message := serialized
		message := []byte("Hello, Gophers!")
		error := ioutil.WriteFile("~/AddresBook", message, 0777)
		if error != nil {
			log.Fatal(error)
		}
		fmt.Println("Запись в файл прошла успешно.")
	}

	if task == 0 {
		return
	}
}
		//fmt.Println("Вы не выбрали что делать дальше.")
	}
	//addresBook["Alexey"] = []int{7900047372}
	//addresBook["Ivan"] = []int{7432637482}
	//addresBook["Bob"] = append(addresBook["Bob"], 79003214546)
	//addresBook["Bob"] = append(addresBook["Bob"], 79003210046)
	



