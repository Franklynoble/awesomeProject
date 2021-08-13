package main

import (
	"fmt"
	"sort"
)

func main(){
	  //   data := []int{1,2,3,4,5,6,7,8,9,10,11,12,13}
		//

	  	    ////// 	fmt.Print(data[i],)
		//
		// }
		// reverse(data)

	     strData := []string{ "abby", "james","coddar", "albert","abby","abby"}
	   // sort.Strings(strData)

	     //  eliminate(strData)
	     eleminate2(strData)

//	fmt.Print(strData)
}


func reverse( arr []int) [] int  {

	if arr == nil{
		return nil
	}else {
		for i := len(arr)-1; i >= 0; i-- {

			 fmt.Print(arr[i])
		}
	}
	return arr[:len(arr)]
}

func eliminate( ss []string )  {
	for i:= 0; i < len(ss) ; i++ {

		if ss[i] == ss[i] {
			ss[i] = ss[i]
	       // ff :=  copy(ss[i:],ss[i+1:])
		//	stringRemove(ss,i)
			fmt.Print(stringRemove(ss,i))
			return
		//  fmt.Print(stringRemove(ss[:len(ss)],i ))

		}

	}


}
func eleminate2(in []string){

	//in := []int{3,2,1,4,3,2,1,4,1} // any item can be sorted
	sort.Strings(in)
	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++  // now increase the counter
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set wha`+-*t is required
		in[j] = in[i]
	}
	result := in[:j+1]
	fmt.Println(result)
}
func stringRemove( slice[]string, i int ) []string{
 // use this to remove string @ a particular index.html
		copy(slice[i:],slice[i+1:]) // here slide the index.html backwards to cover the gap where element is Rmoved
		return slice[:len(slice)-1]
}