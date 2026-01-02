package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Todo-api")

	mux := http.NewServeMux()

	mux.HandleFunc("/todos",todosHandler)
	mux.HandleFunc("/todos/", todoHandler)
	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", mux)


}