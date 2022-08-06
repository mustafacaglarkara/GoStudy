package main

import "fmt"

type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}

type container struct {
	base
	str string
}
type describer interface {
	describe() string
}

func main() {
	co := container{
		base: base{55},
		str:  "Mustafa Çağlar KARA",
	}
	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)
	fmt.Println("also num:", co.base.num)
	fmt.Println("describe:", co.describe())

	var d describer = co
	fmt.Println("describer:", d.describe())

}
