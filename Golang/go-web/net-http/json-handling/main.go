package main

import (
	"encoding/json"
	"fmt"
)

//
// encoding -- marshal

type Course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

// error handling
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func EncodingJSON() {

	plainData := []Course{
		{Name: "reactJS", Price: 499, Platform: "online", Password: "12345", Tags: []string{"web", "js"}},
		{Name: "MERN", Price: 99, Platform: "combined", Password: "abcd", Tags: nil},
		{Name: "Golang", Price: 1000, Platform: "offline", Password: "123ef", Tags: []string{"software", "go", "jin"}},
	}

	// data into json
	dataIntoJSON, err := json.Marshal(plainData)
	checkError(err)

	fmt.Printf("%s\n", dataIntoJSON)

	// if want show  JSON way
	dataIntoJSON2, err := json.MarshalIndent(plainData, "", "\t")
	checkError(err)
	fmt.Printf("%s\n", dataIntoJSON2)

}

//
//Decodde -- Unmarshal JSON

func DecodingJSON() {

	jsonData := []byte(`
	{ 
        "coursename": "Golang",
        "price": 1000,
        "website": "offline",
         "tags": [ "software", "go","jin"]
    }
	`)

	// method one
	var jsonIntoData Course
	checkJSON := json.Valid(jsonData)

	if checkJSON {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonData, &jsonIntoData)
		fmt.Printf("%#v\n", jsonIntoData)
	} else {
		fmt.Println("JSON not valid")
	}

	//method two  , data into key value pair
	var jsonIntoData2 map[string]interface{}

	json.Unmarshal(jsonData, &jsonIntoData2)
	// fmt.Printf("%#v\n", jsonIntoData2)

	for k, v := range jsonIntoData2 {

		fmt.Printf("key is %v and value is %v and type is: %T\n", k, v, v)
	}

}

func main() {

	EncodingJSON()

	DecodingJSON()
}
