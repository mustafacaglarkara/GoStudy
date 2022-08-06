package main

import "fmt"

func main() {

	var a [5]int
	a[2] = 55
	fmt.Println(a)

	var b = [5]int{1, 5, 66, 8, 7}
	fmt.Println(b)

	c := [5]int{15, 66, 5, 8, 9}
	fmt.Println(c)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
