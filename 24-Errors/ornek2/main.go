package main

import (
	"fmt"
)

// Kendi Error Yapımız
type argError struct {
	arg  any
	prob string
}

// Struct' ın Error func'ı
func (e *argError) Error() string {
	return fmt.Sprintf("%s - %s", e.arg, e.prob)
}

func topla(sayi int) (int, error) {
	if sayi == 42 {
		//return -1, errors.New("Sayı 42 olamaz")
		return -1, &argError{sayi, "Sayı 42 olamaz"}
	}
	return sayi + 1, nil
}

func main() {
	sayi, err := topla(42)
	//err dolu ise aşağıdaki kodlar çalışır
	if ae, ok := err.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
	fmt.Println(sayi)

}
