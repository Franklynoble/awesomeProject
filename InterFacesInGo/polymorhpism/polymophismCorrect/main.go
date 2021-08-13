package main

import "fmt"

type Speaker interface {
	Speak() string
}
type cat struct {
}
type dog struct {
	name string
}
type person struct {
	name string
}

func main() {
	c := cat{}

	d := dog{}
	//p := person{name: "James"}
	// cc := []Speaker{c,d,p}
	saySomething(c,d)

}
//Our saySomething() function is using a variadic parameter. If you recall, a variadic
//parameter can accept zero or more arguments for that type. For more information
//on variadic functions, review Chapter 5, Functions. The parameter type is Speaker. An
//interface can be used as an input parameter:
func saySomething(say...Speaker) {
    for _,s := range say {
    	fmt.Println(s.Speak())
    }

}
func (c cat) Speak() string { //cat  implementation of Speaker InterFace
	return "Purr Meow"
}
func (d dog) Speak() string { //dog implementation of Speaker InterFace
	return "Woof Woof"
}
func (p person) Speak() string { // person implementation of Speaker InterFace
	return "Hi my name is " + p.name + "."
}

//Each of our types implicitly implements the Speaker{} interface. Each of the concrete
//types implements it differently from the others:

