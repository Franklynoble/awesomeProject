package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//jsonData represents the JSON that we are given but do not know the structure of:
	jsonData := []byte(`{"checkNum":123,"amount":200,"category":["gift","clothing"]}`)

	//Even though we do not know the JSON structure, we can unmarshal it into an interface.
	//The jsonData gets unmarshaled into v, the empty interface, which will be a map.
	//The map keys are the strings and the values are empty interfaces. The result of printing
	//out v is as follows:
	var v interface{}
	 err := json.Unmarshal(jsonData, &v) //pointer to empty interface

	 if err != nil {
	 	fmt.Println(err)
	 }
	fmt.Println(v)
}
