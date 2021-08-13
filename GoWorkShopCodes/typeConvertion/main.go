package main

import (
	"fmt"
	"math"
)
func convert()string {
	var i8 int8 = math.MaxInt8
	i := 128
	f64 := 3.14
	// here we will convert from smaller int type into a larger int type. this is always safe operation
	m := fmt.Sprintf("int8 = %v  > int64 = %v\n",i8,int64(i8))

	// Now, we will convert from an int that's 1 above int8's maximum size. this would cause an overFlow to int8's minimum size

	m+= fmt.Sprintf("int =%v > in8 =%v\n",i,int64(i8))
	//here we will convert a float into an int. all the decimal data is lost but the whole number is kept as is

//Next, we'll convert out int8 into a float64. This doesn't cause an overflow and the
	//data is unchanged
	//m += fmt.Sprintf("int8 =%v > float32 = %v\n", i8,(f64))
	m += fmt.Sprintf("float64 = %v > int = %v\n", f64, int(f64))

	//here we will convert a float into an int. All the decimal data is lost but the whole number is kept as is
	return m
}

func main() {
	fmt.Print(convert())

}
