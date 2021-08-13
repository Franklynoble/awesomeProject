package main

import (
	"fmt"
	"time"
	)







var count int

func main() {
	//getPointer()
	add5Value(count)
	add5Point(&count)

//
//var count1 *int
//count2 := new(int)
//countTemp := 5
//count3 := &countTemp
//t := &time.Time{}

	//fmt.Printf("count1: %#v\n",count1)
	//fmt.Printf("count2: %#v\n",count2)
	//fmt.Printf("count3: %#v\n",count3)
	//fmt.Printf("time : %v\n",t)

}

/*
  Getting value from a pointer
  differencing a variable
*/

func getPointer() {
	var count1 *int
	count2 := new(int)
	countTemp := 5
	count3 := &countTemp
	t := &time.Time{}

	if count1 != nil {
		fmt.Printf("count1 %#v\n", *count1)
	}
	if count2 != nil {
		fmt.Printf("count2 : %#v\n", *count2)
	}
	if count3 != nil {
		fmt.Printf("count3 :%#v\n", *count3)
	}
	if t != nil {
		fmt.Printf("time : %#v\n", *t)
		fmt.Printf("time : %#v\n",t.String())
	}

}
//function Design With Pointer
func  add5Value(count int) {
	count += 5
	fmt.Println("add5Value: ",count)

}
func add5Point(count *int) {
	 *count += 5
	 fmt.Println("add5Point :", *count)
}