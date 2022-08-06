package main

import "fmt"

func main() {
	nums := []int{55, 28, 21, 52}
	sum := 0
	for _, num := range nums {
		sum = sum + num
	}
	fmt.Println("Sum ", sum)

	for _, num := range nums {
		if num == 28 {
			fmt.Println("Selam Giresun")
		}
	}
	urun := map[string]string{"Dolap": "Arçelik Soğutma", "TV": "LG 45"}
	for key, value := range urun {
		fmt.Printf("%s Kategorisindeki ürün : %s\n ", key, value)
	}

	// rune olarak dönen dataları string e çevirme
	for i, c := range "Mustafa" {
		fmt.Println(i, string(c))
	}
}
