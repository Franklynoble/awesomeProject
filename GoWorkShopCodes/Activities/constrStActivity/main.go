package main

import (
	"fmt"
	"strconv"
)

var COUNT = 100

func main() {

	for i := 1; i < COUNT; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(strconv.Itoa(i), "iFIZZBUZZi")
		} else if i%3 == 0 {
			fmt.Println("FIZZ")
		} else if i%5 == 0 {
			fmt.Println("BUZZ")

		} else {

			fmt.Println(i)
		}

	}
}
