// Echo1 prints its command-line arguments.
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func basicEcho() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		fmt.Println(i)
		fmt.Println(s)
		sep = " "
	}
	//fmt.Println(s)
}
