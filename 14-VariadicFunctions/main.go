package main

import "fmt"

func topla(sayilar ...int) int {
	nums := sayilar
	var toplam int = 0
	fmt.Println("Aktarılan sayılar : ", nums)
	for _, sayi := range nums {
		toplam = toplam + sayi
	}
	return toplam
}

func main() {
	var sayilar = []int{55, 28, 21, 52}
	fmt.Println(topla(sayilar...))
	fmt.Println(topla(1, 2, 5, 1))

}
