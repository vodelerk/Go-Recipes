package main

import "fmt"

func main() {
	x := "hello!"
	for i := 0; i < len(x); i++ {
		x := x[i]
		if x != '!' {
			x := x + 'A' - 'a'  //depending on the go versions can vary
			fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
		}
	}
}
