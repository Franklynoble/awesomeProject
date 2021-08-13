package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	content , err := ioutil.ReadFile("Test")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("file contents: ")
	//For this code snippet, I have a test.txt file that is located in the same location as
	//my executable.
	//Since this is a slice of bytes, it must be converted to a string format for ease of
	//readability. The results of the print statements are as follows:
	fmt.Println(string(content))

}

