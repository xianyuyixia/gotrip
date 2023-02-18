package main

import "fmt"

func main() {
	a := 1

	if a == 1 {
		fmt.Printf("a=1\n")
	} else {
		fmt.Printf("a!=1\n")
	}

	if a == 2 {
		fmt.Printf("a=2\n")
	} else if a == 1 {
		fmt.Printf("a=1\n")
	} else {
		fmt.Printf("a!=1 而且 a!=2\n")
	}

	if b := a == 1; b {
		fmt.Printf("a=1\n")
		fmt.Printf("b=%t\n", b)
	} else {
		fmt.Printf("b!=1\n")
	}
}
