package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)
type Person struct{
	Name string `json:"name"`
	Age int	`json:"age"`
}
func main(){
	port := 8080

	http.HandleFunc("/read",JsonReadHandler)
	http.HandleFunc("/write",JsonWriteHandler)

	log.Println("Rainbowbear Server Starting on port :", port)
    log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

func JsonReadHandler(w http.ResponseWriter, r *http.Request){
	var person Person
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&person)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("이름 : %s, 나이 : %d\n", person.Name, person.Age)

	w.Write([]byte("읽기 성공"))
}
func JsonWriteHandler(w http.ResponseWriter, r *http.Request){
	person := Person{
		Name : "RainbowBear",
		Age : 100,
	}

	responseJSON, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}