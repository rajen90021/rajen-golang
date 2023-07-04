package main

import (
	"fmt"
	"time"
)

func dosomemagic() {

	for i := 0; i < 10; i++ {
		time.Sleep(2 * time.Second) ////if time sleep is not there then control will normally as a function
		///to execute goroutine we have to give a specfic time to control to execute a goruntine
		fmt.Println("helloo")
	}
}
func main() {

	// go dosomemagic() // we call  goroutine
	// dosomemagic()

	fmt.Println("hello rajen ")
	go func() {
		fmt.Println("hiiiiii rajen ") ///same as go for also anoyomous function
	}()
	time.Sleep(2 * time.Second)
}
