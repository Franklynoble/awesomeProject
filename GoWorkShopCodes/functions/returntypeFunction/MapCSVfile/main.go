package main

import (
	"fmt"
	"strings"
)

func main() {
hdr := []string{"empid","employee","address","hours worked","hourly","rate","Manager"}

result := csvHdrCol(hdr)
fmt.Println("result:")
fmt.Println(result)
fmt.Println()
hdr2 := []string{"employee","empid","hours worked","address","manager","hourly rate"}
   result2 := csvHdrCol(hdr2)
   fmt.Println("Result2: ")
   fmt.Println(result2)
   fmt.Println()



}

func csvHdrCol(hdr []string) map[int]string {
	csvIdxCol := make(map[int]string)

	//We range over the header to process each string that is in the slice:
	for i, v := range hdr {
		v = strings.TrimSpace(v)
		//
		//In our switch statement, we lower all the casing for exact matches. As you may
		//recall, Go is a case-sensitive language. We need to ensure the casing is the same for
		//matching purposes. When our code finds the header, it sets the index value for the
		//header in the map:
		switch strings.ToLower(v) {
		case "employee":
			csvIdxCol[i] = v
		case "hours worked":
			csvIdxCol[i] = v
		case "hourly rate":
			csvIdxCol[i] = v

		}
	}
	return  csvIdxCol
}
