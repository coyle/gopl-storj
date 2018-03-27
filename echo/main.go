package main

import (
	"flag"
	"fmt"
	"strings"
)

// creates a new variable of type bool
var n = flag.Bool("n", false, "omit trailing newline") // pointer to flag variable -accessed *n
var sep = flag.String("s", " ", "seperator")           // pointer to flag variable - accessed *sep

func main() {
	flag.Parse()
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}

}
