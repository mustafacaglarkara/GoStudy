package main

import "fmt"

func main() {

	messages := make(chan string, 3)

	messages <- "Bu birinci mesaj"
	messages <- "Bu ikinci mesaj"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
