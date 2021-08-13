package main

import (
	"errors"
	"fmt"
)

func doubler(v interface{}) (string, error) {
	//create  switch using our Argument
	switch t := v.(type) {
	case string:
		return t + t, nil
	case bool:
		if t {
			return "truetrue", nil
		}
		return "false,false ", nil
	//For the floats, we're matching on more than one type. This means we need to do
	//type assertion to be able to work with the value:

	case float32, float64:
		if f, ok := t.(float64); ok {
			return fmt.Sprint(f * 2), nil
		}
		//if this type assertions were to fail, we'd get a panic, but we can rely on the login
		// logic that floa32 can work directly with the result of type assertion:
		return fmt.Sprint(t.(float32) * 2), nil
	//Match all of the int and uint types. We've been able to remove lots of code here by
	//not needing to do the type-safety checks ourselves
	//
	case int:
		return fmt.Sprint(t * 2), nil
	case int8:
		return fmt.Sprint(t * 2), nil
	case int16:

		return fmt.Sprint(t * 2), nil
	case int32:

		return fmt.Sprintf("%T\n", t * 2), nil
	case int64:
		return fmt.Sprint(t * 2), nil
	case uint:
		return fmt.Sprint(t * 2), nil
	case uint16:
		return fmt.Sprint(t * 2), nil
	case uint32:
		return fmt.Sprint(t * 2), nil
	case uint64:
		return fmt.Sprint(t * 2), nil
	default:
		return "", errors.New("unsupported type passed")
	}
}



func main() {

	res, _:= doubler(-5)
	fmt.Println("-5 :",res)
	res,_ =doubler(5)
	fmt.Println("5 :",res)
  res,_ = doubler("Yum")
  fmt.Println("Yum: ",res)
  res,_ = doubler(true)
  fmt.Println("True: ",res)
  res, _ =doubler(float32(3.14))
  fmt.Printf("3.14: %#\n",res )
}
