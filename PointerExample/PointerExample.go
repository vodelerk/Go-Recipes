// simple pointer example
package main

import "fmt"

func main() {
	simplepointer()
}
func simplepointer() {
	x := 1
	p := &x                                    // p, of type *int, points to x
	fmt.Println(*p)                            // "1"
	*p = 2                                     // equivalent to x = 2
	fmt.Println(x)                             // "2"
	var z, y int                               // pointer comparison
	fmt.Println(&z == &z, &z == &y, &z == nil) //true false false
}
