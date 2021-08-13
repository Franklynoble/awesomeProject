package main

import "fmt"

type Speaker interface {
	Speak()string
	//Greet()string
}
type cat struct{

}

type cow struct {

}
func main() {

	c := cat{}
	fmt.Println(c.Speak())
	c.Greeting()


}
	func (c cat )Speak ()string {
    return "Purr Meow"
	}



func (n cat) Greeting(){

	fmt.Println("Meow,Meow!!! mmmeeeoeww")
}


