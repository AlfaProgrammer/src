package main

import (
	"fmt"
	"os"
)

type count int

// signature per fare parte del type(interface) stringer
func (c count) String() string {
	return "test"
}

func main() {
	var c count = 235
	os.Exit(34)
	fmt.Println(c) //test
}
