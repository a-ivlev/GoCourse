package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var (
	SumQuest1Var1, SumQuest1Var2, SumQuest1Var3 float64
	SumApartArea float64
	ApartCount int
	Quest1 = "Организация стоянки автомобилей во дворе дома"
	Quest1Var1 = "Закрытый двор без машин"
	Quest1Var2 = "Распределение мест"
	Quest1Var3 = "Свободный въезд"

)

type Golosovan struct {
	Data        string `csv,"line[0]"`
	Person      string `csv,"line[1]"`
	ApartNumber string `csv,"line[2]"`
	ApartArea   string `csv,"line[3]"`
	Question    Question
}

type Question struct {
	Question1 string `csv,"line[4]"`
}

// Данная программа предназначена для подсчёта голосов.
//Данные уже подготовлены т.е. получены из Гугл формы и записаны в csv файл.
// Программа обработала ответы и вывела результат на экран.
func main() {
	csvFile, err := os.Open("Golosovanie.csv")
	if err != nil {
		log.Fatalln(err)
	}
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var Vote []Golosovan
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		Vote = append(Vote, Golosovan{
			Data:        line[0],
			Person:      line[1],
			ApartNumber: line[2],
			ApartArea:   line[3],
			Question: Question{
				Question1: line[4],
			},
		})

		ApartCount ++
		apartArea, _ := strconv.ParseFloat(line[3], 64)
		SumApartArea = SumApartArea + apartArea
		//Забанил проверку на ошибки, так как выходила ошибка.
		//Т.К. первая строка csv файла у меня поясняет поля.
		//Как запустить чтение со второй строки не разобрался.
		//if err != nil {
		//	log.Fatalln(err)
		//}
		vote1 := line[4]

		if vote1 == Quest1Var1 {
			SumQuest1Var1 += apartArea
		}
		if vote1 == Quest1Var2 {
			SumQuest1Var2 += apartArea
		}
		if vote1 == Quest1Var3 {
			SumQuest1Var3 += apartArea
		}
	}
	fmt.Printf("%v\n", Vote)
	fmt.Printf("Проголосовало %v квартир. Общей площадью %.2f кв.м.\n", ApartCount, SumApartArea)
	fmt.Printf("По первому вопросу %s проголосовали:\n\n", Quest1)
	fmt.Printf("За первый вариант %s проголосовали: %.2f кв.м\n", Quest1Var1, SumQuest1Var1)
	fmt.Printf("За второй вариант %s проголосовали: %.2f кв.м\n", Quest1Var2, SumQuest1Var2)
	fmt.Printf("За третий вариант %s проголосовали: %.2f кв.м\n", Quest1Var3, SumQuest1Var3)
}
