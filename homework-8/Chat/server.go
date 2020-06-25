package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client chan<- string

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string)
)



func main() {
	go listener()

	userInput := ""
	fmt.Println("Для завершения работы сервера наберите exit")
	fmt.Scanln(&userInput)
	if userInput == "exit" {
		return
	}

}

func listener()  {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	// Главная горутина прослушивает и принимает входящие сетевые подключения от клиентов.
	// Для каждого из них создается новая горутина handleConn.
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

// broadcaster хранит информацию о всех клиентах и прослушивает каналы событий и сообщений,
// используя мультиплексирование с помощью select.
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

// handleConn создает новый канал исходящих сообщений для своего клиента
// и объявляет широковещателю о поступлении этого клиента по каналу entуring.
// Затем она считывает каждую строку текста от клиента,
// отправляет их широковещателю по глобальному каналу входящих сообщений,
// предваряя каждое сообщение указанием отправителя.
// Когда от клиента получена вся информация, handleConn объявляет
// об убытии клиента по каналу leaving и закрывает подключение.
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "Для завершения работы наберите exit"
	ch <- "Ваше имя в чате " + who
	messages <- who + " подключился к чату"
	entering <- ch

	input := bufio.NewScanner(conn)
	for input.Scan() {
		if input.Text() == "exit" {
			leaving <- ch
			messages <- who + " вышел из чата"
			conn.Close()
		}
		messages <- who + ": " + input.Text()
	}/*
	leaving <- ch
	messages <- who + " вышел из чата"
	conn.Close()*/
}

// handleConn создает горутину clientWriter для каждого клиента. Она получает широковещательные сообщения
// по исходящему каналу клиента и записывает их в его сетевое подключение. Цикл завершается,
// когда широковещатель закрывает канал, получив уведомление leaving.
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}