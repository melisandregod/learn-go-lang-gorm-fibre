package main

import (
	"fmt"
	"jajar-test/model"      // model
	"jajar-test/repository" // repo
	"jajar-test/service"    // service
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//load env
	err := godotenv.Load()
	if err != nil {
		log.Println("env not found")
	}

	//DB
	db, err := connectDB()
	if err != nil {
		log.Fatal("Database connection error:", err)
	}


	if err := db.AutoMigrate(&model.Transaction{}); err != nil {
		log.Fatal("Migration error:", err)
	}

	
	txRepo := repository.NewTransactionRepository(db)
	txService := service.NewTransactionService(txRepo)

	
	app := fiber.New()

	
	SetupRoutes(app, txService)

	port := getEnv("PORT", "3000") // default = 3000
	fmt.Println("Server running on port " + port)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal("Server start error:", err)
	}
}


func connectDB() (*gorm.DB, error) {

	user := getEnv("DB_USER", "")
	pass := getEnv("DB_PASS", "")
	host := getEnv("DB_HOST", "")
	port := getEnv("DB_PORT", "")
	name := getEnv("DB_NAME", "")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name,
	)

	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func getEnv(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
