package main

import "fmt"

func main() {
	h := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	//y := []int{5, 8, 2, 4, 0, 1, 3, 7, 9, 6}
	sort1(h)

	// fmt.Println(g )

	fmt.Println()
}

func sort1(s []int) {
	//var nums2 []int
	lenn := len(s)
	for i := 0; i < lenn-1; i++ {

		for j := 0; j < lenn-i-1; j++ {

			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}

		}

	}
	fmt.Println("\nAfter Bubble Sorting")
	for _, val := range s {
		fmt.Println(val)
	}
}
