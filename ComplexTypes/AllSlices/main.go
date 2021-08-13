package main

import (
	"fmt"
	"os"
)


func main() {

	   fg := []string{"eu_Au", "me_ME"}
    locales := getLocales(fg)
    	fmt.Println("locales to use :",locales)

}



func getPassedArgs()[]string {
	var args [] string
	for i := 1; i < len(os.Args); i ++ {
		args = append(args,os.Args[i])
	}
	return args
}
func getLocales(extraLocales []string)[]string {
	var locales []string

	locales = append(locales,"en_US","fr_FR")
	locales = append(locales,extraLocales...)
	return locales
}