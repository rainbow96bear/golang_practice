package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "http://example.com/search?querykey=queryvalue&stringkey=stringvalue"

	parsedURL, err := url.Parse(urlString)
	if err != nil{
		fmt.Println("error")
		return
	}
	queryParams := parsedURL.Query()
	queryvalue1 := queryParams.Get("querykey")
	queryvalue2 := queryParams.Get("stringkey")

	fmt.Println("querykey의 값 :",queryvalue1)
	fmt.Println("stringkey의 값 : ", queryvalue2)
}