package controllers

import (
	"bookstore-sql/pkg/config"
	"bookstore-sql/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"

	"bookstore/config"
	"bookstore/models"

	"github.com/gorilla/mux"
	"golang.org/x/tools/go/analysis/passes/defers"
)

func CreateBook(w http.ResponseWriter, r *http.Request)  {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	query := "INSERT INTO books (title, author, price) VALUES (?, ?, ?)"

	result ,err:= config.DB.Exec(query,book.Title,book.Author,book.Price)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	id,_ := result.LastInsertId()

	book.ID=int(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

//get all books
func GetBooks(w http.ResponseWriter, r *http.Request)  {
	var books[] models.Book

	query:= "SELECT id, title, author, price FROM books"
	rows,err:= config.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	//ab har rows pe jake sabi ke detail ko books slice mai add karte jayenge
	for rows.Next(){
		var book models.Book
		rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price)
		books = append(books, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func GetBookByID(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	bookid := vars["id"]

	var book models.Book

	id,_ :=strconv.Atoi(bookid)//return  id in int

	query := "SELECT id, title, author, price FROM books WHERE id = ?"
	
	err := config.DB.QueryRow(query,id).Scan(&book.ID, &book.Title, &book.Author, &book.Price)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var updatedBook models.Book
	json.NewDecoder(r.Body).Decode(&updatedBook)

	query := "UPDATE books SET title=?, author=?, price=? WHERE id=?"
	result, err := config.DB.Exec(query,
		updatedBook.Title,
		updatedBook.Author,
		updatedBook.Price,
		id,
	)

	rows, _ := result.RowsAffected()
	if rows == 0 {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	updatedBook.ID = id
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}
