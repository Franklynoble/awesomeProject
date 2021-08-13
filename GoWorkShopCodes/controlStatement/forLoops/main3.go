package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for {

		r := rand.Intn(8)
 // if r is divisible by 3 skip the rest of the loop and using continue
		if r%3 == 0 {
			fmt.Println("Skip")
			continue
			//If the random number is divisible by 2, then print "Stop" and stop the loop using
		} else if r % 2 == 0{
			fmt.Println("Stop")
			break
		}
		fmt.Println(r)
	}
}
