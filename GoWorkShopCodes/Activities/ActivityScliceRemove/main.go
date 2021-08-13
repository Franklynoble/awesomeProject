package main

import (
	"fmt"
)

func main() {
	removeWord()

}
func removeWord() {
	r1 := []string{"Good","Good","Bad","Good","Good"}

      //  r1 =  r1[:2] + r1[3:]

   r1 = append(r1[:2],(r1[3:])...)
   fmt.Println(r1)

}




