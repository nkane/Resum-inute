package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	jsonRequest = "application/json"
	xmlRequest  = "application/xml"
)

func render(c *gin.Context, data gin.H, template string) {
	switch c.Request.Header.Get("Accept") {
	case jsonRequest:
		c.JSON(http.StatusOK, data["payload"])
	case xmlRequest:
		c.XML(http.StatusOK, data["payload"])
	default:
		c.HTML(http.StatusOK, template, data)
	}
}
