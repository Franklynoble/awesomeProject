package main

import (
	"fmt"
	"reflect"
)

func main() {
	//	n := []interface{1, 2, 3, 4, 5}
	//	//total := 0
	//	var total float64
	//	for i := 0; i < len(n); i++ {
	//
	//
	//		total += n[i]
	//		}
	//	fmt.Println(total)
	c := []int{1,2,3,4,5}
	f := all(c...)
	fmt.Println(f)
}

func all(n ...int) float64 {
	var total = 0
	//var  ff float64
	// var final float64
	for i := 0; i < len(n); i++ {
		  total += n[i]
			total = len(n)/total
			 

	}
	return float64(total)

}





func sum(in interface{}) (int64, error) {
	res := int64(0)
	v := reflect.ValueOf(in)
	switch v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int32, reflect.Int64:
		res = v.Int()
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			sliceRes, err := sum(v.Index(i).Interface())
			if err != nil {
				return 0, err
			}
			res = res + sliceRes
		}
	case reflect.Map:
		for _, k := range v.MapKeys() {
			mapRes, err := sum(v.MapIndex(k).Interface())
			if err != nil {
				return 0, err
			}
			res = res + mapRes
		}
	default:
		return 0, fmt.Errorf("input passed was invalid.")
	}
	return res, nil
}
