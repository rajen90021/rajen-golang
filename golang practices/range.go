package main

import "fmt"

func main() {

	//	nums := []int{2, 4, 5, 6, 7, 8, 8, 95, 5, 6}
	// sum := 0
	// for index, value := range nums {
	// 	sum = sum + value
	// 	fmt.Println(index)
	// }
	// fmt.Println("sum is ", sum)

	// for i, num := range nums {
	// 	if num == 95 {
	// 		fmt.Println("index is", i)
	// 	}
	// }

	kvs := map[string]string{"a ": "apple", "b": "banana"}
	for key, value := range kvs {
		fmt.Println("%S -> %S\n", key, value)
	}

	for b, d := range kvs {
		fmt.Println("key", b)
		fmt.Println("value ", d)

	}
}
