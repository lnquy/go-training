package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string)

	go func() {
		//messages <- "XXX"
		time.Sleep(time.Duration(5) * time.Second)
		//messages <- "ping"
	}()


	msg := <-messages
	fmt.Println(msg)
	msg = <-messages
	fmt.Println(msg)
}
