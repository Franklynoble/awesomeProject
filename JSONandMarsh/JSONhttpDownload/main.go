package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	fileName := "SamplDwonload.DOC"

	URL := "https://api.github.com/search/issues?q=1"

	err := downloadFile(URL, fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("file%s download in current working directory", fileName)

}

func downloadFile(URL, fileName string) error {
	//get the response byte fom the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return errors.New("Received non 200 Response Code")

	}

	// create a  empty file
	file,err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the Bytes to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
