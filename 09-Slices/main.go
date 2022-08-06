package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "Mustafa"
	s[1] = "Çağlar"
	s[2] = "KARA"

	s = append(s, "Ülkü")
	s = append(s, "Öykü")
	s = append(s, "Beren")

	fmt.Println(s)
	fmt.Println(s[0])

	fmt.Println(len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	fmt.Println(s[:2])
	fmt.Println(s[2:6])

	fmt.Println(removeSlice(s, 2))
	fmt.Println("Slice Sonrası :", s)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}

func removeSlice(data []string, index int) []string {
	data = append(data[:index], data[index+1:]...)
	return data
}
