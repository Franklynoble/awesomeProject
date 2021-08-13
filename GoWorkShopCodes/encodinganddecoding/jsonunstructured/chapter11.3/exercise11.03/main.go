//In this exercise, we are going to analyze data from a college administration office and
//see whether we can replace the current college course grade submission application.
//The problem is that the old system's JSON data is not well documented. The data types
//in the JSON are not known, nor is the structure. In some instances, the JSON structure
//is different. We need to write a program that can analyze an unknown JSON structure
//and, for each field in the structure, print the data type and the JSON key-value pair.
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

//Create a main() function and then assign jsonData to a []byte that will represent
//the JSON from the college grade submission program:
func main() {
	jsonData := []byte(`
{"id":2,
"lname":"Bill",
"IsEnrolled":true,
  "grades":[100,76,93,50],
 "class":
  {
    "coursename":"World Lit",
     "coursenum":101,
      "coursehours":3
}
}
`)
if !json.Valid(jsonData) {
	fmt.Printf("JSON is not valid : %s",jsonData)
	os.Exit(1)
}
var v interface{}

err := json.Unmarshal(jsonData,&v)  //add the unmarshalled files to the v variable(empty interface)
if err != nil {
	fmt.Println(err)
	os.Exit(1)
}
  
//Perform type switching on each value in the map. Have a case statement for string,
	//float64, bool, []interface, and default to capture the unknown type of a value.
	//Each of the case statements should print the data type, the key, and the value. Our
	//switch type assertion flow should work as shown in the following diagram:


data1 := v.(map[string]interface{})

for k, v := range data1 {
	switch value := v.(type) {
	case string:
		fmt.Println("(string): ",k, value)
	case float64:
		fmt.Println("(float64):",k, value)
	case bool:
		fmt.Println("(bool):",k, value)
	case []interface{}:
		fmt.Println("(slice):", k)
		for i, j := range value {
			fmt.Println(" ", i,j)
		}
	default:
		fmt.Println("(unknown): ",k,value)
		
	
	}

}
}
