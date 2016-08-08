package main

import (
	"fmt"

	"github.com/mattqs/upandrunning/stringutil"
)

func main() {

	n1, l1 := stringutil.FullName("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, number of char: %v\n\n", n1, l1)

	n2, l2 := stringutil.FullNameNakedReturn("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, number of char: %v\n\n", n2, l2)

}
