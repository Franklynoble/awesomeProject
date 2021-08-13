package main

import (
	"fmt"
	"os"
)

func main() {
	//check that at least one argument gets passed in. if not,Exit:
	if len(os.Args) < 2 {
		fmt.Println("User ID not Passed")
		os.Exit(1)
	}
	//capture the passed argument and callmaingo run  the get user function
	userID := os.Args[1]
	name,exists := getUser(userID)
 // if the key is not found , print a message ,and the  print all the users using a range
  // loop After that,Exit
	if !exists{
		fmt.Printf("Passed user Id (%v) not found.\nUsers:\n",userID)
		for key,value := range getUsers() {
			fmt.Println("ID: ", key, "Name: ",value)
		}
		os.Exit(1)
	}
	// if Everything is okay, print the name we found
	fmt.Println("Name:",name)

}

func getUsers() map[string]string {
	//Define a map and initialize it with data. Then, return the map and close the function:

	return map[string]string{
		"305":"Sue",
		"204" :"Bob",
		"631":"Jake",
		"073":"Tracy",
	}
}

func getUser(id string)(string,bool) {
	users := getUsers()
	user, exists := users[id]
	//Return both values and close the function:
	return user,exists

	}



