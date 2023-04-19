package main

import (
	"net/http"
	"sharecode/controllers"
	"sharecode/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	// loading environment varibles
	utils.LoadEnv()
	utils.PayloadUpgrader()

	// initialize a gin router.
	router := gin.Default()

	// static files setup
	// fs := http.FileServer(http.Dir("assets"))
	// http.Handle("/st/", http.StripPrefix("/st/", fs))

	router.GET("/css", func(c *gin.Context) {
		http.ServeFile(c.Writer ,c.Request, "assets/style.css" )
	})
	router.GET("/ico", func(c *gin.Context) {
		http.ServeFile(c.Writer ,c.Request, "assets/favicon.svg" )
	})
	// routes
	router.GET("/", controllers.MainRoute)
	router.GET("/s/:id", controllers.NewSnippet)

	// run the server
	router.Run()

}
