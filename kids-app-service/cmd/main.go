package main

import (
	"fmt"
	"net/http"

	"../pkg/service"
)

func main() {
	fmt.Println("Hello, world. REST-0")
	err := service.GetQuestions()
	if err != nil {
		fmt.Println(err)
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	mux.HandleFunc("/comment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World comment")
	})

	mux.HandleFunc("/comment/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "return a comment id %s", id)

	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Println(err.Error())
	}
}
