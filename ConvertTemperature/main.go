package main

import "fmt"

func main() {
	//fmt.Print(" ")
	fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
	boilingF := CToF(BoilingC)
	fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
	fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
}
