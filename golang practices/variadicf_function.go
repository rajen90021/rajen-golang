package main

import "fmt"

func sum(a ...int) {
	fmt.Println(a, " ")

	total := 0
	for _, value := range a {
		total = total + value
	}
	fmt.Println(total)
}
func main() {

	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4, 5, 6, 6}
	sum(nums...)

}
