package main

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	str string
}

func (b base) describe() string {
	return fmt.Sprintf("base with num=%v", b.num)
}
func main() {

	co := container{base: base{num: 1}, str: "some name"}

	fmt.Printf("co={num:%v, str: %v}\n", co.num, co.str)
	fmt.Println("also num ", co.base.num)

	fmt.Println("describe ", co.describe())

	type describe interface {
		describe() string
	}
	var s describe = co
	fmt.Println(s.describe())
}
