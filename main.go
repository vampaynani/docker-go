package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", 
	os.Getenv("DB_HOST"), 
	os.Getenv("DB_PORT"), 
	os.Getenv("DB_USER"), 
	os.Getenv("DB_PASSWORD"), 
	os.Getenv("DB_NAME"), 
	os.Getenv("DB_SSL"))
	DBConn, err := gorm.Open(postgres.Open(connStr), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal(err)
	}
	r := gin.Default()
	mode := os.Getenv("GIN_MODE")
	fmt.Println(mode)
	if mode == "release" {
		fmt.Println("Load .env.prod file")
	}

	r.GET("/", func(c *gin.Context){
		tx := DBConn.Exec("SELECT 1")
		fmt.Printf("Error: %v\n Rows: %v", tx.Error, tx.RowsAffected)
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	r.Run(os.Getenv("PORT"))
}