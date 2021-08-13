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


type weekDay int

const (
	SUNDAY weekDay = iota
	MONDAY
	TUESDAY
	WEDNESDAY
	THURSDAY
	FRIDAY
	SATURDAY
)

var TrackHours int
var notLoggedHours int
var hourOverTime = 40


/*
The salary function accepts a function that accepts two integers as arguments and
returns an int. So, any function that matches that signature can be passed as an
argument to the salary function:*/

func salary(x, y int, f func(int, int) int) int {
	/*
		In the body of the salary() function, pay is assigned the value that gets returned
		from the f function. It passes x and y parameters as parameters to the f parameter:
	*/
	pay := f(x, y)
	return pay
}

//Notice that the managerSalary and developerSalary signatures are identical and
//they match the function f for salary. This means that both managerSalary and
//developerSalary can be passed as func(int,int) int:
func managerSalary(basSalary, bonus int) int {
	return basSalary + bonus

}

//devSalary and bossSalary get assigned to the results of the salary function. Since
//developerSalary and managerSalary satisfy the signature of func(int,int) int, they
//each can be passed in as arguments:
//func developerSalary(hourlyRate, hourWorked int) int {
//
//      var hourOverTime int
//	 if hourWorked > overTimeCount {
//	 	all := hourlyRate*2
//		 hourOverTime = all
//	 	  return  hourOverTime
//
//	 }
//	return hourlyRate * hourdWorked
//}
func developerSalary(hourlyRate, hourWorked int) int {
	  //   hourOverTime
	if hourWorked > hourOverTime {
		all := hourlyRate * 2
		hourOverTime = all
		return hourOverTime

	} else {
		return hourlyRate
	}
}
func (d *Developer)HoursWorked()int {
	total := 0
	for _, v := range d.WorkWeek {
		total += v
		}
		return total
}

func nonLoggedHours() func(int) int {

total := 0
return func(i int) int {
	total += i
	return total
}
}
func (d *Developer)PayDay()(int,bool) {
	if d.HoursWorked() > 40 {
		hoursOver := d.HoursWorked() - 40
		overTime := hoursOver * 2 * d.HourlyRate
		regularPay := 40 * d.HourlyRate
		return regularPay + overTime, true
	}
	return d.HoursWorked() * d.HourlyRate,false
}
func (d  *Developer) PayDetails () {

	for i , v := range d.WorkWeek {
		switch i {
		case 0:
			fmt.Println("Sunday hours: ", v)
		case 1:
			fmt.Sprintln("Monday hours", v)

		case 2:
			fmt.Println("Tuesday hours ", v)
		case 3:
			fmt.Println("Wedneday hours: ", v)
		case 4:
			fmt.Println("Thursday hours :", v)
		case 5:
			fmt.Println("Friday hours: ", v)
		case 6:
			fmt.Println("Saturday hours: ", v)
		}
	}
	fmt.
		Printf("\nHours Worked this week: %d\n",d.HoursWorked())
	pay,overTime := d.PayDay()
	fmt.Println("Pay for the week: $ ",pay)
	fmt.Println("is the Overtime pay: ",overTime)
	fmt.Println()
	}
func main() {
	devSalary := salary(50, 280, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)

	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Developer salary: %d\n", devSalary)
	d := Developer{Individual: Employee{Id: 1,FirstName: "Tony",Lastname: "Stark"},HourlyRate: 10}
    x := nonLoggedHours()
    fmt.Println("Tracking hours worked thus far today: ",x(2))
	fmt.Println("Tracking hours worked thus far today: ",x(3))
	fmt.Println("Tracking hours worked thus far today: ",x(5))
	fmt.Println()
    d.LogHours(MONDAY,8)
    d.LogHours(TUESDAY,10)
	d.LogHours(WEDNESDAY,10)
	d.LogHours(THURSDAY,10)
	d.LogHours(FRIDAY,6)
	d.LogHours(SATURDAY,8)
    d.PayDetails()
}


func (d *Developer)LogHours(day weekDay,hours int) {

	d.WorkWeek[day] = hours
}