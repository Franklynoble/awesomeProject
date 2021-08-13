package main

import "fmt"

func square(x int) func ()int {

	f := func() int {
	return x*x
	}
	return f
}

func main() {

	v := square(9)
	fmt.Printf("%v\n",v())

}
