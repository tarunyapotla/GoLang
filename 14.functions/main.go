package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()

	// func greeterTwo()  { // function inside function is not allowed
	// 	fmt.Println("second function")
	// }
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, msg := proAdder(2, 5, 8, 7, 4)
	fmt.Println("Proresult is: ", proRes)
	fmt.Println("msg is: ", msg)

}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Pro result function"
}

func adder(a int, b int) int {
	return a + b
}

func greeter() {
	fmt.Println("greeter function")
}
func greeterTwo() {
	fmt.Println("second function")
}
