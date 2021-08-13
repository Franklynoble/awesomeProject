package main

import "fmt"

func main() {
	l1,l2,l13 := linked()
	fmt.Println("Linked: ", l1,l2,l13)
 nl1,nl2 := noLink()
 fmt.Println("No Link : ",nl1,nl2)

 cl1,cl2 := capNoLink()
 fmt.Println("cap No Link: ",cl1,cl2)

 copy1,copy2,copied := copyNoLink()
 fmt.Println("Copy No link: ",copy1,copy2)
 fmt.Printf("Number of elements copied %v)\n",copied )

 //fmt.Printf()

	//fmt.Println(message())
	//s1, s2, s3 := getSlice()
	////fmt.Println(s1,s2,s3)
	//fmt.Printf("s1: len = %v cap = %v\n", len(s1), cap(s1))
	//fmt.Printf("s1: len = %v cap = %v\n", len(s2), cap(s2))
	//fmt.Printf("s1: len = %v cap = %v\n", len(s3), cap(s3))
}

func message() string {

	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//We'll extract the first value, first directly as an int, then as a slice using both low
	//and high, and finally using just high and skipping low. We'll write the values to a
	//message string:
	m := fmt.Sprintln("First :", s[0], s[0:2], s[3:])
	fmt.Println(m)
	m += fmt.Sprintln("Last : ", s[len(s)-1], s[len(s)-1:len(s)], s[len(s)-1:])
	fmt.Println(m)

	// let's get the first five values and add them to the message:
	m += fmt.Sprintln("First 5: ", s[:5])
	//we'll get the last four values and add them to the message as well:
	m += fmt.Sprintln("Last 4: ", s[5:])
	fmt.Println(m)
	m += fmt.Sprintln("Middle 5:", s[2:7])
	return m
}

func getSlice() ([]int, []int, []int) {
	var s1 []int
	s2 := make([]int, 10)
	s3 := make([]int, 10, 50)
	return s1, s2, s3
}

func linked()(int, int, int) {
	s1 := []int{1, 2, 3, 4, 5}
	//we'll make a simple variable copy of that slice:
	s2 := s1
	//Create a new slice by copying all the values from the first slice as part of a slice
	//range operation:

	s3 := s1[:]
//Change some data in the first slice. Later, we'll see how this affects the second and
	//third slice:
	s1[3] = 99

	return s1[3],s2[3],s3[3]

}
func noLink()(int,int) {
	s1:= []int{1,2,3,4,5}
	s2 := s1
	s1 = append(s1,6)
	s1[3] = 99
	return s1[3],s2[3]
}

func capNoLink() (int, int) {
	s1 := make([]int, 5, 10)
	s1[0], s1[1], s1[2], s1[3], s1[4] = 1, 2, 3, 4, 5
	s2 := s1
	s1 = append(s1, []int{10: 11}...)
	s1[3] = 99
	return s1[3], s2[3]
}
func copyNoLink()(int,int,int) {
	s1 := []int{1,2,3,4,5}
	s2 := make([]int, len(s1))
	copied := copy(s2,s1)
	s1[3] = 99
	return s1[3],s2[3],copied
}
func appendNoLink()(int,int){
	s1 := []int{1,2,3,4,5}
	s2 := append([]int{}, s1...)
	s1[3] = 99
	return s1[3],s2[3]

}