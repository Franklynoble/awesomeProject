package main

import (
	"fmt"
	"sort"
)

var ages map[string]int

func main() {


	useOk()

	//fmt.Print("prints from range sort")
	//rangeAndSort()
	//creating map using the make func
	//ages1 := make(map[string]int)
	//ages1["jax"] = 0
	//ages1["Ampil"] = 2
	//ages1["bob"] = ages1["bob"] + 1 // assign 1 to key "bob"
	//ages1["afgan"] ++               // assign 1 to key
	//for kk, vv := range ages1 {
	//	fmt.Println(kk, vv)
	//}

}

func mapdemo1() {
	ages = map[string]int{"alice": 31, "charlie": 34}

	for _, ageInt := range ages {
		fmt.Print(ageInt)

	}
}
func rangeAndSort() {
	ages = map[string]int{"alice": 31, "charlie": 34, "goldberg": 5}
	var names []string

	for name := range ages {
		names = append(names, name)
	}

	sort.Strings(names)
	for _, name := range names { // the key has been thrown, we only have the value
		fmt.Printf("%s\t%d\n", name, ages[name]) // ages[name] would access the key
	}
	//you must allocate a map  before you can store into it
	//	ages["carol"] = 21 // panic: assignment to entry in nil map
}
func useOk() {
	/*  For many pur pos es that’s fine, but sometimes you need to know
	whet her the element was really there or not. For example, if the element typ e is numer ic, you
	mig ht have to distinguish bet ween a nonexistent element and an element that happens to have
	the value zero, using a test like this: */
	ages = map[string]int{"alice": 31, "charlie": 34, "goldberg": 5}
	var names []string

	for name := range ages {
		names = append(names, name)
	}
	if age,ok := ages["bob"]; !ok{ /* if bob is not a key on this map,*/
		fmt.Println("print ok func")
		fmt.Println(`no key with name "bob"`,age)
	}else {
		sort.Strings(names)
		for _, name := range names { // the key has been thrown, we only have the value
			fmt.Printf("%s\t%d\n", name, ages[name]) // ages[name] would access the key
		}
	}
	//you must allocate a map  before you can store into it
	//	ages["carol"] = 21 // panic: assignment to entry in nil map

}
