package main

import (
	"errors"
	"fmt"
)

func getdivision(num1, num2 int) (int, error) {
	if num2 == 0 {
		return -1, errors.New("divide by zeroo")
	}
	return num1 / num2, nil
}

func main() {

	result, err := getdivision(20, 0)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("the division of two number is ", result)
	}

}
