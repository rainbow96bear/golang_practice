package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	port := 8080

	http.HandleFunc("/rainbowbearWorld", rainbowbearHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v",port),nil))
}

func rainbowbearHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Welcome to rainbowbear World")
}