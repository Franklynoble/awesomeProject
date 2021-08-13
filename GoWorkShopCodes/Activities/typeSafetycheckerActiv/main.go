package main

import (
	"fmt"
)
//the Activity

func getData()[]interface{} {
	return []interface{}{
		1,
		3.14,
		"Hello",
		true,
		struct{}{},

	}
}


func main() {
data := getData()
for i := 0; i < len(data); i++ {
	fmt.Printf("%v is %v\n",data[i],check1(data[i]))
}

//ch := true
//	cc := float32(2.3)
	//fmt.Println(cc)
//	name := "James"


}

func passAll( c interface{}) interface{}{
	return nil
}

func check1(v interface{}) (c interface{}) {
	m := ""
	switch v.(type) {

	//case int:
	//	if b,ok := t.(int); ok {
	//	   m,_ := fmt.Printf("%T\n",b)
	//        d = strconv.Itoa(m)
	//	   return  fmt.Sprintf("%T\n",d)
	//	}
	//case all:
	//	   for k,y := range all{
    //        fmt.Println(k,y)
	//	   }
	//
	//	m = fmt.Sprintf("%#v is %T\n", v, v.(all))
	//	c = m
	//	return c
	//}

		case float64:
			m = fmt.Sprintf("%#v is %T\n", v, v.(float64))
			c = m
			return c
		case float32:
			m = fmt.Sprintf("%#v is %T\n", v, v.(float32))
			c = m
			return c
		case string:
			m = fmt.Sprintf("%#v\nis %T\n", v, v.(string))
			c = m
			return c

//
//case struct:
//	m = fmt.Sprintf("%#v\nis %T\n", v, v.(string))
//	c = m
//	return c
		default:
			m = fmt.Sprintf("%T\n is Unkown ", v)
				c = m
				return c
			}
		}

