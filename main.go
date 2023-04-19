package main

import (
	"github.com/gin-gonic/gin"
	"sharecode/controllers"
	"sharecode/utils"
)

func main() {
	// loading environment varibles
	utils.LoadEnv()
	utils.PayloadUpgrader()
	// initialize a gin router.
	router := gin.Default()

	// routes
	router.GET("/", controllers.MainRoute)
	router.GET("/s/:id", controllers.NewSnippet)

	// run the server
	router.Run()

}
