package main

import "fmt"

func main() {
	var arr [10]int
	arr = fillArray(arr)
	arr = opArray(arr)
	fmt.Println(arr)




	 var arr1[4]string

	 arr1[0] = "Kok"

	arr1[1] = "O,O,O"

	arr1[2] =  "One"

	arr1[3] = "Olumba,Olumba"
	for i := 0; i < len(arr1); i ++ {

		if arr1[i] == "Kok"{
			arr1[i] = "King,of Kings"
		}
		fmt.Println(arr1[i])
	}






	//fmt.Println(n)
	//fmt.Println(opertionOnArray())
	// allMessages :=  writeToArray()
	// fmt.Println(allMessages)

	//com1,com2,arr3 := arrayInit()
	//fmt.Println("[10]int ==[...]{9:0} :",com1)
	//fmt.Println("[10]int==[10]int{1,9:10,4:5}:",com2)
	//fmt.Println("arr3 :",arr3)
	//

	//comp1,comp2,comp3 := compArrays()
	//fmt.Println(comp1,comp2,comp3)

	//fmt.Printf("%#v\n",defineArray())
}
func defineArray() [10]int {
	var arr [10]int
	return arr
}

//compare arrays
func compArrays() (bool, bool, bool) {
	var arr1 [5]int
	arr2 := [5]int{0}
	arr3 := [...]int{0, 0, 0, 0, 0}
	arr4 := [5]int{0, 0, 0, 0, 9}
	return arr1 == arr2, arr1 == arr3, arr1 == arr4
}

//initializing arrays using key

func arrayInit() (bool, bool, [10]int) {
	var arr1 [10]int
	arr2 := [...]int{9: 0}
	arr3 := [10]int{1, 9: 10, 4: 5}

	for i := 0; i < len(arr3); i++ {
		fmt.Println("from Array 3:", arr3[i])
	}
	return arr1 == arr2, arr1 == arr3, arr3
}

// reading a single item from an Array

func message() string {
	arr := [...]string{
		"Ready",
		"Get",
		"Go",
		"To",
	}

	return fmt.Sprintln(arr[0], arr[1], arr[2], arr[3])
}

// Writing to An Array
func writeToArray() string {

	arr := [4]string{"ready ", "Get ", "Go ", "to "}

	arr[1] = " It's"
	arr[0] = " time"
	return fmt.Sprint(arr[1], arr[0], arr[3], arr[2])

}

// Looping Over An Array

func opertionOnArray() string {
	m := ""
	arr := [4]int{1, 2, 3, 4}
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] * arr[i]
		m += fmt.Sprintf("%v: %v\n ", i, arr[i])
	}
	return m
}

func arrayModif(n []int) {
  count := 5
 //  var ans[]int
for i := 0; i < count; i ++ {

	n[count] = n[i]
	 fmt.Println(n[i])
}

}

func fillArray(arr[10]int)[10]int{
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}
	return arr
}

func opArray(arr [10]int)[10]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i]* arr[i]
	}
	return arr
}
