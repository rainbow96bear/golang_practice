package main

import (
	"fmt"
	"log"
	"net/http"
)
func welcomeFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "welcome")
}

func rainbowFunc(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Rainbow Area")
}

type BearHandler struct {
	name string
}

func (b *BearHandler) ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "bear name is %s", b.name)
}

func main(){
    port := 8080
    mux := http.NewServeMux()

	myBear := &BearHandler{name:"GoodBear"}

	mux.HandleFunc("/",welcomeFunc)
	mux.HandleFunc("/rainbow",rainbowFunc)
	mux.Handle("/bear",myBear)

    log.Println("Rainbowbear Server Starting on port :", port)
    log.Fatal(http.ListenAndServe(fmt.Sprint(":", port), mux))
}
