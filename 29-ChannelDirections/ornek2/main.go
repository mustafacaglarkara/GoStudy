package main

import "fmt"

//Gönderici
func sender(channel chan<- string, message string) {
	channel <- message
}

//Alıcı
func comm(receiver <-chan string, sender chan<- string) {
	message := <-receiver
	sender <- message
}

func main() {
	scott := make(chan string, 1)
	tiger := make(chan string, 1)

	sender(scott, "My name is Bond.James Bond")
	comm(scott, tiger)

	fmt.Println(<-tiger)
}
