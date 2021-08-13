package main

import (
	"flag"
	"fmt"
)

func main() {

	i := flag.Int("age",-1,"your age")

	n := flag.String("name","","are you married")

     b := flag.Bool("married",false,"are You married")

     flag.Parse()
     fmt.Print("Name: ",*n)
     fmt.Println("Age: ",*i)
     fmt.Println("Married: ",*b)

}
