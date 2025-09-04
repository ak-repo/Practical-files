package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {

	// handleWebRequest()

	handleURL()

}

// Web request handling : GET
const github = "https://github.com/"

func handleWebRequest() {
	//
	response, err := http.Get(github)

	if err != nil {
		panic(err)
	}

	fmt.Printf("res type:%T \n", response)
	defer response.Body.Close()
	fmt.Println("done")

	//reading
	body, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("body: ", string(body))
}

// url handling

const myUrl = "https://myweb:8080/users?username=ak&email=ak.@gmail.com"

func handleURL() {

	fmt.Println(myUrl)

	//parsing
	result, _ := url.Parse(myUrl)
	fmt.Println("result: ", result.Scheme)
	fmt.Println(result.Host)
	fmt.Println("path:", result.Path)
	fmt.Println("port", result.Port())
	fmt.Println("RawQuery:", result.RawQuery)

	queryParameters := result.Query()
	fmt.Println("params", queryParameters)

	for i, val := range queryParameters {
		fmt.Printf("%v : %v \n", i, val)
	}

	//constructiong usrl
	//https://pkg.go.dev/net/http#Cookie

	partsOfURL := &url.URL{
		Scheme:  "https",
		Host:    "pkg.go.dev",
		Path:    "/net/http",
	}

	fmt.Println("constructed utl: ", partsOfURL.String())

}
