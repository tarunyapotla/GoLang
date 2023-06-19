package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	//no inheritance, no super or parent in golang
	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("hitesh details are: %v\n", hitesh)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("Name is %v and email is %v\n", hitesh.Name, hitesh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is: ", u.Email)
}
