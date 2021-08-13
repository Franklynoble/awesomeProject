package main

import "fmt"

func main() {
	//fizzBuzz()
	for i := 1; i <= 15; i++ {
		 gh := fact(i)
		 fmt.Printf("%d\n ",gh)
		//n, s := fizzBuzzReturn(i)
		//fmt.Printf("Result: ,%d %s\n", n, s)
	}
}

func fact(n int) int {
	if n == 1 {
		n *= 1
		return n

	}
	if n == 2 {
		n *= 2
		return n
	}
	if n == 3 {
		n *= 3
		return n
	}
 return  0
}
func fizzBuzz() {
	for i := 1; i <= 30; i++ {
		if i%15 == 0 {
			fmt.Println("Fizz")
		} else if i%3 == 0 {

		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func fizzBuzzReturn(kk int) (int, string) {
	switch {
	case kk%15 == 0:
		return kk, "FizzBuzz"
	case kk%3 == 0:
		return kk, "Fizz"
	case kk%5 == 0:
		return kk, "Buzz"
	}
	return kk, ""

}
