package main

import "fmt"

func main() {

x := make([] float64, 5) // slice of type float and with an ArraySize of 5
cc := [] int{1,2,3,3}
 aa := cc
 fmt.Println(aa[3])

x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
 for i,v :=range x {
		fmt.Println(i,v)
	}

	//another way to create Array is t
  //xl := make([] float64,5,10)

  // slice2 := []float64{1,2,3}
   slice1 := []int{1,2,3}
   slice2 := append(slice1,4,5)
   fmt.Println(slice1, slice2)

    slice3 := append(slice2,2,344,566,77,788,)
    fmt.Println(slice3)


}
