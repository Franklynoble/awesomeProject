package main

import (
	"fmt"
)

func main() {

// nonempty is an example of an in-place algorithm
/*
	data := [] string{"one","", "three"}
	fmt.Printf("%q\n",nonempty(data))  // `["one", "three"]`
	fmt.Printf("%q\n",data )
*/

	// nonempty is an example of an in-place algorithm
	data1 := [] string{"one","", "three", "four", "Four"}

	fmt.Printf("%q\n",nonempty2(data1))  // `["one", "three"]`
	fmt.Printf("%q\n",data1 )

	fmt.Print("removing from index position")
	s := []int{5,6,7,8,9}
	fmt.Println(remove(s,2)) // "[5 6 8 9]" // remove from s @ index position 2


}


func nonempty(strings []string) [] string {
	i := 0
	for _,s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
	/*
	The subtle part is that the input slice and the output slice share the same underlying array.
	This avoids the need to allocate another array, thoug h of course the contents of data are par tly
	overwritten, as evidence d by the second print statement:
	 */
}

/*

Thus we would usually write: data = nonempty(data).
The nonempty function can also be written using append:
 */

func nonempty2(strings []string) [] string {

	out := strings[:0] // zero-length slice of original
	for _,s := range strings {
		if s != ""{
			out = append(out, s)
		}
	}
	return out
}


/*
To remove an Element from the Middle of a slice, preserving the order of the remaining elements,use
copy to slide the higher-numbered elements down by one to fill the gap
 */
func remove(slice []int, i int) []int {
	copy(slice[i:],slice[i+1:]) // here slide the index backwards to cover the gap where element is Rmoved
	return slice[:len(slice)-1]
}

/*
if we do not want to preserve the order we can just move the last element into the gap

 */
func remove2(slice []int, i int) []int{
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}