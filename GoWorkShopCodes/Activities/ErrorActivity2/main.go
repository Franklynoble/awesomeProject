package main

import (
	"errors"
	"fmt"
)


/*
Activity Solved 2
 */
type DirectDeposit struct {
	FirstName string
	LastName  string
	BankName  string

	RoutingNumber int
	AccountNumber int
}

func (di *DirectDeposit) validateRoutingNumber(rout int )  error {
	if di.RoutingNumber < 100 {
		// di.RoutingNumber = rout
		return ErrInvalidRoutingNumber
	}
	di.RoutingNumber = rout
	//fmt.Printf( "%v\n",di.RoutingNumber)
	return  nil
}
func (ls *DirectDeposit) validateLastName(lsn string) error {

	if ls.LastName == "" {
		return ErrInvalidLastName
	}
	ls.LastName = lsn

	return nil
}

func Report(d DirectDeposit)  {
	if err := d.validateRoutingNumber(d.RoutingNumber); err != nil {
		 fmt.Printf("%v\n",ErrInvalidRoutingNumber)


	}

	if err := d.validateLastName(d.LastName); err != nil {
		fmt.Printf("%v\n ",ErrInvalidLastName)

	}else {
		fmt.Printf("**************************\n\tFirst Name: %s\n\t Last Name : %s\n\t Bank Name: %s\n\t Routing Number: %v\n\t Account Number:%v\t",d.FirstName,d.LastName,d.BankName,d.RoutingNumber,d.AccountNumber)

	}
}




var (
	ErrInvalidLastName      = errors.New("Invalid last name")
	ErrInvalidRoutingNumber = errors.New("invalid routing number")
)


func main() {


	    dn := DirectDeposit{
		    FirstName:     "James Brown",
		    LastName:      "",
		    BankName:      "FirstBank",
		    RoutingNumber: 60,
		    AccountNumber: 213233,
	    }


	    Report(dn)
	//name := "James ohn"
	//number := 124455
	//allNames(name, number)

}
func allNames(n string, number int) {
	nn := 1234455

	defer func() {
		if r := recover(); r != nil {
			if r == ErrInvalidLastName {
				fmt.Printf("the name entered was: %s\n %v", n, r)
			}
			if r == ErrInvalidRoutingNumber {
				fmt.Printf("the ruting number entered: %v\n %v", number, r)
			}
		}
		fmt.Printf("the Last Name was%s:, and the routing number was%v: ", n, number)
	}()

	if n != "James John" {
		panic(ErrInvalidLastName)
	}
	if number != nn {
		panic(ErrInvalidRoutingNumber)

	} else {
		fmt.Printf("Your Last Name is : %s\n and Your routing number is%d :", n, number)
	}

}
