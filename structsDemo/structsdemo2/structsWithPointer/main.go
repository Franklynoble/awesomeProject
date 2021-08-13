package main

import (
	"fmt"
	"time"
)
type Point struct {x,y int }

type Employee struct {

	ID int
	Name string
	Address string
	DoB  time.Time
	Position string
	Salary int
	ManagerID int
}
var dilbert Employee

func main() {

	 pp := &Point{1,2} // this but &Point{1, 2} can be used
	// direc tly within an expression, such as a function cal l.


 	fmt.Print(pp)
	//  pp := new(Point)
	  *pp = Point{1,2}



 dilbert.Salary -= 5000 // demoted for writing very few lines codes

 // or take the address and access it through a pointer
 position := &dilbert.Position
 *position = "Senior "+  *position   // promoted, for outsourcing to Elbonia
 // the dot noytation also work for pointer to a struct

 var employeeOfTheMonth *Employee = &dilbert
  employeeOfTheMonth.Position += "(proactive team player)"

}

//For efficiency, larger str uct typ es are usually passed to or retur ned from functions indirec tly
//using a pointer,
func AwardAnnualRaise (e  *Employee)  {
	e.Salary = e.Salary * 105 / 100
	fmt.Println(e.Salary)
}
func Bonus (e *Employee, percent int ) int {
	return e.Salary * percent / 100
}
