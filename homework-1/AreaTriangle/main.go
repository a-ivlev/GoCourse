package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

var katetA int = 0
var katetB int = 0
var lenght string = "0"
var hypotenuse float64 = 0
var areaTringle float64 = 0
var perimetr float64 = 0

func main() {
	fmt.Printf("Введите длину катета (a) прямоугольного треугольника\n")
	fmt.Scanf("%.f", &lenght)
	katetA, err := strconv.Atoi(lenght)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Введите длину катета (b) прямоугольного треугольника\n")
	fmt.Scanf("%.f", &lenght)
	katetB, err := strconv.Atoi(lenght)
	if err != nil {
		log.Fatalln(err)
	}
	areaTringle = 0.5 * float64(katetA) * float64(katetB)
	fmt.Printf("Площадь данного треугольника равняется %.f\n", areaTringle)

	hypotenuse = math.Sqrt(math.Pow(float64(katetA), 2) + math.Pow(float64(katetB), 2))
	fmt.Printf("Гипотенуза данного прямоугольного треугольника равна %.f\n", hypotenuse)

	perimetr = float64(katetA) + float64(katetB) + hypotenuse
	fmt.Printf("Периметр данного треугольника составляет %.f\n", perimetr)
}
