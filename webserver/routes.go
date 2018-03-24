package main

import "github.com/gin-gonic/gin"

func initRoutes(router *gin.Engine) {
	router.GET("/", indexPage)

	resumeRouter := router.Group("resume")
	{
		resumeRouter.GET(":id", GetResumeByID)
	}

	videoRouter := router.Group("video")
	{
		videoRouter.POST("", PostVideo)
	}

}

func indexPage(c *gin.Context) {
	render(c, gin.H{"title": "Resum-inute"}, "index.html")
}
