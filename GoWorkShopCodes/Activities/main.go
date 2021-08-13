package main

import (
	"fmt"
	 pr "github.com/Franlky01/Exercises/GoWorkShopCodes/Activities/payroll"
	"os"
)

var s  pr.Developer

//var d payroll.Manager
var employeeReview = make(map[string]interface{})
var m = pr.Manager{Individual: pr.Employee{Id: 2, FirstName: "Mr.", LastName: "Boss"}, Salary: 150000, CommissionRate: .07}
var d = pr.Developer{Individual: pr.Employee{Id: 1, FirstName: "Eric", LastName: "Davis"}, HourlyRate: 35, HoursWorkedInYear: 2400,Review: employeeReview}

func init() {
	fmt.Println("Welcome to the Employee Pay And Performance Review\n"+ "++++++++++++++++++++++++++++++++++")
}
func init() {


	employeeReview["WorkQuality"] = 5
	employeeReview["TeamWork"] = 2
	employeeReview["Communication"] = "Poor"
	employeeReview["Problem-solving"] = 4
	employeeReview["Dependability"] = "Unsatisfactory"

}


func main() {


	err := d.ReviewRating()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	pr.PayDetails(d)
	pr.PayDetails(m)


}
