package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}

}

func handlePostFormRequest() {

	const myYRL = "http://localhost:8000/postform"

	//formdata
	data := url.Values{}

	data.Add("Name", "Ak")
	data.Add("age", "23")
	data.Add("Email", "ak@gmail.com")

	response, err := http.PostForm(myYRL, data)
	checkError(err)

	//close the request
	defer response.Body.Close()

	//read
	context, err := io.ReadAll(response.Body)
	checkError(err)

	var strResponse strings.Builder

	strResponse.Write(context)
	fmt.Println("PostForm response: ", strResponse.String())

}

func main() {

	handlePostFormRequest()
}
