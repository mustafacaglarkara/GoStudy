package main

import "fmt"

func Yazdir[T any](Deger T) {
	fmt.Println(Deger)
}

func main() {
	var isim = "Mustafa Çağlar KARA"
	Yazdir(isim)
}
