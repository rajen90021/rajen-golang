package main

import "fmt"

func multiplywithchennel(ch chan int) {
	fmt.Println(100 * <-ch)
}

func main() {

	///var firstchan chan int      ////creating a chennel  /// if we use this syntax  the output will be <nil>
	// firstchan := make(chan int) //declaration a chennel as a int type
	// fmt.Printf(" the value of the chennel is  ", firstchan)  //0xc00001e120
	// fmt.Printf("type of the chennel  %T ", firstchan) ///chna type
	ch := make(chan int)
	fmt.Println("hello from main ")
	///ch <- 10/// it will give error
	go multiplywithchennel(ch)
	ch <- 10
	fmt.Println("bye from main ")

	ch := make(chan int)
	/////elem, done := <-ch error du to deadlock condition
	close(ch)
	elem, done := <-ch
	fmt.Println("checking if chennel is close or not ", done, elem) /////false  means chennel is closed ,  0 that means there is no element
}
