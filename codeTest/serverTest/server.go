package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Name string `json:"name"`
}

func main(){
	port := 8080

	r := mux.NewRouter()

	r.HandleFunc("/test",testHandler).Methods("GET")

	http.ListenAndServe(fmt.Sprintf("%d:",port),r)
}

func testHandler(w http.ResponseWriter, r *http.Request){
	var user User
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.Write([]byte(fmt.Sprintf("%s님 반갑습니다.", user.Name)))
}