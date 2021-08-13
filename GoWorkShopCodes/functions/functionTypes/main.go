package main

import "fmt"

func main() {
	calculator(add,5,6)

}

func calculator(f func(int,int) int, i,j int ) {
	fmt.Println(f(i,j))
}

func add(i,j int ) int {
	return i + j
}
func subtract(i,j int )int {
	return i - j
}
