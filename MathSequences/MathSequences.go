package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//const one, two = 100, 10
	s := os.Args[0] //not working due to some i dont know yet
	i, err := strconv.Atoi(s)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(s, i)
	d := os.Args[1]
	e, err := strconv.Atoi(d)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	fmt.Println(d, e)
	fmt.Printf("print GCD for 4 and 5, and fibonacci for 5")
	//fmt.Fprint(gcd(one, two))
	fmt.Printf("GCD %d\n", gcd(i, e))
	fmt.Printf("fibonacci %d\n", fib(e))
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
