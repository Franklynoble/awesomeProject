package main

import "fmt"

func main() {
var nelement[10]int
    n1 := fillArray(nelement)
    fmt.Printf("%d\n",n1)
}


func fillArray(n [10]int)[10]int{
	for i := 0; i < len(n); i ++ {

		n[i] = 1+ i
	}
	return n

}