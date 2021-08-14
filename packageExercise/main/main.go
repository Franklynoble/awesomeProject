package main

import "github.com/Franlky01/Exercises/packageExercise/shape"

func main() {

	t := shape.Rectangle{
		Length: 8,
		Width:  6,
	}
	r := shape.Rectangle{
		Length: 21.1,
		Width:  22.1,
	}
	s := shape.Square{
		Side: 12,
	}
	shape.PrintShapeDetails(t,r,s)

}
