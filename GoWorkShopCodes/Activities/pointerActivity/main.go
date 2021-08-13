package main

import "fmt"
//activity passed
func main() {
	a,b := 5,10
	fmt.Println(a,b)
	//swap1(n1,n2)
	//fmt.Println(b,a)
	  swap(&a,&b)

	fmt.Print(a == 10 , b == 5)

}

func swap(a *int, b *int) {
	   *a,*b = *b,*a

	    fmt.Println(*a,*b)


}
func swap1(a int, b int ) {

	a,b = b,a
	b,a = a,b
	fmt.Print(b,a)

}
