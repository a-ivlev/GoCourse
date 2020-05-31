package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

var katetA float64 = 0
var katetB float64 = 0
var lenLn string = "0"
var lenInt int = 0
var hypotenuse float64 = 0
var areaTringle float64 = 0
var perimetr float64 = 0

func main() {
	fmt.Printf("Введите длину катета (a) прямоугольного треугольника\n")
	fmt.Scan(&lenLn)
	lenInt, err := strconv.Atoi(lenLn)
	if err != nil {
		log.Fatalln(err)
	}
	katetA = float64(lenInt)
	fmt.Printf("Введите длину катета (b) прямоугольного треугольника\n")
	fmt.Scanln(&lenLn)
	lenInt, err = strconv.Atoi(lenLn)
	if err != nil {
		log.Fatalln(err)
	}
	katetB = float64(lenInt)
	areaTringle = 0.5 * katetA * katetB
	fmt.Printf("Площадь данного треугольника равняется %.2f\n", areaTringle)

	hypotenuse = math.Sqrt(math.Pow(katetA, 2) + math.Pow(katetB, 2))
	fmt.Printf("Гипотенуза данного прямоугольного треугольника равна %.2f\n", hypotenuse)

	perimetr = katetA + katetB + hypotenuse
	fmt.Printf("Периметр данного треугольника составляет %.2f\n", perimetr)
}
