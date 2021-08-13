package main

import (
	"fmt"
	"os"
)


func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n ", minArgs)
		os.Exit(1)
	}
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

func findLongest(args []string) string {
	var longest string
	for i := 0; i < len(args); i++ {
		if len(args[i]) > len(longest) {

			longest = args[i]

		}
	}
	return longest
}

func main() {
	// The first argument
	// is always program name
	myProgramName := os.Args[0]

	// this will take 4
	// command line arguments
	cmdArgs := os.Args[4]

	// getting the arguments
	// with normal indexing
	gettingArgs := os.Args[2]

	toGetAllArgs := os.Args[1:]

	// it will display
	// the program name
	fmt.Println(myProgramName)

	fmt.Println(cmdArgs)

	fmt.Println(gettingArgs)

	fmt.Println(toGetAllArgs)

}
