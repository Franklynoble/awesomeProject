package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
)

//Create a struct to be our client-side user model:
type UserClient struct {
	ID   string
	Name string
}

//Create a struct to be our client-side transaction. Tx is a common shorthand for
//transaction:
type TxClient struct {
	ID          string
	User        *UserClient //points to the UserClient add
	AccountFrom string
	AccountTo   string
	Amount      float64
}

//Create a struct to be our server-side user model. This model doesn't match the
//client model because it doesn't have the Name property:
type UserServer struct {
	ID string
}

//Create a struct to be our server-side transaction. Here, the user is not a pointer.
//The amount is a pointer, however, and the pointer is for a float32, not a float64:
type TxServer struct {
	ID          string
	User        UserServer
	AccountFrom string
	AccountTo   string
	Amount      *float32
}

func main() {
	//create the dummy netWork, which is a buffer from the bytes package
	var net bytes.Buffer

	//create the dummy data using the client-side structs:
	clientTx := &TxClient{
		ID:"123456789",
		User: &UserClient{
			ID: "ABCDEF",
			Name: "James",
		},
		AccountFrom: "Bob",
		AccountTo: "Jane",
		Amount: 9.99,
	}

	//Encode the data. The Target for the encoded data is our dummy network
	enc := gob.NewEncoder(&net)

	//check for errors and exit if any are found
	if err := enc.Encode(clientTx);err != nil {
		log.Fatal("error",err)
	}
	// send the data to the server
	serverTx,err := sendToServe(&net)

	if err != nil {
		log.Fatal("Server Error: ",err)
	}
	fmt.Printf("%#v\n",serverTx)


}

 //Create our sendToServer function. this function takes  a single io.Reader interface
 //and returns a server-side transaction
func sendToServe(net io.Reader)(*TxServer,error) {
	//create a variable to be the Target for decoding
	tx := &TxServer{}

	//create a decoder with the network as the Source
	dec := gob.NewDecoder(net)
	//Decode and capture an Error
	err := dec.Decode(tx)
	return tx,err
}