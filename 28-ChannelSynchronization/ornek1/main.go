package main

import (
	"fmt"
)

func worker(done chan bool, sayiArgs ...int) {
	fmt.Println("working...")
	for _, sayi := range sayiArgs {
		fmt.Println(sayi)
	}
	//time.Sleep(time.Second)
	fmt.Println("false")
	done <- false
	fmt.Println("true")
	done <- true

}

func main() {
	sayi := []int{55, 26, 32, 14}
	done := make(chan bool, 1)
	go worker(done, sayi...)
	<-done
}
