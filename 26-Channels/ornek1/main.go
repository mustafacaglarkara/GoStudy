package main

import "fmt"

func main() {
	messages := make(chan string)

	go func() {
		messages <- "Merhaba ben kanaldan geldim"
	}()

	fmt.Println(<-messages)
}
