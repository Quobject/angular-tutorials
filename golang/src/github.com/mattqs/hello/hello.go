package main

import (
	"fmt"
	"github.com/mattqs/stringutil"
)

type t struct {
	name  string // name of the object
	value int    // its value
}

func main() {
	fmt.Println(stringutil.Reverse("Hello, new gopher!"))
}
