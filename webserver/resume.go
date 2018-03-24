package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resume struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetResumeByID(c *gin.Context) {
	// TODO(nick): implement
	c.AbortWithStatus(http.StatusNotFound)
}
