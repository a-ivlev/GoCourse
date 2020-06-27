package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

func main() {
	go listenerConn()
	userInput := ""
	fmt.Println("Для завершения работы сервера наберите exit")
	fmt.Scanln(&userInput)
	if userInput == "exit" {
		return
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func listenerConn()  {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}