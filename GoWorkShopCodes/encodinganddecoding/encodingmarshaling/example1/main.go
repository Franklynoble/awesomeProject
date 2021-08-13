package main

import (
	json2 "encoding/json"
	"fmt"
)

type greeting struct {
	SomeMessage string
}
//We have two structs: a person struct and an address struct. The address struct is
//embedded inside the person struct. Both structs have the JSON key names defined in
//the JSON tags. The address struct will be a separate object inside the JSON:
type person struct {
	LastName string `json:"lname"`
	FirstName string `json:"fname"`
	Address address `json:"address"`
}
type address struct{
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	Zipcode int  `json:"zipcode"`
}
type book struct {
	// using omit empty to exclude file not to print when there is no value, instead zero value or empty string
	ISBN string `json:"isbn"`
	Title string `json:"title"`
	YearPublished int `json:"yearpub,omitempty"`
	Author string `json:"author"`
	CoAuthor string `json:"coauthor,omitempty"`
}
type book2 struct {
	// using omit empty to exclude file not to print when there is no value, instead zero value or empty string
	//In the above code, `json:",omitempty"` does not have a value for a field. Notice the
	//JSON tag value starts with a comma.
	//• `json:",omitempty"` will have the field in the JSON if there is a value for the key. If
	//Author has a value set, it will appear in the JSON as the "Author" :"somevalue" key:

	ISBN string `json:"isbn"`
	Title string `json:"title"`
	YearPublished int `json:",omitempty"`
	Author string `json:",omitempty"`
	CoAuthor string `json:"-"`
}


func main () {

	peExec()

	//book1exc()
	fmt.Println("++++++++++++++")

	var b book
	b.ISBN = "9933HIST"
	b.Title = "Greatest of all times"
	b.Author = "John Adams"

	 json1, err := json2.Marshal(b)
    if err != nil {
    	fmt.Println(err, "ERROR!")
	}
	fmt.Printf("%s",json1)

    fmt.Println("*******************")


















	var v greeting
	//Since we did not provide the JSON tag for the struct greeting, SomeMessage, the Go
	//Marshal encodes the exportable fields and its values. The Go Marshal uses the name of
	//the field, SomeMessage ,as the name of the key field in the JSON data.
	v.SomeMessage = "Marshal Me!"

	//When we call the Marshal function, we are passing it a struct. The function will return
	//back an error and the JSON encoding of g.
	json,err := json2.Marshal(v)
   if err != nil {
   	fmt.Println(err)
   }
   fmt.Printf("%s",json)

}

func book1exc() {
var b book2

b.ISBN = "9933HIST"
b.Title = "Greatest of all Books"
b.Author = "John Adams"
b.CoAuthor = "can't see me"
json,err := json2.Marshal(b)

if err != nil {
	fmt.Println(err)
}
fmt.Printf("%s",json)
}


func peExec() {
	p := person{LastName: "Vader", FirstName: "Darth"}
	p.Address.Street = "Galaxy Far Away"
    p.Address.City = "Dark Side"
    p.Address.State = "Tatooine"
    p.Address.Zipcode = 12345

//The prettyPrint variable is the JSON encoding of p, by using json.MarshalIndent(). We
	//set the prefix argument to an empty string and the indent argument to four spaces.
	//As with the json.Marshal() function, we also check for any errors returned from
	//the json.MarshalIndent() function. We can see these various steps using the json.
	//MarshalIndent() method depicted in the following diagram:
	//Figure
	json,err := json2.MarshalIndent(p,"","  ")

       if err != nil {
       	fmt.Println(err)
	   }
	 //  fmt.Println("first Print :",json)
	fmt.Println()
	   fmt.Println(string(json)) // the json marshalled file Must be Encoded
}