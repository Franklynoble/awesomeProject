package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Stat("junk.txt")

	if err != nil {
		if os.IsNotExist((err)) {
			fmt.Println("junk.txt: File does not Exist")
			fmt.Println(file)
		}
	}
	fmt.Println()
	file, err = os.Stat("test.txt")
	if err != nil {
		if os.IsExist((err)) {
			fmt.Println("Test.txt: File does not exist! ")
		}
	}
	fmt.Printf("File name: %s\nIsDir: %tnModTime: %vnMode: %v\nSize: %d\n",file.Name(),file.IsDir(), file.ModTime(),file.Mode(), file.Size())
    file, err = os.Stat("junk.txt")
    if os.IsExist(err) {
    	fmt.Println("junk.txt: file does not exist!")
    	fmt.Println(file)
	}
}
