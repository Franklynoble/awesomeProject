package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
	)

func main() {
	counts := make(map[rune]int)  // counts of unicode characters
	var utflen [utf8.UTFMax + 1]int // count of legnth of UTF-8 encoding
	invalid := 0                    //counts of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes,error
		if err == io.EOF {         // if the error is equal to end of line, exit i.e stop(break)
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		 if unicode.IsLetter(r){
			 counts[r]++
			 utflen[n]++

		 }

		fmt.Printf("rune\tcounts\n")
		for c, n := range counts {
			fmt.Printf("%q\t%d\n", c, n)
		}

		fmt.Print("nlen\tcount\n")
		for i, n := range utflen {
			if i > 0 {
				fmt.Printf("%d\t%d\n", i, n)
			}
		}
		if invalid > 0 {
			fmt.Printf("n%d invalid UTF-8 characters\n", invalid)
		}
	}



}
