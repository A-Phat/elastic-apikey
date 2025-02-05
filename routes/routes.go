package routes

import (
	"elastic-apikey/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/insert", controllers.InsertDocument)
	r.GET("/search", controllers.SearchDocument)

	return r
}
