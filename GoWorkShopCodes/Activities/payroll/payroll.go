package payroll

import (
	"errors"
	"fmt"
)

type Payer interface {
	Pay() (string, float64)
}
type Developer struct {
	Individual        Employee
	HourlyRate        float64
	HoursWorkedInYear float64
	Review            map[string]interface{}
}
type Employee struct {
	Id        int
	FirstName string
	LastName  string
}
type Manager struct {
	Individual     Employee
	Salary         float64
	CommissionRate float64
}
func PayDetails(p Payer) {
	fullName, yearPay := p.Pay()
	fmt.Printf("%s got paid %.2f for the year\n", fullName, yearPay)
}
func (d Developer) Pay() (string, float64) {
	fullName := d.FullName()
	return fullName, d.HourlyRate * d.HoursWorkedInYear
}
func (d Developer) FullName() string {
	fullName := d.Individual.FirstName + " " + d.Individual.LastName
	return fullName
}
func (d Developer) ReviewRating() error {
	total := 0
	for _, v := range d.Review {
		rating, err := OverallReview(v)
		if err != nil {
			return err
		}
		total += rating
	}
	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n",d.FullName(),averageRating)
	return nil
}
func (m Manager) Pay() (string, float64) {
	fullName := m.Individual.FirstName + " " + m.Individual.LastName
	return fullName, m.Salary + (m.Salary * m.CommissionRate)
}
func OverallReview(i interface{}) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := ConvertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("unknown type")
	}
}
func ConvertReviewToInt(str string) (int, error) {
	switch str {
	case "Excellent":
		return 5, nil
	case "Good":
		return 4, nil
	case "Fair":
		return 3, nil
	case "Poor":
		return 2, nil
	case "Unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + str)
	}
}