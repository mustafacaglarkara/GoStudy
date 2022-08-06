package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]
	fmt.Println("v1: ", v1)

	fmt.Println("len:", len(m))

	delete(m, "k2")
	fmt.Println("map:", m)

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

	sehir := make(map[string]int)
	sehir["Samsun"] = 55
	sehir["Giresun"] = 28
	sehir["Ordu"] = 52

	var durum = kontrol(sehir, "Diyarbakır")
	fmt.Println("Durum :", durum)
}

func kontrol(sehir map[string]int, ad string) bool {
	_, ok := sehir[ad]
	return ok
}
