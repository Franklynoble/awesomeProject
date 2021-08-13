package main

import (
	"fmt"
)

type Employee struct {
	Id                  int
	FirstName, Lastname string
}
//var Individual Employee
type Developer struct {
	Individual Employee
	HourlyRate int
	WorkWeek   [7]int
}


var developers Developer

var startTime, endTime int
var totalHours int

type Weekday int
const (
	SUNDAY  Weekday = iota  // starts at zero
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

func main() {
	d := Developer{Individual: Employee{Id: 1,FirstName: "Tony",Lastname: "Stark"},HourlyRate:10 }
    d.LogHours(MONDAY,8)
	d.LogHours(TUESDAY,10)
	fmt.Println("Hours worked on Monday: ",d.WorkWeek[MONDAY])
	fmt.Println("Hours worked on Tuesday: ",d.WorkWeek[TUESDAY])
	fmt.Printf("Hours worked this week: %d",d.HoursWorked())


}

func (d *Developer)LogHours(day Weekday,hours int) {

	d.WorkWeek[day] = hours
}

func (d *Developer)HoursWorked()int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
	}
	return total

}