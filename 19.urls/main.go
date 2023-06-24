package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs"

func main() {

	//parsing url
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}

	fmt.Println(result.RawQuery)

	qparams := result.Query() //type - url.Values
	fmt.Println(qparams, qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}

	//construct url
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/learn",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
