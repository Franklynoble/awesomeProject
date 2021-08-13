package main

import "fmt"

type name string
type location struct {
	x int
	y int
}
//create s size struct with two fields, that is width,and height
type size struct {
	width int
	height int
}
//Create a struct named dot.
//This embeds each of the preceding structs in it:
type dot struct {
	name
	location
	size
}
//create a function that  returns a slice of dots:
func getDots()[]dot {
	// our First Dot Uses the var notation. this will result in all the fields having zero
	var dot1 dot
	//with dot two we are also initializing with zero values
	dot2 := dot{}
	// to set the name, we use the type's name as if it were a field
	dot2.name = "A"
	// for size and location, we will use the promoted fields. for name,the result is the
	//the same but for location and size, you need to put more into this
	dot2.x = 5
	dot2.y = 6
	dot2.width = 10
	dot2.height = 20

	//when initializing embedded types, you can'not use promotion ,for name, the result is the same but for location, and size, you need to put More Work into this
	dot3 := dot{
		name:"B",
		location:location{
			x: 13,
			y: 27,
		},
		size:size{
			width: 5,
			height: 7,
		},
	}
	dot4 := dot{}
	dot4.name = "C"
	dot4.location.x = 101
	dot4.location.y = 209
	dot4.size.width = 87
	dot4.size.height = 43
	//return all the dots in a sclice and then close the function
	return []dot{dot1,dot2,dot3,dot4}
}
func main(){

dots := getDots()
for i := 0; i < len(dots); i ++ {
	
		fmt.Printf("dot%v: %#v\n",i+1,dots[i])
}
}
