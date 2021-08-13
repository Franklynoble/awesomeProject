package main

import "fmt"

type Speaker interface {
	Speak()string
}
type cat struct {

}


func main() {
	c := cat{}
   catSpeak(c)
}

//cat matches the Speak() method of the Speaker{} interface, so a cat is a Speaker{}:
func (c cat) Speak() string  { // implementing the Speaker interface
	return  "Purr Meow"
}

func catSpeak(c cat) {
	fmt.Println(c.Speak())
}
