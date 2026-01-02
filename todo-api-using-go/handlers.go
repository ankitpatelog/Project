package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func todosHandler(w http.ResponseWriter,r *http.Request){
	switch r.Method{
	case http.MethodGet:
		json.NewEncoder(w).Encode(todos)

	case http.MethodPost:
		var t Todo
		err := json.NewDecoder(r.Body).Decode(&t)
		if err!=nil{
			http.Error(w,"invalied json",http.StatusBadRequest)
			return
		}

		t.ID=nextID
		nextID++
		todos = append(todos,t)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)

	default:
		http.Error(w,"Method not allowed",http.StatusMethodNotAllowed)
	}
	
} 

func todoHandler(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, t := range todos {
		if t.ID == id {

			switch r.Method {

			case http.MethodGet:
				json.NewEncoder(w).Encode(t)

			case http.MethodPut:
				var updated Todo
				json.NewDecoder(r.Body).Decode(&updated)
				todos[i].Title = updated.Title
				todos[i].Completed = updated.Completed
				json.NewEncoder(w).Encode(todos[i])

			case http.MethodDelete:
				todos = append(todos[:i], todos[i+1:]...)
				w.WriteHeader(http.StatusNoContent)

			default:
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			}
			return
		}
	}

	http.Error(w, "Todo Not Found", http.StatusNotFound)
}
