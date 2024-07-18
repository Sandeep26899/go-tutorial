package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"name"`
	Price    int      `json:"price"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{
			"ReactJs Bootcamp",
			299,
			"example.com",
			"abc123",
			[]string{"web-dev", "js"},
		},
		{
			"NodeJs Bootcamp",
			499,
			"example.com",
			"abc1234",
			[]string{"backend", "js"},
		},
		{
			"NestJs Bootcamp",
			599,
			"example.com",
			"abc1235",
			nil,
		},
	}

	//package this data as json data

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"name": "ReactJs Bootcamp",
		"price": 299,
		"platform": "example.com",
		"tags": [
				"web-dev",
				"js"
		]
     }
	`)

	var myCourses course

	isValid := json.Valid(jsonDataFromWeb)

	if isValid {
		fmt.Println("Valid Json")
		json.Unmarshal(jsonDataFromWeb, &myCourses)
		fmt.Printf("%#v\n", myCourses)
	} else {
		fmt.Println("Json was not valid!")
	}

	//some cases where you just want to add data to key value

	var onlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &onlineData)
	fmt.Printf("%#v\n", onlineData)

	for k, v := range onlineData {
		fmt.Printf("Key is: %v and value is: %v and type is: %T\n", k, v, v)
	}

}
