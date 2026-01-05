package controllers

import (
	"bokk-management-sql/pkg/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/bokk-management-sql/pkg/models"
	"github.com/bokk-management-sql/pkg/utils"
	"github.com/gorilla/mux"
)

//
// GET ALL BOOKS
//
func GetBook(w http.ResponseWriter, r *http.Request)  {
	books,error := models.GetAllBooks()
	if error!=nil {
		http.Error(w,error.Error(),http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	id, err := strconv.ParseInt(bookId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book,err := models.GetBookById(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, r *http.Request)  {
	//sabse pehle client data in json frmat go struct mai covert hoga
	//then usko book struct mai save karenge then uska hi mehtod createbook ko call karnge
	//then vreatebook struct mai saved data ko se karke db mai data book save karega

	book := &models.Book{}
	err :=json.NewDecoder(r.Body).Decode(&book)
	if err!=nil {
		http.Error(w,"Json not converted",http.StatusInternalServerError)
		return
	}
	//ab tak json -->go conversion done and stored in book struct

	createdbook,error := book.CreateBook()

	if error != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create book"))
		return
	}

	// Step 4: response bhejo
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdbook)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	//data aate hi client se struct mai fill kar denge json to go struct banake 
	//phir getbookbyid se book nikal lo jo model mai alreafdy functio bana hua hai wo book return  karega hoga to us id pe 
	//warna error de dega booknahi milega to
	//phir jo data struct mai hoga update karne ke liye then use replace kar denge book ke data se
	//Book{
//   Name: "",
//   Author: "New Author",
//   Publication: "",   liek this isme client se sirf update ke liye author hai steuct mai itna hi data hai
//	 // ab struct wale data ko book dat mai save kard enge and model mai updatebook ko call kar denge jo ki book lega as argument and then then db.save kar dega

	updatedData := &models.Book{}
	err := json.NewDecoder(r.Body).Decode(&updatedData)
	if err!=nil {
		http.Error(w,"error in json to go struct",http.StatusInternalServerError)
		return
	}

	//id nikalo 
	vars := mux.Vars(r)
	bookid := vars["bookId"]

	id,err := strconv.ParseInt(bookid,10,64)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}

	book,err := models.GetBookById(id)
	if err != nil {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	//check which data client want to change

	if updatedData.Author !="" {
		//means client want  to chnge the author name
		book.Author = updatedData.Author
	}

	if updatedData.Name !="" {
		book.Name=updatedData.Name
	}

	if updatedData.Publication != "" {
		book.Publication = updatedData.Publication
	}

	//now book contains all the updated value

	err := models.UpdateBook(book)

	





}


