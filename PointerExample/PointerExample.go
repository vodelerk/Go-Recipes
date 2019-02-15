// simple pointer example
package main

import "fmt"

func main() {
	//simplepointer()
	//fmt.Println(f() == f()) // "false"
	v := 1
	incr(&v)              // side effect: v is now 2
	fmt.Println(incr(&v)) // "3" (and v is 3)
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
func alternativepointer() {

}
func incr(p *int) int {
	*p++ // increments what p points to; does not change p
	return *p
}

var p = f()

func f() *int {
	v := 1
	return &v
}
