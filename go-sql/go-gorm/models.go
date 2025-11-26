package main

import (
	"fmt"
	"log"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name        string
	Author      string
	Description string
	Price       uint
}

func CreateBook(db *gorm.DB, book *Book) {
  result := db.Create(book)
  if result.Error != nil {
    log.Fatalf("Error creating book: %v", result.Error)
  }
  fmt.Println("Book created successfully")
}

func GetBook(db *gorm.DB, id uint) *Book {
  var book Book
  result := db.First(&book, id)
  if result.Error != nil {
    log.Fatalf("Error finding book: %v", result.Error)
  }
  return &book
}

func UpdateBook(db *gorm.DB, book *Book) {
  result := db.Save(book)
  if result.Error != nil {
    log.Fatalf("Error updating book: %v", result.Error)
  }
  fmt.Println("Book updated successfully")
}

func DeleteBook(db *gorm.DB, id uint) {
  var book Book
  result := db.Delete(&book, id)
  if result.Error != nil {
    log.Fatalf("Error deleting book: %v", result.Error)
  }
  fmt.Println("Book deleted successfully")
}


func getBooksSortedByCreatedAt(db *gorm.DB) ([]Book, error) {
  var books []Book
  result := db.Order("created_at desc").Find(&books)
  if result.Error != nil {
    return nil, result.Error
  }
  return books, nil
}

func getBooksByAuthorName(db *gorm.DB, authorName string) ([]Book, error) {
  var books []Book
  result := db.Where("author = ?", authorName).Find(&books)
  if result.Error != nil {
    return nil, result.Error
  }
  return books, nil
}

func GetBooks(db *gorm.DB) []Book {
  var books []Book
  result := db.Find(&books)

  if result.Error != nil {
    log.Fatalf("Error get book: %v",result.Error)
  }
  return books
}