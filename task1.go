package main

import "fmt"

func PrintSegitiga(rows int) {
	var col int = ((2 * rows) - 1)

	for i := 0; i < rows; i++ {
		for j := 0; j < col; j++ {
			if i > 0 && j < i {
				fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}
		col -= 1
		fmt.Printf("\n")
	}
}
