package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              // "-" remove or hides the field from json
	Tags     []string `json:"tags,omitempty"` //remove or omits empty value field
}

func main() {
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "abc1234", []string{"fullstack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "abc1235", nil},
	}

	//package this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //Marshal or MarshalIndent
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": [
				"web-dev",
				"js"
		]
    }
	`)

	var lcoCourse course
	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Printf("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Printf("JSON IS NOT VALID")
	}

	//add data to key value
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}

}
