package main

import "fmt"

func main() {
	monthsT()

	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 2)
	copy(slice2, slice1) //dest slice2, src slice . copy slice1 into slice2
	fmt.Println(slice1, slice2)
	fmt.Println("The Capacity is ", cap(slice2), len(slice2)) // this prints the length and capacity of the slice array element



}

func monthsT() {

	months := []string{"", "january", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

	Q2 := months[4:7]
	fmt.Println(Q2)
	summer := months[6:9]
	fmt.Println(summer)

	//fmt.Println(summer[:20])  // panic: out of range
	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)

	// this sample, inefficient test for common element

	for _,s := range summer {
			for _,q := range Q2 {
				if s == q {
					fmt.Printf("%s appears in both\n", s)
				}
			}
	}

}
