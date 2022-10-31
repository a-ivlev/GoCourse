package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(60 * time.Millisecond)
	time.Sleep(10 * time.Second)
}

func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}