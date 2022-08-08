package main

import "fmt"

// string gonmessage ye gonderiliyor.
// Gönderilen msg
// alici gonmessage
func sender(gonMessage chan<- string, msg string) {
	gonMessage <- msg
}

// alicimsg e gonmessage gonderiliyor
// gonderici gonmessage alici alicimsg
func receiver(gonMessage <-chan string, aliciMsg chan<- string) {
	message := <-gonMessage
	aliciMsg <- message

}

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	sender(ch1, "Mustafa Çağlar KARA")
	receiver(ch1, ch2)
	fmt.Println(<-ch2)

}
