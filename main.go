package main

import (
	handlers "songPrime/Handlers"
	migrations "songPrime/Migrations"
	utils "songPrime/Utils"
	"songPrime/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	utils.InitLogger()
	utils.InitDB()
	migrations.RunMigrations()

	r := gin.Default()

	docs.SetupSwagger(r)

	r.GET("/songs", handlers.GetSongs)
	r.GET("/songs/:id/text", handlers.GetSongText)
	r.DELETE("/songs/:id", handlers.DeleteSong)
	r.PUT("/songs/:id", handlers.UpdateSong)
	r.POST("/songs", handlers.CreateSong)

	r.Run(":8080")
}
