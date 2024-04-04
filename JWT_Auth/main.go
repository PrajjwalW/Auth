package main

import (
	"log"
	"myJwtAuth/routes"

	"github.com/gin-gonic/gin"
)

const port = "8080"

func main() {
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	err := router.Run(":" + port)
	if err != nil {
		log.Fatal("Error while starting server: ", err)
	}
}
