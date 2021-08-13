package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {

//sigs := make(chan bool)

	//Next, add a done channel. The done channel is used to let us know when the
	//program can exit:
//	done := make(chan bool)

	//we will add signal.Notify method. the Notify method works by
	//sending  values of the os.Signal type to a channel


//Recall that the last parameter of the signal.Notify method is a variadic parameter
	//of the os.Signal type.




	sigs := make(chan os.Signal, 1)
	done := make(chan bool)
//	signal.Notify(sigs, syscall.SIGINT,syscall.)

	go func() {
		for {
			s := <-sigs
			switch s {
			case syscall.SIGINT:
				fmt.Println()
				fmt.Println("My process has been interrupted.  Someone might of pressed CTRL-C")
				fmt.Println("Some clean up is occuring")
				cleanUp()
				done <- true
			//case syscall.SIGTSTP:
				fmt.Println()
				fmt.Println("Someone pressed CTRL-Z")
				fmt.Println("Some clean up is occuring")
				cleanUp()
				done <- true
			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught(ctrl-z, ctrl-c)")
	<-done
	fmt.Println("Out of here")
}

func cleanUp() {
	fmt.Println("Simulating clean up")
	for i := 0; i <= 10; i++ {
		fmt.Println("Deleting Files.. Not really.", i)
		time.Sleep(1 * time.Second)
	}
}



