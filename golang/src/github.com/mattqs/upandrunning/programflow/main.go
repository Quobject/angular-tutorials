package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//conditional()
	//switch2()
	loops()
}

func loops() {

	sum := 1
	fmt.Println("Sum:", sum)

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	sum = 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	sum = 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum: ", sum)
		if sum > 200 {
			goto endofprogram
		}

		if sum > 500 {
			break
		}
	}

endofprogram:
	fmt.Println("end of program")

}

func switch2() {
	rand.Seed(time.Now().Unix())
	dow := rand.Intn(6) + 1
	fmt.Println("Day", dow)

	result := ""

	switch dow {
	case 1:
		result = "It's Sunday"
	case 7:
		result = "It's a Saturday"
	default:
		result = "It's a weekday"
	}

	fmt.Println("Day", dow, ",", result)

	x := -42
	switch {
	case x < 0:
		result = "Less than zero"
	case x == 0:
		result = "Equals to zero"
	default:
		result = "Greater than zero"
	}

	fmt.Println(result)

}

func conditional() {
	//var x float64 = -42
	var result string

	if x := 42; x < 0 {
		result = "Less than zero"
	} else if x == 0 {
		result = "Equal to zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

}
