package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in go")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	// fruitList[2] = "Apple"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is ", fruitList)
	fmt.Println("Fruit list is ", len(fruitList))

	var vegList = [3]string{"Potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList)

}
