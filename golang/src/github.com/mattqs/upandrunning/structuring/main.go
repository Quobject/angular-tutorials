package main

import "fmt"

func main() {
	// doSomething()

	// sum := addValues(23, 54)
	// fmt.Println("Sum:", sum)

	// sum = addAllValues(12, 54, 79)
	// fmt.Println("Sum:", sum)

	// n1, l1 := FullName("Zaphod", "Beeblebrox")
	// fmt.Printf("Fullname: %v, number of char: %v", n1, l1)

	n1, l1 := FullNameNakedReturn("Zaphod", "Beeblebrox")
	fmt.Printf("Fullname: %v, number of char: %v\n\n", n1, l1)

}

//FullNameNakedReturn a
func FullNameNakedReturn(f, l string) (full string, length int) {
	full = f + " " + l
	length = len(full)
	return
}

//FullName lll
func FullName(f, l string) (string, int) {
	full := f + " " + l
	length := len(full)
	return full, length
}

func addValues(value1, value2 int) int {
	return value1 + value2
}

func addAllValues(values ...int) int {
	sum := 0
	// for i := 0; i < len(values); i++ {
	// 	sum += values[i]
	// }

	for i := range values {
		sum += values[i]
	}

	//fmt.Printf("%T\n", values)
	return sum
}

func doSomething() {
	fmt.Println("Doing something!")
}
