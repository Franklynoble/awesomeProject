package main

import (
	"fmt"
	"unicode"
)

func main() {
	appsn()
	if passWordChecker("This15A") {
		fmt.Println("Good Password")
	} else {
		fmt.Println("password Bad")
	}

}




func appsn() {
	fmt.Println(1 > 20)
}
//program to Measure password Complexity
func passWordChecker(pw string) bool {
	pwR := []rune(pw)
	if len(pwR) < 8 {
		return false
	}
	//Define some bool variables. We'll check these at the end:
	hasUpper := false
	hasLower := false
	hasNumber := false
	hasSymbol := false

	// loop Over the Multi-byte characters one at a time
	for _, v := range pwR {
		// using the unicode package, check whether this character is uppercase.
		//this function returns a bool that we can use directly in the if statement

		if unicode.IsUpper(v) {
			hasUpper = true
		}
		if unicode.IsLower(v) {
			hasLower = true
		}
		if unicode.IsNumber(v) {
			hasNumber = true
		}
		//For symbols, we'll also accept punctuation. Use the or operator, which works with
		//Booleans, to result in true if either of these functions returns true:
		if unicode.IsPunct(v) || unicode.IsSymbol(v) {
			hasSymbol = true
		}

	}
	return hasUpper && hasLower && hasNumber && hasSymbol

}
