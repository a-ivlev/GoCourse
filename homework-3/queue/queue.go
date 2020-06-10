package queue

import (
	"errors"
	"fmt"
)

var s []string

func Queue()  {
	fmt.Println("Очередь состоит из следующих фраз:")
	fmt.Println("Этот текст")
	fmt.Println("Будет находиться в очереди")
	fmt.Println("До первого обращения к pop")
	Push("Этот текст", "Будет находиться в очереди", "До первого обращения к pop")
	fmt.Println("Теперь достанем все элементы из очереди\n")
	n := 1
	for i := 0; i < n; i++ {
		element, err := Pop()
		fmt.Println(element)
		if err == nil {
			n++
		}
	}
}

// Push добавляет новый строковый элемент в конец очереди.
func Push(elems ...string) {
	s = append(s, elems...)
}

// Pop берёт первый элемент очереди.
func Pop() (string, error) {
	if len(s) < 1 {
		return "", errors.New("очередь закончилась")
	}
	elem := s[0]
	s = s[1:]

	return elem, nil
}
