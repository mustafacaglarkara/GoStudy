package main

import (
	"fmt"
)

func Yazdir[T any](Deger T) {
	fmt.Println(Deger)
}

func SayiArttir[age int64 | float32](myAge age) {
	val := int(myAge) + 1
	fmt.Println(val)
}

func main() {
	var isim = "Mustafa Çağlar KARA"
	var yas int64 = 38
	Yazdir(isim)
	Yazdir(yas)
	SayiArttir(yas)
}
