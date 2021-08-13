package main

import (
	"errors"
	"fmt"
	"os"
)

var m string

type locales struct {
	Country  []string
	Language []string
	//France,Francea string
	//Russia, Russian string
	//Canada ,Canadian string

}

var nn = locales{
	Country:  []string{"France", "Russia", "Canada"},
	Language: []string{"en_US", "en_CN", "fr_CN", "fr_FR", "ru_RU"},
}

var errov = errors.New("Errors")


func main() {

	lc := "en_GB"
	locales1 := struct {
		allKeys map[string]string
	}{
		map[string]string{"en_US": nn.Language[0], "en_CN": nn.Language[1], "fr_CN": nn.Language[2], "fr_FR": nn.Language[3], "ru_RU": nn.Language[4]},
	}
	//	mm := map[string]string{"en_US": nn.Language[0], "en_CN": nn.Language[1], "fr_CN": nn.Language[2], "fr_FR": nn.Language[3], "ru_RU": nn.Language[4]}
	//var n string
	if k, v := locales1.allKeys[lc]; v {
		fmt.Println(k)
	} else {

		fmt.Printf("%s\n", locales1.allKeys[k])
		//	os.Exit(1)
	}
}
	//nmd := map[string]string{"en_US": nn.Language[0], "en_CN": nn.Language[1], "fr_CN": nn.Language[2], "fr_FR": nn.Language[3], "ru_RU": nn.Language[4]}
	//
	//showLocale(nn.Language[3], nmd)
	//
	//	allLocales := struct {
	//		Names  []string;
	//		//Region string;
	//	}{
	//		[]string{nn.language[0], nn.language[1],nn.language[2],nn.language[3]},
	//
	//	}
	//	   i := 0
	//	if allLocales.Names >
	//}


func PassedArgs() []string {
	var args []string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return args
}

//	if f, v := nm["fr_CN"]; v {
//		fmt.Println(f)
//	} else {
//       fmt.Println("Error",f)
//	}
//}

func showLocale(i string, w map[string]string) {
	//n = map[string]string{}
	if k, v := w[i]; v {

		fmt.Println(k)
	} else {
		fmt.Println("Errors")
	}

}

//
//getLocale("ru_RU")
//
//func getLocalec(n string) string {
//	language := locales{
//		America: "en_US",
//		 Enlgish: ""
//		France:  "fr_CN",
//		Russia:  "ru_RU",
//		Canada:  "en_CN",
//	}
//
//	if n == language.Russia {
//		n = language.Russia
//		return n
//	} else if n == language.America {
//		n = language.America
//		//	fmt.Printf("%s :", n)
//	} else if n == language.France {
//		n = language.France
//		//	fmt.Printf("%s :", n)
//	} else if n == language.Canada {
//		n = language.Canada
//		//fmt.Printf("%s:", n)
//	} else {
//		fmt.Println("Local not Supported: ", n)
//	}
//	fmt.Sprintf("%s", n)
//
//}
