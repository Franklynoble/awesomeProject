package main

import "fmt"

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}


/*
Go lets us declare a field with a typ e but no name; such fields are cal le d anonymous fields. The
type of the field must be a named typ e or a pointer to a named typ e. Below, Circle and Wheel
have one anonymous field each. We say that a Point is embedded within Circle, and a
Circle is embedde d within Wheel.
*/

var w Wheel
func main() {
	// w = Wheel{Circle{Point{8, 8}, 5},  20}
	w = Wheel{
		Circle: Circle{
			Point:  Point{X: 8, Y: 8},
			Radius: 5}, Spokes: 20, // Note: Trailing comma necessary here(and at Radius)
	}
	fmt.Printf("%#v\v",w)
	//OutPut:
	//// Wheel{Circle:Circle{Point:Point{X:8, Y:8}, Radius:5}, Spokes:20}
  w.X = 42
  fmt.Printf("%#v\n",w)
//  w.Radius = 50
	// Output:
	// Wheel{Circle:Circle{Point:Point{X:42, Y:8}, Radius:5}, Spokes:20}

//	Notice how the # adverb causes Printf’s %v verb to display values in a form simi lar to Go syntax.
	//For str uct values, this form includes the name of each field.
}
