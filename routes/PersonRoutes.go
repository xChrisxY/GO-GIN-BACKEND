package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-backend/controllers"
)

func PersonRoutes(router *gin.Engine) {

	routes := router.Group("api/v1/persons")
	routes.GET("", controllers.GetPersons)
	routes.POST("", controllers.CreatePerson)
	routes.GET("/:id", controllers.GetPersonById)
	routes.PUT("/:id", controllers.UpdatePerson)
	routes.DELETE("/:id", controllers.DeletePerson)

}