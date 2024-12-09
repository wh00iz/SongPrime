package migrations

import (
	"log"
	model "songPrime/Model"
	utils "songPrime/Utils"
)

func RunMigrations() {
    if utils.DB == nil {
        log.Fatal("Database connection is not initialized")
    }
    utils.DB.AutoMigrate(&model.Song{})
}