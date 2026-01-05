package models

import "bokk-management-sql/pkg/config"
import "github.com/jinzhu/gorm"

var db *gorm.DB

// Book model → maps to `books` table
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// init runs automatically when package loads
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//
// CREATE
//
//its a method of the struct Book
func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(b).Error; err != nil {
		return nil, err
	}
	return b, nil
}

//
// READ ALL
//
func GetAllBooks() ([]Book, error) {
	var books []Book

	if err := db.Find(&books).Error; err != nil {
		return nil, err
	}

	return books, nil
}

//
// READ ONE (by ID)
//
func GetBookById(id int64) (*Book, error) {
	var book Book

	result := db.First(&book,id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &book, nil
}

//
// DELETE
//
func DeleteBook(id int64) (*Book, error) {
	var book Book

	// first check if book exists
	if err := db.First(&book, id).Error; err != nil {
		return nil, err
	}

	// then delete
	if err := db.Delete(&book).Error; err != nil {
		return nil, err
	}

	return &book, nil
}

func UpdateBook(book *Book) error {
	return db.Save(book).Error
}
