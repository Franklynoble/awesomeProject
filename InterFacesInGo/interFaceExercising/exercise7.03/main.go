package main

import "fmt"

type record struct {
	Key       string
	valueType string
	data      interface{}
}

type person struct {
	lastName  string
	age       int
	isMarried bool
}
type animal struct {
	name     string
	category string
}

func main() {
	//function, we will initialize our map. The map is initialized to a
	//string for the key and an empty interface for the value. We then assign a to an
	//animal struct literal and p to a person struct literal. Then, we start adding various
	//key-value pairs to the map:

	m := make(map[string]interface{})
	a := animal{name: "oreo", category: "cat"}
	p := person{lastName: "Doe", isMarried: true, age: 19}

	m["person"] = p
	m["animal"] = a
	m["age"] = 54
	m["isMarried"] = true
	m["lastName"] = "Smith"

	//we initialize a slice of record. We iterate over the map and add records to rs:
	rs := []record{}
	for k, v := range m {
		r := newRecord(k, v)
		rs = append(rs, r)
	}
	//print out the record field values. We range over the slice of records and print
	//each record value:
	for _, v := range rs {
		fmt.Println("Key", v.Key)
		fmt.Println("Data: ", v.data)
		fmt.Println("Type: ", v.valueType)
		fmt.Println()
	}

}

func newRecord(key string, i interface{}) record {
	//we initialize record{} and assign it to the r variable.
	//We then assign r.key to the key input parameter

	//The switch statement assigns the type of i to the v variable. The v variable type gets
	//evaluated against a series of case statements. If a type evaluates to true for one of
	//the case statements, then the valueType record gets assigned to that type, along
	//with the value of v to r.data, and then returns the record type:

	r := record{}
	r.Key = key
	switch v := i.(type) { // compare the type, if v is equal to the compared type, int, then assign int to the value na d ....

	case int:
		r.valueType = "int"
		r.data = v
		return r
	case bool:
		r.valueType = "bool"
		r.data = v
		return r
	case string:
		r.valueType = "string"
		r.data = v
		return r
	case person:
		r.valueType = "person"
		r.data = v
		return r
	default:
		r.valueType = "unkown"
		r.data = v
		return r

	}
}
