package main

import "fmt"

type user struct {
	name string
	age int
	balance float64
	member bool
}

type point struct {
	x int
	y int
}
//here we create function that returns two boolean
func compare()(bool,bool) {
	point1 := struct {
		x int
		y int
	}{10,10}
	// with the second anonymous struct, were initializing it to zero and then changing the value
    point2 := struct{
    	x int
    	y int
	}{}
    point2.x = 10
    point2.y = 5
    // the final struct to create uses the named struct type we created previously
    point3 := point{10,10}

    //compare them. then return and close the function
    return point1 == point2,point1 == point3
}

//We'll create a function that returns a
//slice of our newly defined struct type:
func getUsers()[]user{
	u1 := user{
		name: "Tracy",
		age: 51,
		balance: 98.43,
		member: true,
	}
	u2 := user{
		age: 19,
		name: "Nick",
	}
	u3 := user{
		"Bob",
		25,
		0,
		false,
	}
	//This var notation will create a
	//struct where all the fields have zero values:
	var u4 user
		u4.name = "Sue"
		u4.age = 31
		u4.member = true
		u4.balance = 17.09
		//Now, we can set values on the fields using . and the field name:
		return []user{u1,u2,u3,u4}
}
func main() {

	a,b := compare()
	fmt.Println("point1 == point2: ",a)
	fmt.Println("point1 == point3: ",b)

	//
	//users := getUsers()
 //
 //for i := 0; i < len(users); i++ {
 //	fmt.Printf("%v: %#v\n",i,users[i])
 //
 //}
}
