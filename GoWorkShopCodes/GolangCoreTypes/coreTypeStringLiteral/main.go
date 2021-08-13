package main

import "fmt"

func main() {
	comment1 := `This is the Best thing Ever!`
    comment2 := `This is the Best\nthing ever!`
    comment3 := "This is the Best\nthing ever!"
    fmt.Println(comment1)
    fmt.Println(comment2)
    fmt.Println(comment3)
    fmt.Print(" ")
	commenta1 := `In "Windows" the user directory is "C:\Users\"`
	commenta2 := "In \"Windows\" the user directory is \"C:\\Users\\\""
	fmt.Println(commenta1)
	fmt.Println(commenta2)


}
