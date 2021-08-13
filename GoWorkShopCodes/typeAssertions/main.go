package main

import (
	"errors"
	"fmt"
)

//Create a function that accepts an interface{}
//and returns a string and an error:
func doubler(v interface{})(string, error ) {

	//we'll check to see if our argument is an int, and if it is, we'll multiply it by 2
	//and return it:
	if i, ok := v.(int);ok {
		return fmt.Sprint(i*2), nil
	}
	//we'll check if it's a string and if it is, w
	//we'll concatenate it to itself and return it:

	if s,ok := v.(string); ok {
		return s + s,nil
	}
	//if we do not have any Matches, return an Error. then close the Function
	return "" ,errors.New("unsupported type passed")
}


func main() {
	res,_ := doubler(5)
	fmt.Println("5: ",res)
	res,_ = doubler("Yum")
	fmt.Println("Yum ",res)
	_, err := doubler(true)
	fmt.Println("True:",err)

}
