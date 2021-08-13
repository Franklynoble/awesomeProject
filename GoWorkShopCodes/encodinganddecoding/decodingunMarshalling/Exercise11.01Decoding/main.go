package main

import (
	"encoding/json"
	"fmt"
)

type student struct {
	StudentId     int      `json:"id"`
	LastName      string   `json:"lname"`
	MiddleInitial string   `json:"minitial"`
	FirstName     string   `json:"fname"`
	IsEnrolled    bool     `json:"enrolled"`
	Course        []course `json:"classes"`
}
type course struct {
	Name   string `json:"coursename"`
	Number int    `json:"coursenum"`
	Hours  int    `json:"coursehours"`
}

func main() {

	// in the main function ,add the hson data that we will be unmarshalling into the (student and course)

	data := []byte(`{"id":123,
		             "lname":"Smith",
                        "minitial":"null"
                            ,"fname":"John",
                                "enrolled": true,
                            "classes":
          [{"coursename":"Intro to Golang",
                             "coursenum":101,
               "coursehours":3},
{

"coursename":"Engliish Lit",
     "coursenum":101,
             "coursehours":3},
			{"coursename":"World History",
										
				"coursenum":101,
					"coursehours":3}
                             ]
}
`)
	var s student

	//now we will unmarshal the json into our student
	err := json.Unmarshal(data, &s)
	if err != nil {
		fmt.Println(err)
	}
	//We will print the student struct so that we can see that all the data from the JSON
	//is present:
	fmt.Println(s)

}
