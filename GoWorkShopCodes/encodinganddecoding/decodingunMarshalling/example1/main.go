package main

import (
	"encoding/json"
	"fmt"
)

type greeting struct {
	Message string
}

type person struct {
	Lastname  string  `json:"lname"`
	Firstname string  `json:"fname"`
	Address   address `json:"address"`
}

type address struct {
	Street  string `json: "street"`
	City    string `json:"city"`
	State   string `json:"state"`
	ZipCode int    `json:"zipcode"`
}

type greeting1 struct {
	SomeMessage string `json:"anymessage"` // this naymessages corresponds to the json Key value
}

func main() {

disp()


	fmt.Println("*****************************")

	//FIRST JSON
	data := []byte(`{"message": "Greetings fellow gopher!"
}`)
	//The json.Unmarshal struct requires that the JSON encoded data must be a byte of slices:
	var v greeting
	err := json.Unmarshal(data, &v)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v.Message)

	//Second JSON
	//second unmarshalling
	data2 := []byte(`{"anymessage": "Greeting Fellows from 2"}`)
	var g greeting1

	err = json.Unmarshal(data2, &g)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(g.SomeMessage)

}

func disp() {
	data := []byte(`{"lname":"Smith",
"fname":"John",
"address":{
"street":"Sulphur springs Rd",
"city":"Park City",
"state":"VA",
"zipcode":12345}
}`)

var p person

err := json.Unmarshal(data,&p)
if err != nil {
	fmt.
		Println(err)
}
fmt.Println(p)

}
