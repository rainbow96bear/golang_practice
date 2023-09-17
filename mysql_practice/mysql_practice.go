package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gorilla/pat"
)

type BookInfo struct{
	ID int `json:"id"`
	Kind string `json:"kind"`
	Title string `json:"title"`
	Author string `json:"author"`
}
type responseMsg struct {
    Status int `json:"status"`
    Msg string `json:"msg"`
}

var db *sql.DB

func main(){
	port := 8080
	var err error
	db, err = sql.Open("mysql", "root:1q2w3e4r!@tcp(localhost:3333)/Books")
	if err != nil {
		fmt.Println("DB 열기 실패 :", err)
	}
	defer db.Close()

	router := pat.New()

	router.Get("/Books", GetList)
	router.Post("/Books", CreateData)
	router.Put("/Books", UpdateData)
	router.Delete("/Books", DeleteData)
	fmt.Println(port, "포트 서버 생성")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func CreateData(w http.ResponseWriter, r *http.Request){
	var bookInfo BookInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bookInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	queryParams := r.URL.Query()
	bookInfo.Kind = queryParams.Get("kind")

	query := "INSERT INTO books (kind, title, author) VALUES (?, ?, ?)"
	_, execErr := db.Exec(query, bookInfo.Kind, bookInfo.Title, bookInfo.Author)
	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
		return
	}

	res := responseMsg{Status : http.StatusOK, Msg: "저장 완료"}
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func GetList(w http.ResponseWriter, r *http.Request){
	queryParams := r.URL.Query()
	kind := queryParams.Get("kind")

	query := "SELECT * FROM books WHERE kind = ?"
	rows, queryErr := db.Query(query, kind)
	if queryErr != nil {
		http.Error(w, queryErr.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var data []BookInfo
	for rows.Next() {
		var d BookInfo
		if err := rows.Scan(&d.ID, &d.Kind, &d.Title, &d.Author); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		data = append(data, d)
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func UpdateData(w http.ResponseWriter, r *http.Request){
	var bookInfo BookInfo
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bookInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	
	queryParams := r.URL.Query()
	bookInfo.ID, _ = strconv.Atoi(queryParams.Get("id"))
	bookInfo.Kind = queryParams.Get("kind")

	query := "UPDATE books SET kind = ?, title = ?, author = ? WHERE id = ?"
	_, execErr := db.Exec(query, bookInfo.Kind, bookInfo.Title, bookInfo.Author, bookInfo.ID)
	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
		return
	}

	res := responseMsg{Status : http.StatusOK, Msg: "수정 완료"}
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

func DeleteData(w http.ResponseWriter, r *http.Request){
	queryParams := r.URL.Query()
	id := queryParams.Get("id")

	query := "DELETE FROM books WHERE id = ?"
	_, execErr := db.Exec(query, id)
	if execErr != nil {
		http.Error(w, execErr.Error(), http.StatusInternalServerError)
		return
	}

	res := responseMsg{Status : http.StatusOK, Msg: "삭제 완료"}
	responseJSON, err := json.Marshal(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
}

