package main

import "fmt"

type arac struct {
	marka, model string
	yas          int
}

func (a arac) Ilerle() string {
	return "İlerliyorum"
}

type IArac interface {
	Ilerle() string
}

func Ilerle(a IArac) {
	fmt.Println(a.Ilerle())
}

func main() {
	var megane = arac{
		marka: "Renault",
		model: "Megane 2",
		yas:   2009,
	}
	Ilerle(megane)
}
