package main

import "fmt"

func main() {

	i := 1
	for i <= 5 {
		fmt.Println(i)
		//i++
		i = i + 1
	}
	fmt.Println("**************")
	for j := 1; j <= 10; j++ {
		fmt.Println(j)
	}
	fmt.Println("**************")

	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("******* printing odd number *******")
	for n := 0; n <= 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}

}
