package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name, age: 38}
	return &p
}

func main() {
	var insan = newPerson("Mustafa Çağlar KARA")
	fmt.Println(insan)

	fmt.Println(person{"Ülkü KARA", 39})
	fmt.Println(&person{name: "Öykü Beren KARA", age: 9})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
	fmt.Println(s.age)
}
