package main

import (
	"fmt"
)

func printSegitiga(n int) {
	for i := n; i >= 1; i-- {
		for j := 0; j < n-i; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < i*2-1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

// func main() {
// 	printSegitiga(7)
// }
