package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var (
	ErrWorkingFileNotFound = errors.New("The Working file is not found")
)

func main() {
	backupFile := "backUpFile"
	workingFile := "notes"
	data := "note"

	err := createBackup(workingFile, backupFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for i := 1; i <= 10; i++ {
		note := data + " " + strconv.Itoa(i)
		err := addNotes(workingFile,  note)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}
func createBackup(working, backup string) error {
	// check to see if our working file exists
	//before backing it up
	_, err := os.Stat(working)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrWorkingFileNotFound
		}
		return err
	}
	//We need to read the contents of the workFile. We will be using the ioutil.ReadAll
	//method to get all the contents of the workFile. The workFile is of the os.File type,
	//which satisfies the io.Reader interface; this allows us to pass it to ioutil.ReadFile.

	workingFile, err := os.Open(working)
	 if err != nil {
	 	return err
	 }

	content, err := ioutil.ReadAll(workingFile)
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(backup, content, 0644)
	if err != nil {
		fmt.Println(err)
	}
	return nil
}

func addNotes(workingFile, notes string) error {
	notes += "\n"
	f, err := os.OpenFile(workingFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write([]byte(notes))
	err != nil{
		return nil
	}
	return nil
}









