package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	sigs := make(chan os.Signal, 4)

	// We create a channel of the os.Signal
	//The Notify method works by sending
	//values of the os.Signal type to a channel. The sigs channel is used to receive these
	//notifications from the Notify method:

	//The done channel is used to let us know when the program can exit:
	done := make(chan bool)

	//The signal.Notify method will receive notifications on the sigs channel, which is
	//of the syscall.SIGINT type:
	signal.Notify(sigs, syscall.SIGINT)

	go func() {

		for {
			s := <-sigs // pass sig chanel to s and use s variable in switch
			switch s {
			case syscall.SIGINT:
				fmt.Print()
				fmt.Println("My process has been interrupted. Someone might have pressed CTRL c")
				fmt.Println("some clean up is occurring")

				done <- true

			}
		}
	}()
	fmt.Println("Program is blocked until a signal is caught")
	<- done
	fmt.Println("Out of here")
}
