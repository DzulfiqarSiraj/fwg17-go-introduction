package main

import "fmt"

func GenPass(password string) {
	if len(password) < 6 {
		fmt.Println("password length must be inputed at least 6 characters")
	} else {
		fmt.Println(password)
	}
}
