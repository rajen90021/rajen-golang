package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("seven is odd ")
	} else {
		fmt.Println("7 is odd ")
	}
	fmt.Println("****************")
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
		fmt.Println("****************")
		if num := 9; num < 0 {
			if 9%2 == 0 {
				fmt.Println(num, "is positive  ")
			}
		} else if num < 10 {
			fmt.Println(num, " has one digit")
		} else {
			fmt.Println(num, "has multiple digits")
		}

	}

}
