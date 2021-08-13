package main

import "fmt"

/*If all the fields of a str uct are comparable, the str uct its elf is comparable, so two expressions of
that typ e may be compared using == or !=. The == operat ion compares the corresponding
fields of the two str ucts in order, so the two printed expressions below are equivalent:   */


type Point struct{x,y int }

//Comparable str uct typ es, li ke other comparable typ es, may be used as the key typ e of a map.
type address struct{
	hostname string
	port int
}

func main() {

	 m := map[string]int{"James":22,"Join":44}
        for _,gg := range m{
        	 fmt.Print(gg)
       		   }



   hits := make(map[address]int)

    hits[address{"golang.org",443}]++
 fmt.Println(hits)
	p := Point{1,2}
	q := Point{2,1}
	fmt.Println(p.x == p.x && p.y == q.y) // false
	fmt.Println(p == q)  // false
}

