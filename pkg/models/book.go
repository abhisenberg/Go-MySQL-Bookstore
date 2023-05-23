package models

import (
	"fmt"

	"github.com/abhisenberg/go-books/pkg/config"

	"gorm.io/gorm"
)

var db *gorm.DB

/*
The model has 3 levels to it:
- DB representation: For this we declare our struct as GORM model
- Application representation: The Go model defined with `type struct`
- JSON representation: For this we declare json struct tags
*/
type Book struct {
	gorm.Model
	Name   string `json:"name"`
	ISBN   string `json:"isbn"`
	Author string `json:"author"`
}

func init() {
	config.ConnectToDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{}) //This is done to create Book table in the DB
}

/*
The below functions deal only with converting the
"Go objects (application representation)" -> "DB objects"
These functions receive Go objects and use the ORM to store them
into DB.
*/
func (b *Book) CreateBook() *Book {
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var allBooks []Book
	fmt.Println(db)
	db.Find(&allBooks)
	return allBooks
}

func GetBookByID(id int64) *Book {
	var book Book
	db.Where("ID=?", id).Find(&book)
	return &book
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
