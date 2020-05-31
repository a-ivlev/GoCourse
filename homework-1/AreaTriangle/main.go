package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

func main() {
	fmt.Printf("Введите длину катета (a) прямоугольного треугольника\n")
	lenLn := "0"
	fmt.Scan(&lenLn)
	katetA, err := strconv.ParseFloat(lenLn, 64)
	if err != nil {
		log.Fatalln(err)
	}
	//katetA := float64(lenInt)

	fmt.Printf("Введите длину катета (b) прямоугольного треугольника\n")
	fmt.Scanln(&lenLn)
	katetB, err := strconv.ParseFloat(lenLn, 64)
	if err != nil {
		log.Fatalln(err)
	}
	//katetB := float64(lenInt)

	areaTringle := 0.5 * katetA * katetB
	fmt.Printf("Площадь данного треугольника равняется %.2f\n", areaTringle)

	hypotenuse := math.Sqrt(math.Pow(katetA, 2) + math.Pow(katetB, 2))
	fmt.Printf("Гипотенуза данного прямоугольного треугольника равна %.2f\n", hypotenuse)

	perimetr := katetA + katetB + hypotenuse
	fmt.Printf("Периметр данного треугольника составляет %.2f\n", perimetr)
}
