package main

import (
	"fmt"
	"time"
)

type Point struct{x,y int}

type Employee struct{
	ID   int
	Name string
	Address string
	DOB  time.Time
	Position string
	Salary int
	ManagerID int
}




func main() {
fmt.Print(Scale(Point{1,2},5))
}


func Scale(p Point, factor int) Point {
	 return Point{p.x * factor, p.y * factor}
}

/* for efficiency larger struct are usually passed to or returned from functions inderectly
using a pointer
 */

