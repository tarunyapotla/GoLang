package main

import "fmt"

func main() {
	fmt.Println("Loops in golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)

	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days { //comma ok syntax is also allowed
		fmt.Printf("index %v and value is %v \n", index, day)
	}

	n := 1
	for n < 10 {
		if n == 2 {
			goto lco
		}
		if n == 5 {
			break
		}
		fmt.Println("value is ", n)
		n++
	}

lco:
	fmt.Printf("Hello world")
}
