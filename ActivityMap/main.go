package main

import (
	"fmt"
)

func main() {
	findkey(5)


}
func findkey(k1 int)  {
   // key := ""
	n := map[int]string{305: "Sue", 204: "Bob", 631: "Jake", 73: "Tracy"}
	 if k, v :=  n[k1]; v {
			fmt.Println("Hi",k)
		}else {
		fmt.Println("errors!")
	 }
}
