package main

import (
	"fmt"
	"os"
)

// Prints its command line arguments
func main() {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	fmt.Println(os.Args[0])
}
