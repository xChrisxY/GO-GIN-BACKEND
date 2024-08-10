package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-backend/configs"
	"github.com/go-backend/routes"
)

func init() {

	configs.ConnectToDB()

}

func main() {

	r := gin.Default()

	routes.PersonRoutes(r)

	r.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"message":"Hello world from server GO"})

	})

	r.Run()

}