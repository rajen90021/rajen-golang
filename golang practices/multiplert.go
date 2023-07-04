package main

import "fmt"

func vals() (int, int, string) {
	return 3, 7, "ada"
}

func main() {

	a, b, c := vals()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	a, c, _ := vals() //error
	fmt.Println(c)
}
