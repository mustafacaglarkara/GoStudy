package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	const isim = "Mustafa Çağlar KARA"

	fmt.Println("Len :", len(isim))

	for i := 0; i < len(isim); i++ {
		fmt.Printf("%x ", isim[i])
	}
	fmt.Println()
	fmt.Println("Rune count :", utf8.RuneCountInString(isim))

	//Burdan itibaren bakılacak
	for idx, runeValue := range isim {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(isim); i += w {
		runeValue, width := utf8.DecodeRuneInString(isim[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}
func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
