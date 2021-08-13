package main

import "fmt"

func main() {
	var runes []rune

	for _,r := range "Hello 你好" {
		runes = append(runes, r)
	}
	//fmt.Printf("%q\n",runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' 'B' 'F']"
 strl := []rune("Ok, now proceed") // here  the  more convinient way to convert to rune
 fmt.Printf("%q",strl)



}
