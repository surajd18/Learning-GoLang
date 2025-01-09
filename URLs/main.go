package main

import (
	"fmt"
	"net/url"
)

const urls string = "https://github.com/surajd18"

func main() {
	fmt.Println(urls);

	result,_ := url.Parse(urls);
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
}