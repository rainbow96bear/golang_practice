package main

import (
	"fmt"
	"net/http"
)

func main (){
	port := ":8080"

	http.HandleFunc("/", testHandler)

	http.ListenAndServe(port,nil)
}

func testHandler(w http.ResponseWriter, r *http.Request){
	queryParams := r.URL.Query()

	name := queryParams.Get("name")

	if name != ""{
		fmt.Fprintln(w,"Hello,", name)
	}else {
		fmt.Fprintln(w, "Welcome!")
	}
}