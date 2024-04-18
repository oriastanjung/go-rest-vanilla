package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	fmt.Println("Rest API in GO 1.22")

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	mux.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "get by id : "+id)
	})
	mux.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {
		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fmt.Fprintf(w, "POST Person : %+v", person)
	})

	if err := http.ListenAndServe("localhost:8000", mux); err != nil {
		fmt.Println(err.Error())
	}
}
