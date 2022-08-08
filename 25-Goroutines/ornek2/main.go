package main

import (
	"fmt"
	"time"
)

func say() {
	var sayilar = []int{5, 8, 44, 68, 84}
	for sayi := range sayilar {
		fmt.Print(sayi)
	}
}

func oku() {
	var isim = "Mustafa Çağlar KARA"
	for _, karakter := range isim {
		fmt.Print(string(karakter))
	}
}

func main() {
	go oku()
	go say()

	go func(msg string) {
		fmt.Println(msg)
	}("going")

	time.Sleep(time.Second * 5)
	fmt.Println("\nProgram Sonlandı")
}
