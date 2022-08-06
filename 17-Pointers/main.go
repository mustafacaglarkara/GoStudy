package main

import "fmt"

func sayisalDeger(ival int) {
	ival = 0
}

func sayisalPointer(ivan *int) {
	*ivan = 0
}

func main() {
	i := 1
	fmt.Println(i)

	sayisalDeger(i)
	fmt.Println("sayisalDeger", i)

	sayisalPointer(&i)
	fmt.Println("sayisalPointer", i)
	fmt.Println("pointer", &i)
}
