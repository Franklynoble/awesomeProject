	package main

import (
	"fmt"
	"strings"
)

func main() {
	hdr := []string{"empid","employee","address","hours worked","hourly rate","manager"}
    csvHdrCol(hdr)
	hdr2 := []string{"employee","empid","hours worked","address","manager","hourly rate"}
   csvHdrCol(hdr2)

}

// the function that we are going
//to create will be taking a slice of column headers from a
//CSV file. It will print out a map of an index.html value of the headers we are interested in


func csvHdrCol (header []string) {
	csvHeadersToColumnIndex := make(map[int]string)

	//we assign a variable to a key-value pair of int and string. key(int) will be the
	//index.html of our header(string) column. The index.html will map to a column header.
	//
	for i,v := range header {
		//
		//For each string, remove any trailing spaces in front of and after the string. In
		//general, we should always make the assumption that our data may have some
		//erroneous characters:
		/*
			In our switch statement, we lower all the casing for exact matches. As you may
			recall, Go is a case-sensitive language. We need to ensure that the casing is the
			same for matching purposes. When our code finds the header, it sets the index.html
			value for the header in the map:
		*/
		v = strings.TrimSpace(v)
		switch strings.ToLower(v) {
		case "employee":
			csvHeadersToColumnIndex[i] = v
		case "Hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hours worked":
			csvHeadersToColumnIndex[i] = v
		case "hourly rate":
			csvHeadersToColumnIndex[i] = v
		}
	}
	/*
	Typically, we would not print out the results. We should return the
	csvHeadersToColumnIndex, but since we have not gone over how to return a value,
	we will print it for now:
	 */
       fmt.Println(csvHeadersToColumnIndex)
		}

