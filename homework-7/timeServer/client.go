package main

import (
	"io"
	"log"
	"net"
	"os"
)

func exitProg() ch string {
	ch := make(chan string)
	go func{
	userInput := ""
	fmt.Println("Для завершения работы наберите exit")
	fmt.Scan(&userInput)
		if userInput == "exit" {
			ch <- userInput
		}
	}()
return ch
}


func main() {

	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	if userInput == "exit" {
		defer conn.Close()
		io.Copy(os.Stdout, conn) // игнорируем ошибки
	}

}
