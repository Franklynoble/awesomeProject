package main

import "fmt"

type Speaker interface {
	Speak() string
}
type Stringer interface {
	String() string
}
type person struct {
	name      string
	age       int
	isMarried bool
}
type cat struct {
}

func main() {
	p := person{name: "Caily",
		age:       44,
		isMarried: true,
	}
	fmt.Println(p.Speak())
	fmt.Println(p.String())

	c1 := cat{}
	chatter(c1)

}

// Create a String() method for person and return a string value. This will satisfy
//the Stringer{} interface, which will now allow it to be called by the fmt.Println()
//method:
func (p person) String() string {
	return fmt.Sprintf("%v(%v years old).\nMarried status: %v", p.name,
		p.age, p.isMarried)
}

//Create a Speak() method for person that returns a string. The person type has a
//Speak() method that has the same signature as the Speak() method of the Speaker{}
//interface. The person type satisfies the Speaker{} interface by having a Speak()
//method that returns the string. To satisfy interfaces, you must have the same
//methods and method signatures of the interface:
func (p person) Speak() string {
	return "Hi my name is: " + p.name
}
//• In the preceding code, we declare a cat type and create a method for the cat type
//called Speak(). This fulfills the required method sets for the Speaker{} interface.
//• We create a method called chatter that takes the Speaker{} interface as an
//argument.
//• In the main() function, we are able to pass a cat type into the chatter function,
//which can evaluate to the Speaker{} interface This satisfies the required method
//sets for the interface.

func chatter(s Speaker){
	fmt.Println(s.Speak())
}

func (c cat) Speak() string {

	return "Pure mmeoo!!"
}
