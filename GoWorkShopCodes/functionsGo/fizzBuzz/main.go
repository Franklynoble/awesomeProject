package main

import "fmt"

func fizzBuzz(){
	for i := 0; i < 30; i ++ {
		if i% 15 == 0{
			fmt.Println("FizzBuzz")
		}else if i%3 == 0 {
			fmt.Println("Fizz")
		}else if i%5 == 0 {
			fmt.Println("Buzz")
		}else {
			fmt.Println(i)
		}
	}
}


func itemsSold(){
	items := make(map[string]int)
	items["John"] = 41
	items["Celina"] = 109
	items["Micah"] = 24

	for k,v := range items {
		fmt.Printf("%s sold %d items and ",k,v)
		if v < 40 {
			fmt.Println("is below Expectation")
		}else if v > 40 && v <= 100 {
			fmt.Println("Meets expectations.")
		}else if v > 100 {
			fmt.Println("Exceed expectation")
		}
	}
}

func fizBuzzWithParam(end int) {
 for i := 1; i < end; i ++ {
 	if i%15 == 0 {
 		fmt.Println("FizzBuzz")
	}else if i%3 == 0 {
		fmt.Println("Fizz")
	}else {
		fmt.Println(i)
	}
 }
}

func greeting(name string, age int){
	fmt.Printf("%s is %d", name,age)
}
func main() {
	fmt.Println("Main is in Control")
	//itemsSold()
//greeting("Cayden",45)
	//fizzBuzz()

    fizBuzzWithParam(100)
	fmt.Println("Back to main")

}


