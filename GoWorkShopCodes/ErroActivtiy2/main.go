package main

import (
	"errors"
	"fmt"
)
//Activity Solved

var (
	ErrInvalidLastName      = errors.New("Invalid last name Head")
	ErrInvalidRoutingNumber = errors.New("invalid routing number Head")
)

type DirectDeposit struct {
	FirstName string
	LastName  string
	BankName  string

	RoutingNumber int
	AccountNumber int
}

func (di *DirectDeposit) validateRoutingNumber(rout int) int {
	defer func() {
		if err := recover(); err != nil {
			if err == ErrInvalidRoutingNumber {
				//  panic(errors.New("invalid routing number"))
				fmt.Printf("%vrr%d\ninvalid routing number:", err, rout)
			}
			fmt.Printf("invalid routing number entered  is%v :", rout)

		}
		fmt.Printf("the routing number  is  %v  ",rout)


	}()

	if di.RoutingNumber < 100 {
		//di.RoutingNumber = rout
		return 0
	}else {
		di.RoutingNumber = rout
		return di.RoutingNumber

		//fmt.Printf( "%v\n",di.RoutingNumber)
	}
}

func (ls *DirectDeposit) validateLastName(lsn string) string {
	defer func() {

		if err := recover(); err != nil {
			if err == ErrInvalidLastName {
				//  panic(errors.New("invalid routing number"))
				fmt.Printf("%vrr%d\n invalid last name ifErr:", err, lsn)
			}
			fmt.Sprintf("invalid name or field empty %v :", lsn)
		}



	}()

	if ls.LastName == "" {
		//fmt.Printf("%v", ErrInvalidLastName)
	}
	ls.LastName = lsn
	return ls.LastName
}
func Report(d DirectDeposit) {
	if d1 := d.validateRoutingNumber(d.RoutingNumber); d1 < 100 {
	//	fmt.Printf("%v\n", ErrInvalidRoutingNumber)
		 fmt.Println(ErrInvalidRoutingNumber)
		return
	}

	//panic(errors.New("invalid routing number"))

	if bb := d.validateLastName(d.LastName); bb == "" {
		 fmt.Println(ErrInvalidLastName)
	} else {

		fmt.Printf("**************************\n\tFirst Name: %s\n\t Last Name : %s\n\t Bank Name: %s\n\t Routing Number: %v\n\t Account Number:%v\t", d.FirstName, d.LastName, d.BankName, d.RoutingNumber, d.AccountNumber)

	}
}







func main() {

	dn := DirectDeposit{
		FirstName:     "James Brown",
		LastName:      "",
		BankName:      "FirstBank",
		RoutingNumber: 500,
		AccountNumber: 213233,
	}

	Report(dn)
	//name := "James ohn"
	//number := 124455
	//allNames(name, number)

}
//
//func allNames(n string, number int) {
//	nn := 1234455
//
//	defer func() {
//		if r := recover(); r != nil {
//			if r == ErrInvalidLastName {
//				fmt.Printf("the name entered was: %s\n %v", n, r)
//			}
//			if r == ErrInvalidRoutingNumber {
//				fmt.Printf("the ruting number entered: %v\n %v", number, r)
//			}
//		}
//		fmt.Printf("the Last Name was%s:, and the routing number was%v: ", n, number)
//	}()
//
//	if n != "James John" {
//		panic(ErrInvalidLastName)
//	}
//	if number != nn {
//		panic(ErrInvalidRoutingNumber)
//
//	} else {
//		fmt.Printf("Your Last Name is : %s\n and Your routing number is%d :", n, number)
//	}
//
//}
