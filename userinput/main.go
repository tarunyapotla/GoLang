package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza:")

	// comma ok || err ok
	input, _ := reader.ReadString('\n') //reads until \n  (use _ if u dont care about anything like error or input)
	// input, err := reader.ReadString('\n') // if u want to hold or catch errors
	// _, err (dont care abt input only when error is needed)
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of rating is %T", input)

}
