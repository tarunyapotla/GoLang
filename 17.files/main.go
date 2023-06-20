package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	content := "This needs to go in a file"

	file, err := os.Create("./gofile.txt")
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./gofile.txt")

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("file data is \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
