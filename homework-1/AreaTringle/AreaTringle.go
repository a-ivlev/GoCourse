package AreaTringle

import (
	"fmt"
	"math"

	"github.com/GoCourse/mypackage"
)

// Функция считает площадь, периметр и гипотенузу прямоугольного треугольника.

func AreaTringle() {

	katetA := mypackage.UserInputFloat("Введите длину катета (a) прямоугольного треугольника")
	katetB := mypackage.UserInputFloat("Введите длину катета (b) прямоугольного треугольника")

	areaTringle := 0.5 * katetA * katetB
	fmt.Printf("Площадь данного треугольника равняется %.2f\n", areaTringle)

	hypotenuse := math.Sqrt(math.Pow(katetA, 2) + math.Pow(katetB, 2))
	fmt.Printf("Гипотенуза данного прямоугольного треугольника равна %.2f\n", hypotenuse)

	perimetr := katetA + katetB + hypotenuse
	fmt.Printf("Периметр данного треугольника составляет %.2f\n", perimetr)
}
