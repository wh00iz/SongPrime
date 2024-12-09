package Handlers

import (
	"net/http"
	model "songPrime/Model"
	utils "songPrime/Utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetSongs(c *gin.Context) {
	var song []model.Song
	utils.DB.Find(&song)
	c.JSON(http.StatusOK, song)
}

func GetSongText(c *gin.Context) {
    id, _ := strconv.Atoi(c.Param("id"))
    var song model.Song
    utils.DB.First(&song, id)
    c.JSON(http.StatusOK, song.Text)
}

func DeleteSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var song model.Song
	utils.DB.Delete(&song, id)
	c.Status(http.StatusNoContent)
}

func UpdateSong(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var song model.Song
	utils.DB.First(&song, id)
	c.BindJSON(&song)
	utils.DB.Save(&song)
	c.JSON(http.StatusOK, song)
}

func CreateSong(c *gin.Context) {
	var song model.Song
	c.BindJSON((&song))
	utils.DB.Create(&song)
	c.JSON(http.StatusCreated, song)
}
