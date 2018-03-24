package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var (
	router         *gin.Engine
	globalServerDB *ResuminuteDB
)

func main() {
	var err error
	globalServerDB, err = CreateDBSession()
	if err != nil {
		os.Exit(1)
	}
	defer globalServerDB.dbContext.Close()
	router = gin.Default()
	router.LoadHTMLGlob("../public/templates/*")
	router.Static("/js", "../public/js")
	initRoutes(router)
	router.Run()
}
