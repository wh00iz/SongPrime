package utils

import (
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" 
    "log"
    "os"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

   
    DB.DB().SetMaxOpenConns(10) //открытые
    DB.DB().SetMaxIdleConns(10) //пулл

    log.Println("Successfully connected to the database!")
}