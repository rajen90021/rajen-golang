package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 34
	fmt.Println("***********************")
	fmt.Println("map are ", m)
	fmt.Println("***********************")

	v1 := m["k1"]
	fmt.Println("***********************")

	fmt.Println("v1 is ", v1)
	fmt.Println("***********************")

	fmt.Println("len is ", len(m))
	fmt.Println("***********************")

	delete(m, "k2")
	fmt.Println("map are now ", m)
	fmt.Println("***********************")

	_, prs := m["k2"]
	fmt.Println("prs ", prs)
	fmt.Println("***********************")

	n := map[string]int{"foot": 1, "footi": 2}
	fmt.Println("map", n)
}
