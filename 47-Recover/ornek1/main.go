package main

import "fmt"

func HataPanic() {
	panic("Problem Oluştu")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()

	HataPanic()
	fmt.Println("Hata Panic Sonrası")
}
