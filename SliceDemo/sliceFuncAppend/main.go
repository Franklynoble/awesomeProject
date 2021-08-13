package main

import "fmt"

func main(){
	// appending a slice to  a slice
	slice1 := []int{1,2,3,4}
	slice2 := append(slice1,11,22,33,44,55,66)
	fmt.Println(slice2)

}
