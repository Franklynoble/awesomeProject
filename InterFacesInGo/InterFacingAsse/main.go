package main

import "fmt"
type cat struct {
	name string
}

func main () {
	//• A type assertion returns two values, the underlying value and a Boolean value.
	//• isValid is assigned to a return type of bool. If it returns true, that indicates that str
	//is of the int type. It means that the assertion is true. We can use the Boolean that
	//was returned to determine what action we can take on str.
	//• When the assertion fails, it will return false. The return value will be the zero value
	//that you are trying to assert to. It also will not panic.
	var str interface{} = "the book club"
	v, isValid := str.(string)
	fmt.Println(v,isValid)

	c := cat{"oreo"}
    i := []interface{}{42,"The Book club ",true,c}
    typeExample(i)

}


func typeExample(i []interface{}) {
	//i.(type)
	//The syntax is similar to that of the type assertion, i.(int), except the specified type, int
	//in our example, is replaced with the type keyword. The type being asserted of type i is
	//assigned to v; then, it is compared to each of the case statements.
	//case S:
	//In the switch type, the statements evaluate for types. In regular switching, they evaluate
	//for values. Here, it is evaluated for a type of S.
	for _,x := range i {
		switch v := x.(type) {
		case int:
			fmt.Printf("%v is as\n",v)
		case string:
			fmt.Printf("%v is a ",v)
		case bool:
			fmt.Printf("%v is a\n ",v)
		default:
			fmt.Printf("Unknown type %T\n",v)

		}
	}
}

