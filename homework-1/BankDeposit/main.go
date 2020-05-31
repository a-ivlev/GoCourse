package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	fmt.Printf("Введите сумму в рублях на которую вы хотите открыть вклад?\n")
	sumLn := "0"
	fmt.Scan(&sumLn)
	sumInt, err := strconv.Atoi(lenLn)
	if err != nil {
		log.Fatalln(err)
	}
	katetA := float64(lenInt)
}
