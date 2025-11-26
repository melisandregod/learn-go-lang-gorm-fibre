package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
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
	db.AutoMigrate(&User{})
	fmt.Println("MySQL Connected & Migration completed!")

	// Setup Fiber
	app := fiber.New()

	app.Post("/register", func(c *fiber.Ctx) error {
		return createUser(db, c)
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		return loginUser(db, c)
	})

	app.Use("/books", authRequired)
	app.Get("/books", func(c *fiber.Ctx) error {
		return getBooks(db, c)
	})
	// Start server
	log.Fatal(app.Listen(":8000"))
}

func getBooks(db *gorm.DB, c *fiber.Ctx) error {
	return c.JSON(GetBooks(db))
}

func authRequired(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	jwtSecretKey := "test"
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})

	if err != nil || !token.Valid {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.Next()
}
