package main

import (
	"html/template"
	"net/http"
	"strconv"
)

type Person struct {
	ID   int
	NAME string
}

var index = template.Must(template.ParseFiles("index.html"))
var req = template.Must(template.ParseFiles("req.html"))

func request(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, _ := strconv.Atoi(r.URL.Query().Get("id"))
		name := r.URL.Query().Get("name")
		person := Person{
			ID:   id,
			NAME: name,
		}
		req.Execute(w, person)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	dummy := 1
	index.Execute(w, dummy)
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/request", request)
	http.ListenAndServe(":8080", nil)
}
