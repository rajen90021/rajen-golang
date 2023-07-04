package main

import (
	"fmt"
)

func main() {

	// i := 5
	// fmt.Println("write ", i, "as")
	// switch i {
	// case 1:
	// 	fmt.Println("four")
	// case 5:
	// 	fmt.Println("five ")
	// case 6:
	// 	fmt.Println("six")
	// }
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("its is the weekend")
	// default:
	// 	fmt.Println("its a weekday")
	// }
	// t := time.Now()
	// fmt.Println(t)
	// switch {
	// case t.Hour() > 12:
	// 	fmt.Println("its before noon ")
	// default:
	// 	fmt.Println("its is a after noon")
	// }
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
