package main

import "fmt"

func main() {
	const one, two = 4, 5
	fmt.Printf("print GCD for 4 and 5, and fibonacci for 5")
	//fmt.Fprint(gcd(one, two))
	fmt.Printf("GCD %d\n", gcd(one, two))
	fmt.Printf("fibonacci %d\n", fib(two))
	//fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))

}
func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
