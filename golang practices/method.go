package main

import "fmt"

type rectangle struct {
	width  int
	height int
}

func (n rectangle) area() int {
	return n.height * n.width
}
func (n rectangle) perim() int {
	return 2*n.height + 2*n.width
}

func main() {

	r := rectangle{width: 10, height: 5}
	fmt.Println("area ", r.area())
	fmt.Println("perimeter is ", r.perim())
	rp := &r
	fmt.Println("area ", rp.area())
	fmt.Println("perimeter is ", rp.perim())
}
