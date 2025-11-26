package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"
	port     = 3306
	user     = "myuser"
	password = "mypassword"
	dbname   = "mydb"
)

func main() {
	// DSN format ของ MySQL
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbname,
	)

	// New logger for detailed SQL logging
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect to database: " + err.Error())
	}

	// migrate schema
	db.AutoMigrate(&Book{})
	fmt.Println("MySQL Connected & Migration completed!")

	// //add
	// newBook := Book{Name: "momo", Author: "jajarr", Description: "test", Price: 200}
	// CreateBook(db, &newBook)

	// // // Get a book
	// book := GetBook(db, 1) // Assuming the ID of the book is 1
	// // fmt.Println("Book Retrieved:", book)

	// // // Update a book
	// book.Name = "The Go Programming Language, Updated Edition"
	// UpdateBook(db, book)

	// // Delete a book
	// DeleteBook(db, book.ID)

	// Setup Fiber
	app := fiber.New()

	// CRUD routes
	app.Get("/books", func(c *fiber.Ctx) error {
		return getBooks(db, c)
	})
	// app.Get("/books/:id", func(c *fiber.Ctx) error {
	// 	return getBook(db, c)
	// })
	// app.Post("/books", func(c *fiber.Ctx) error {
	// 	return createBook(db, c)
	// })
	// app.Put("/books/:id", func(c *fiber.Ctx) error {
	// 	return updateBook(db, c)
	// })
	// app.Delete("/books/:id", func(c *fiber.Ctx) error {
	// 	return deleteBook(db, c)
	// })

	// Start server
	log.Fatal(app.Listen(":8000"))
}

func getBooks(db *gorm.DB,c *fiber.Ctx) error {
	return c.JSON(GetBooks(db))
}

// func getBookById(c *fiber.Ctx) error {
// 	bookId, err := strconv.Atoi(c.Params("id"))

// 	//cant convert
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	for _, book := range books {
// 		if book.ID == bookId {
// 			return c.JSON(book)
// 		}
// 	}
// 	return c.SendStatus(fiber.StatusNotFound)
// }

// func createBook(c *fiber.Ctx) error {
// 	book := new(Book)
// 	if err := c.BodyParser(book); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}
// 	books = append(books, *book)
// 	return c.JSON(book)
// }

// func updateBook(c *fiber.Ctx) error {
// 	bookId, err := strconv.Atoi(c.Params("id"))

// 	//cant convert
// 	if err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}

// 	bookUpdate := new(Book)
// 	if err := c.BodyParser(bookUpdate); err != nil {
// 		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
// 	}
	
// 	for i, book := range books {
// 		if book.ID == bookId {
// 			books[i].Title = bookUpdate.Title
// 			books[i].Author = bookUpdate.Author
// 			return c.JSON(books[i])
// 		}
// 	}
// 	return c.SendStatus(fiber.StatusNotFound)
// }

// func deleteBook(c *fiber.Ctx) error {
// 	id, err := strconv.Atoi(c.Params("id"))
// 	if err != nil {
// 		return c.SendStatus(fiber.StatusBadRequest)
// 	}

// 	for i, book := range books {
// 		if book.ID == id {
// 			books = append(books[:i], books[i+1:]...)
// 			return c.SendStatus(fiber.StatusNoContent)
// 		}
// 	}

// 	return c.SendStatus(fiber.StatusNotFound)
// }