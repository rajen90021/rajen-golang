package main

import (
	"fmt"
	"math"
)

const s string = " constant value"
const PI = 3.14

func main() {

	fmt.Println(s)
	//PI=3.45  cannot reassign the value
	const n = 500000000
	fmt.Println(PI)
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
