package main

import (
	"fmt"
	"strconv"
)

// this is an activity but to be revisited, one more thing to complete this, is to Add the Review Activity functionality (rating)

type Employee struct {
	id        int
	FirstName string
	LastName  string
}
type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}
type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}

type Payer interface {
	Pay() (string, float64)
}

func main() {
	d := Developer{Individual: Employee{FirstName: "Eric",
		LastName: "Davis"},HoursWorkedInYear: 5000, HourlyRate: 50,Review: map[string]interface{}{"Work Quality": 4, "Team Work": 3, "Skill": 1}}
	m := Manager{Individual: Employee{
		id:        0,
		FirstName: "Mr Boss",
		LastName:  "",
	},Salary: 2000.1,CommissionRate: 100.0}

	  // reviewFunc()
	  var dc Developer
	 dc.Review["Work Quality"] = 4
	dc.Review["Work Quality"] = 4
	//reviewFunc("Work Quality",4)
	payDetails(d, dc, m)



}

//func (s Developer) Pay()(string,float64) {
//
//	return
//
//}
func (d Developer) Pay() (string, float64) {
	devYearPay := d.HourlyRate * d.HoursWorkedInYear
	return d.Individual.LastName, devYearPay

}
func (dm Manager) Pay() (string, float64) {
	Salary := dm.Salary + (dm.Salary * dm.CommissionRate)
	m := fmt.Sprintf(dm.Individual.FirstName + " " + dm.Individual.LastName + "Got Paid")
	return m, Salary
}

func payDetails(p ...Payer) {
	for _, v := range p {
		fmt.Println(v.Pay())
	}

}

func reviewFunc(rating string, i interface{}) (b Developer) {
	b.Review[rating] = rating
	n := []int{1, 2, 3, 4, 5}
	n2 := []string{"1", "2", "3", "4", "5"}

	switch v := i.(type) {

	case int:

		if v = n[4]; v == n[4] {
			//c := float64(v)
			//total :=  len(b.Review)

			b.Review[rating] = v
			b.Review[rating] = "Excellent"
			return b
		}
		if v == n[1] {
			b.Review[rating] = v
			b.Review[rating] = "Fair"
			return b

		}
		if v == n[3] {
			//f = sum(b.Review[rating])
			b.Review[rating] = v
			//f = sum(v)
			b.Review[rating] = "Good"

			return b

		}
		if v == n[4] {
			b.Review[rating] = v
			b.Review[rating] = "Excellent"
			return b
		}
		if v == n[1] {
			b.Review[rating] = v
			b.Review[rating] = "Unsatisfactory"
			return b
		}
		//return b
	case string:
		// strconv.Atoi(n2[1])
		if v == "1" {
			strconv.Atoi(n2[1])
			b.Review[rating] = v
			b.Review[rating] = "Unsatisfactory"
		}
		if v == "2" {
			strconv.Atoi(n2[2])
			b.Review[rating] = v
			b.Review[rating] = "fair"
		}
		if v == "3" {
			strconv.Atoi(n2[3])
			b.Review[rating] = v
			b.Review[rating] = "Good"
		}
		if v == "4" {
			strconv.Atoi(n2[4])
			b.Review[rating] = v
			b.Review[rating] = "Excellent"
		}

	default:
		return b
	}
	return b

}

//func sum(i interface{}) int  {
//	total := 0
//
//
//	switch c := i.(type) {
//	case int:
//		for ss := 0; ss < c; ss ++ {
//            total += c
//			//return total
//		}
//return total
//
//
//		}
//
//	}
//}
