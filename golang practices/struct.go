package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	fmt.Println(person{"bob", 20})
	fmt.Println(person{name: "alice", age: 30})
	dog := struct {
		name   string
		isGood bool
	}{"Rex", true}
	fmt.Println(dog)

}
