package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//openFileCreate()
	//	openFileReadWrite()
	openCreatAppend()

}

func readAll() {
	f, err := os.Open("testall") // this would exist without err, only if the file is present but empty

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file contents: ")
	fmt.Println(string(content))
	r := strings.NewReader("no File here.")
	c, err := ioutil.ReadAll(r)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Println("contents of strings.NewReader: ")
	fmt.Println(string(c))
}

func openFileCreate() {

	//Using os.OpenFile with the os.O_CREATE file mode will create the junk101.txt file if
	//it does not exist, and then open it.
	f, err := os.OpenFile("junk101.txt", os.O_CREATE, 0644) // fil mode 0644 gives you the Rom to read/write file Mode
	if err != nil {
		panic(err)
	}
	defer f.Close()

}

func openFileReadWrite() {
	//Using os.OpenFile with the os.O_CREATE flag will create the junk101.txt file if it does
	//not exist and then open it. If it does exist, it will just open the file. It will also allow
	//the reading and writing of the file while it is open because of the os.O_WRONLY flag:
	f, err := os.OpenFile("junk102", os.O_CREATE|os.O_WRONLY, 0644)
	//• Since we used the os.O_WRONLY flag, we can write to the file while it is open.
	if _, err = f.Write([]byte("adding stuff\n")); err != nil {
		panic(err)
	}
	//after creating the file now reading the file to console
	//• Since we used the os.O_WRONLY flag, we can write to the file while it is open.
	fs, err := os.ReadFile("junk102")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(fs))

}

func openCreatAppend() {

	f, err := os.OpenFile("Junk103", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err := f.Write([]byte("adding stuff to this File\n")); err != nil {
		panic(err)
	}
	f, err = os.OpenFile("Junk103", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if cvc, err := os.ReadFile("Junk103"); err != nil {
		panic(err)
	} else {
		fmt.Println(string(cvc))
	}

}
