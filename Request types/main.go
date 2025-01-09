package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerformPostRequest()

}

func PerformGetRequest() {
	const url = "http://localhost:8000/get";

	response,err := http.Get(url);

	if err != nil {
		panic(err)
	}

	defer response.Body.Close();

	content,_ := io.ReadAll(response.Body);

	fmt.Println(string(content))

}

func PerformPostRequest(){
	const url = "http://localhost:8000/post";

	requestBody := strings.NewReader(`
		{
			"coursename":"Golang",
			"Price":0
		}
	`)

	response,err := http.Post(url,"application/json",requestBody)

	if err != nil{
		panic(err)
	}


	defer response.Body.Close();

	content,_ := io.ReadAll(response.Body);

	fmt.Println(string(content));
}