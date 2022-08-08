package main

import (
	"fmt"
	"sort"
)

func main() {
	var isimler = []string{"Mustafa", "Çağlar", "Öykü", "Ülkü", "Beren"}

	sort.Strings(isimler)
	fmt.Printf("Sıralanmış olarak data %q", isimler)

	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
