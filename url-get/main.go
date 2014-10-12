package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Person struct {
	ID   int    `json:"id"`
	NAME string `json:"name"`
}

var t = template.Must(template.ParseFiles("index.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			log.Fatal(err)
		}
		name := r.URL.Query().Get("name")
		person := Person{
			ID:   id,
			NAME: name,
		}
		t.Execute(w, person)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
