package main

import "fmt"

type Speaker interface {
	Speak()string
}
//is an interface that is a type that can describe itself as a string. Interface names
//typically follow the method name but with the addition of the er suffix:
type Stringer interface {
	String()string
}
type cat struct {
	name string
	age int
}


func main() {
	c := cat{
		name: "Oreo",
		age: 9,
	}
	fmt.Println(c.Speak())
	fmt.Println(c.String())

}

func (c cat)Speak()string {
	return "Purr Meow"
}

func (c cat) String()string {
	return fmt.Sprintf("%v is (%v years Old)",c.name, c.age)
}