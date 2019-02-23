package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Print(" ")
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	//medals example

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}

	//formatting examples
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	xc := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", xc)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF

	//unicode example
	ascii := 'a'
	unicode := 'D'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 D 'D'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"

	//print exponent
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d eA = %8.3f\n", x, math.Exp(float64(x)))
	}

}
