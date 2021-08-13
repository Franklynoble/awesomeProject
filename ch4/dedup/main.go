package main

import (
	"bufio"
	"fmt"
	"os"
)
var  m = make(map[string]int)

func main() {
	ss :=   []string{"peter","James","John","Abby"}
	cs := k(ss)
	fmt.Println(cs)
	Add(ss)
	 adc := count(ss)
	 fmt.Println(adc)

}

func dedupProgram(){
	/*
	Go does not provide a set type, but since the keys of a map are distinct, a map can ser ve this
	purpos e. To illustrate, the program dedup reads a sequence of lines and prints only the first
	occurrence of each distinct line. (It’s a var iant of the dup prog ram that we showed in
	Section 1.3.) The dedup prog ram uses a map whose keys represent the set of lines that have
	already appeared to ensure that subsequent occ urrences are not printed.
	 */
		seen := make(map[string]bool) // a set of strings
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {  // receive parameters from user
			line := input.Text()
			if !seen[line] {
				seen[line] = true
				fmt.Println(line)
			}

		}
		if err := input.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "dedup: %v\n",err)
			os.Exit(1)
		}
	}

	func k(list []string) string{
		return fmt.Sprintf("%q",list)
	}
	func Add(list[]string){
		m[k(list)] ++

	}

	func count(list[]string) int{
		return m[k(list)]
	}
