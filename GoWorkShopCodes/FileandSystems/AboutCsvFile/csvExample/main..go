package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)

func main() {

	in := `firstName, lastName,age
Celina,Jones, 18
Ceilyn, Henderson,13
Cayden, Smith, 24
`

	//The following creates a reader type and returns it:
	r := csv.NewReader(strings.NewReader(in))
	//The NewReader method takes an argument of io.Reader and returns a type of Reader that
	//is used to read the CSV data:
	 header := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(record) }
	if !header{
		for idx , value := range record {

			switch idx {
			case 0:
				fmt.Println("First Name: ",value)
			case 1:
				fmt.Println("Last Name: ", value)

			case 2:
				fmt.Println("Age: ",value)

			}

			}

		}
		header = false
	}

}
