package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Listof all langs: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	//delete key value
	delete(languages, "RB")
	fmt.Println(languages)

	//loops in maps
	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n", key, value)
	}
}
