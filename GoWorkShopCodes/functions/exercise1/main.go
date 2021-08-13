package main

import "fmt"

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
func developerSalary(hourlyRate, hourdWorked int) int {
	return hourlyRate * hourdWorked
}

func main() {
	devSalary := salary(50, 280, developerSalary)
	bossSalary := salary(150000, 25000, managerSalary)

	fmt.Printf("Boss salary: %d\n", bossSalary)
	fmt.Printf("Developer salary: %d\n", devSalary)

}
