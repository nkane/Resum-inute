package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	tmpPath = "/tmp/video-data/"
)

type Video struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	FilePath string
}

func GetVideoByID(c *gin.Context) {
	// TODO(nick): implement
	c.AbortWithStatus(http.StatusNotFound)
}

func PostVideo(c *gin.Context) {
	videoFileHeader, err := c.FormFile("video")
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	} else {
		if _, err := os.Stat(tmpPath); os.IsNotExist(err) {
			log.Printf("Directory %s doesn't exist - creating!\n", tmpPath)
			os.Mkdir(tmpPath, os.ModeDir|0744)
		}
		filePath := tmpPath + videoFileHeader.Filename
		err := c.SaveUploadedFile(videoFileHeader, filePath)
		if err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusNotFound)
			return
		}
		log.Println("File saved - path: ", filePath)
		video := Video{
			Name:     videoFileHeader.Filename,
			FilePath: filePath,
		}
		id, err := globalServerDB.CreateVideo(video)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		video.ID = id.String()
		v, err := json.Marshal(video)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return

		}
		c.JSON(http.StatusOK, v)
	}
}
