package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type UserInfo struct {
   Name string `json:"name"`
   Age int `json:"age"`
}

var userInfos []UserInfo = []UserInfo{{Name: "무지개곰", Age:100}}

func main() {
	port := 8080

	http.HandleFunc("/user",HandleUsers)

	log.Println("Rainbowbear Server Starting on port :", port)
    log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), nil))
}

func HandleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        GetUsers(w, r)
	case http.MethodPost:
        PostUsers(w, r)
    }
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