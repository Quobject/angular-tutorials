package main

import (
	"fmt"
	"sort"
)

func main() {
	//pointers()
	//arrays()
	//slices()
	//memory()
	structures()
}

type Doc struct {
	Breed  string
	Weight int
}

func structures() {
	poodle := Doc{"Poodle", 34}

	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v\nWeight: %v", poodle.Breed, poodle.Weight)
}

func memory() {
	m := make(map[string]int)
	m["key"] = 42
	fmt.Println(m)

	states := make(map[string]string)

	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println(states)

	california := states["CA"]
	fmt.Println(california)

	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	for k, v := range states {
		fmt.Printf("%v: %v\n", k, v)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}

	sort.Strings(keys)

	fmt.Println(keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}

}

func slices() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5, 5)
	numbers[0] = 1
	numbers[1] = 7
	numbers[2] = 32
	numbers[3] = 12
	numbers[4] = 156

	fmt.Println(numbers)

	numbers = append(numbers, 235)
	fmt.Println(numbers)
	fmt.Println(cap(numbers))

	sort.Ints(numbers)
	fmt.Println(numbers)
}

func arrays() {
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Green"
	colors[2] = "Blue"
	fmt.Println(colors)
	fmt.Println(colors[1])

	var numbers = [5]int{5, 3, 1, 2, 4}
	fmt.Println(numbers)
	fmt.Println("Number of colors:", len(colors))
	fmt.Println("Number of numbers:", len(numbers))
}

func pointers() {
	var p *int

	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	var v = 42
	p = &v

	if p != nil {
		fmt.Println("Value of p:", *p)
	} else {
		fmt.Println("p is nil")
	}

	value1 := 42.13
	pointer1 := &value1
	fmt.Println("Value 1:", *pointer1)

	*pointer1 = *pointer1 / 31
	fmt.Println("Value 1:", *pointer1)
	fmt.Println("Value 1:", value1)

}
