package main

import (
	"fmt"
	"time"
)

var dayBorn = time.Sunday
func main() {
	//caseWeek()
	bornDay2()

	switch dayBorn {
	case time.Monday:
		fmt.Println("Moday's child is fair of face")

	case time.Tuesday:
		fmt.Print("Tuesday's child is full of grace")

	case time.Wednesday:
		fmt.Println("Wednesday's child is full of Woe")
	case time.Thursday:
		fmt.Println("Thursday's child has far to go")

	case time.Friday:
		fmt.Println("Friday's child is loving and giving")
	case time.Saturday:
		fmt.Println("Saturday's child works hard for a living")
	case time.Sunday:
		fmt.Println("Sunday's child  is bonny and blithe")
	default:
		fmt.Println("Error Born day incorrect")
	}
}
// check whether the born day was weekend or not,
//checking for Multiple options
func caseWeek(){

	switch dayBorn{
	case   time.Monday,time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Born on a weekDay")
	case time.Saturday,time.Sunday:
		fmt.Println("Born on Weekend")
	default:
		fmt.Println("Error, day born not valid")

	}
}

//Expressionless switch Statements
//
func bornDay2() {
	//4. Our switch expression is using the initial statement to define our variable. The
	//expression is left empty as we'll not be using it:
	switch bornDay1 := time.Sunday; {

	case bornDay1 == time.Sunday || bornDay1 == time.Saturday:
		fmt.Println("Born on the Weekend")
	default:
		fmt.Println("Born some other day")
	
	}
}