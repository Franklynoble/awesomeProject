package main

import (
	json2 "encoding/json"
	"fmt"
	"os"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	Middleinitial string   `json:"minitial"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled,omitempty"`
	isMarried     bool     `json:"_"`
	Courses       []course `json:"classes,omitempty" `
}
type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {

	allprint()
	fmt.Println("***********")
	s := newStudent(1, "Williams", "s", "Felicia", false, false)

	//Next, marshal s to JSON. We want the indenting of the JSON to be four spaces for
	//each field for ease of readability:
	student1, err := json2.MarshalIndent(s, "", " ")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(student1))
}
func allprint() {
	//Create another student using the newStudent() function:
	s2 := newStudent(2, "WashingTon", "", "Bill", true, true)
	//We will now add various courses to s2:
	c := course{
		Name: "World Lit", Number: 101, Hours: 3}
	//adds c courses to the Courses variable Object in s2

	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Biology", Number: 201, Hours: 4}
	s2.Courses = append(s2.Courses, c)
	c = course{Name: "Intro to Go", Number: 101, Hours: 4}
	s2.Courses = append(s2.Courses, c)

	student2, err := json2.MarshalIndent(s2, " ", "")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(student2))

}

func newStudent(studentID int, lastName, middleInitial, firstName string, isEnrolled bool, isMarried bool) student {
	s := student{
		StudentId:     studentID,
		LastName:      lastName,
		Middleinitial: middleInitial,
		FirstName:     firstName,
		IsEnrolled:    isEnrolled,
		isMarried:     isMarried,
	}
	return s
}
