package main

import "fmt"

func main() {
	userName:= "Sir_King_Uber"
	for i := 0; i < len(userName); i ++ {
		fmt.Println(userName[i])
	}
	fmt.Println("After Converting")
	//convert()
	rangeOverString()
	usingLength()

}
//convert the byte back to string

func convert(){
	//You can see that when using len directly on a string, you get the wrong answer.
	//Checking the length of data input using len in this way would end up with invalid data.
	//For example, if we needed the input to be exactly 8 characters long and somebody
	//entered a multi-byte character, using len directly on that input would allow them to
	//enter less than 8 characters.
	//To safely work with interindividual characters of a multi-byte string, you first must
	//convert the strings slice of byte types to a slice of rune types.
	userName := "Sire_King_Uber"
	runes := []rune(userName)
	for i := 0; i < len(runes); i++ {
		//var runes [] rune

		fmt.Print(string(runes[i]), " ",i)
	}
}

//5. Declare the string with a multi-byte string value:
//We'll then loop over the string using range to give us each character, one at a time. We'll
//then print out the byte index.html and the character to the console:

func rangeOverString() {
	logLevel := "デバッグ"

	for index,values := range logLevel {
		fmt.Println(index,string(values))
	}

}
func usingLength(){
	username := "Sir_King_Über"
	//length of a string
	fmt.Println("Bytes: ",len(username))
	fmt.Println("Runes: ",len([]rune(username)))
	//limit to 10 characters
	fmt.Println("string length: ",string(username[:10]))
	fmt.Println("string rune Length: ",string([]rune(username)[:10]))
}