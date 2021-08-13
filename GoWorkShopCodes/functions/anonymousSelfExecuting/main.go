package main

import "fmt"

func main() {
	message := "Greetings"
	func(str string){
		fmt.Println(str)

	}(message)

}
