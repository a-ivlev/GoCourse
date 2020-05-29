package main

import (
	"fmt"
	"math"
)

var katetA float64 = 0
var katetB float64 = 0
var hypotenuse float64 = 0
var areaTringle float64 = 0
var perimetr float64 = 0

func main() {
	fmt.Printf("Введите длину катета (a) прямоугольного треугольника\n")
	fmt.Scanf("%.f", &katetA)
	fmt.Printf("Введите длину катета (b) прямоугольного треугольника\n")
	fmt.Scanf("%.f", &katetB)
	//katetB, err := strconv.Atoi(lenght)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	areaTringle = 0.5 * katetA * katetB
	fmt.Printf("Площадь данного треугольника равняется %.f\n", areaTringle)

	hypotenuse = math.Sqrt(math.Pow(katetA, 2) + math.Pow(katetB, 2))
	fmt.Printf("Гипотенуза данного прямоугольного треугольника равна %.f\n", hypotenuse)

	perimetr = katetA + katetB + hypotenuse
	fmt.Printf("Периметр данного треугольника составляет %.f\n", perimetr)
}
