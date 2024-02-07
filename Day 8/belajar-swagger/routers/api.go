package routers

import (
	"belajar-swagger/controllers"
	_ "belajar-swagger/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Car API
// @version 1.0
// @description This is a sample service for managing cars
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email soberkoder@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0html
// @host localhost:8080
// @BasePath /
func StartServer() *gin.Engine {
	router := gin.Default()

	// Read
	router.GET("/cars/:id", controllers.GetOneCar)
	// Create
	router.POST("/cars", controllers.CreateCars)
	// Read All
	router.GET("/cars", controllers.GetAllCars)
	// Update
	router.PATCH("/cars/:id", controllers.UpdateCars)
	// Delete
	router.DELETE("/cars/:id", controllers.DeleteCars)

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
