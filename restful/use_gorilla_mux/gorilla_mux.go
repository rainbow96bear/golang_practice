package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type UserInfo struct {
	Name string `json:"name"`
	Age int `json:"age"`
 }
 
 var userInfos []UserInfo = []UserInfo{{Name: "고릴라", Age:100}}

func main() {
	port := 8080
    r := mux.NewRouter()

    r.HandleFunc("/user", GetUsers).Methods("GET")
    r.HandleFunc("/user", PostUsers).Methods("POST")

	log.Println("Rainbowbear Server Starting on port :", port)
    log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), r))
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(userInfos)
}

func PostUsers(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    var userInfo UserInfo
    decoder := json.NewDecoder(r.Body)
    err := decoder.Decode(&userInfo)
    if err != nil {
        http.Error(w, "Bad Request", http.StatusBadRequest)
        return
    }
    userInfos = append(userInfos, userInfo)
    json.NewEncoder(w).Encode(userInfos)
}