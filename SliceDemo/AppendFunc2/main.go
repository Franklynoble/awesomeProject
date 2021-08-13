package main

import (
	"fmt"
)

func main() {
	//var x, y []int
	//for i := 0; i < 10; i++ {
	//	y = appendInt(x, i)
	//	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	//	x = y // append y to the array element x
	//}

	checkSlice()
}

func appendInt(x []int, y int) []int {

	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// there is room to grow, extend the slice.
		z = x[:zlen]
	} else {
		// there is insufficient space, allocate a new array
		// grow by doubling , for amortized linear complexity
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap) // make func.. zlen this is the legnth of the array, zcap, is the capacity of the capacity
		copy(z, x)                  // a built-in function; add x to z
	}
	z[len(x)] = y
	return z
}


func checkSlice() {
	var x [] int
	 x = append(x,1)
	 x = append (x,2,3)
	 x = append(x,4,5,6)
	 x = append(x,x...) // append the slice x to x
	 fmt.Println(x)     //"[123456123456]"

}
