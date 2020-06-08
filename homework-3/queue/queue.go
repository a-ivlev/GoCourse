package queue

import (
	"errors"
)

var s []string

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
