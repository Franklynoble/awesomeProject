package main

import "fmt"

type Shape interface {
	Area() float64
	Name() string
}


//Next, we will create triangle, rectangle, and square struct types. These types will
//each satisfy the Shape{} interface. triangle, rectangle, and square have appropriate
//fields that are needed to calculate the area of the shape:
type triangle struct {
	base   float64
	height float64
}

type rectangle struct {
	length float64
	width  float64
}
type square struct {
	side float64
}

func main() {
	//Inside the main() function, set the fields for triangle, rectangle, and square. Pass
	//all three to the printShapeDetail() function. All three can be passed because they
	//each satisfy the Shape interface:
	t := triangle{base:15,height: 20.1}
	r := rectangle{length: 20,width: 10}
	s := square{side: 10}
	printShapedDetails(t,r,s)

}

//We create the Area() and Name() methods for the triangle struct type. The area of
//a triangle is base * height\2. The Name() method returns the name of the shape:
func (t triangle) Area() float64 {

	return (t.base * t.height) / 2
}
func (t triangle) Name() string {
	return "triangle"
}
//We create the Area() and Name() methods for the rectangle struct type. The area of
//a rectangle is length * width. The Name() method returns the name of the shape:
func (r rectangle) Area() float64 {
	return r.length * r.width
}

func(r rectangle)Name()string {
	return "rectangle"
}
//We create the Area() and Name() methods for the square struct type. The area of a
//square is side * side. The Name() method returns the name of the shape:
func(s square)Area()float64 {
	return s.side * s.side
}

func(s square) Name()string {
	return "square"
}
//Now, each of our shapes (triangle, rectangle, and square) satisfies the Shape
//interface because they each have an Area() and Name() method with the
//appropriate signatures:

//
//We will now create a function that accepts the Shape interface as a variadic
//parameter. The function will iterate over the Shape type and will execute each of its
//Name() and Area() methods:

func printShapedDetails(shapes ...Shape) {
	for _,item := range shapes {
		fmt.Printf( "The Area of %s is: %.2f\n",item.Name(),item.Area())
	}
}