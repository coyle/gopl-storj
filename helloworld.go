package main

import (
	"fmt"
)

type foo struct {
	s string
	i int
	f bool
}

//
var s, i string
var bar int

func main() {
	fmt.Println("Hello 世界")

	b, _ := foobar("hello", "world", 5)
	fmt.Println(b)

}

func foobar(s, i string, f int) (bool, int) {
	return false, 0
}
