package main

import "fmt"

func main() {
	fmt.Println("************************")
	func(x int) {
		fmt.Println("the value of x is ", x)
	}(100)
	fmt.Println("************************")

	value := func(x int) int {
		fmt.Println("the value of x is ", x)
		return -1
	}(200)
	fmt.Println("the value of ", value)
	fmt.Println("************************")

	func_variable := func() {

		fmt.Println("te function was executed ")
	}
	fmt.Printf("%T\n", func_variable)
	func_variable()
	fmt.Println("************************")
	complex_func_variable := func(x int, value string) string {
		fmt.Println("the value of x is ", x)
		fmt.Println("the value of value is ", value)
		return "executed "
	}
	output := complex_func_variable(80, "anand")
	fmt.Println("output is ", output)
	fmt.Println("************************")
//  function closures 
  func_var :=  Clouser_sample_func(100)
	func Clouser_sample_func (value int) (int ,string) {
		x := 100
		fmt.Println(x)
		func (value int )(int)  {
			fmt.Println("the value of the passed variable ", value)
			fmt.Println("the value of the x variable ",x)
			x=x+1
			return x
		}(200)
		return x,"value"
	}
	



}
