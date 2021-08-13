package main

import "fmt"

type Speaker interface {
	Speak()string
}
type cat struct {

}
type dog struct {
	name string
}
type person struct {
	name string
}


func main () {
	c := cat{}

	d := dog{}
	p := person{name: "James"}

	catSpeak(c)
	dogSpeak(d)
	personSpeak(p)

}
func (c cat)Speak()string { //cat  implementation of Speaker InterFace
	return "Purr Meow"
}
 func (d dog)Speak()string { //dog implementation of Speaker InterFace
 	return "Woof Woof"
 }
func (p person) Speak()string { // person implementation of Speaker InterFace
	return "Hi my name is "+p.name +"."
}
//Each of our types implicitly implements the Speaker{} interface. Each of the concrete
//types implements it differently from the others:
func catSpeak(c cat){
	fmt.Println(c.Speak())
}
func dogSpeak(d dog) {
	fmt.Println(d.Speak())
}
func personSpeak(p person) {
	fmt.Println(p.Speak())
}
