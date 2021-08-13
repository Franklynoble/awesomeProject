package main

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidLastName      = errors.New("Invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)

func main() {
	name := "James ohn"
	number := 124455
	allNames(name, number)

}
func allNames(n string, number int) {
	nn := 1234455

	defer func() {
		if r := recover(); r != nil {
			if r == ErrInvalidLastName {
				fmt.Printf("the name entered was: %s\n %v", n, r)
			}
			if r == ErrInvalidRoutingNumber {
				fmt.Printf("the ruting number entered: %v\n %v", number, r)
			}
		}
		fmt.Printf("the Last Name was%s:, and the routing number was%v: ", n, number)
	}()

	if n != "James John" {
		panic(ErrInvalidLastName)
	}
	if number != nn {
		panic(ErrInvalidRoutingNumber)

	} else {
		fmt.Printf("Your Last Name is : %s\n and Your routing number is%d :", n, number)
	}

}
