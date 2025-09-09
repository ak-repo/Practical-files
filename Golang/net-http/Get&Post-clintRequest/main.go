package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

// error handling

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// post
func handlePostRequest() {
	const myURL = "http://localhost:8000/post"

	// fake json data
	var requestData = strings.NewReader(`{ "name":"ak", "age":"23" }`)

	response, err := http.Post(myURL, "application/json", requestData)
	checkError(err)

	// close
	defer response.Body.Close()

	//read method one
	content, err := io.ReadAll(response.Body)
	checkError(err)

	var strResponse strings.Builder
	strResponse.Write(content)

	fmt.Println("Post data response: ", strResponse.String())
}

// get

func handleGetRequest() {
	const myURL = "http://localhost:8000/get"

	response, err := http.Get(myURL)
	checkError(err)

	defer response.Body.Close()

	//read method two
	context, err := io.ReadAll(response.Body)
	checkError(err)
	fmt.Println("Get request response: ", string(context))

}

func main() {

	handlePostRequest()
	handleGetRequest()
}
