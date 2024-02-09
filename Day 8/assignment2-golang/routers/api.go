package routers

import (
	"assignment2-golang/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.GET("/orders", controllers.GetAllOrders)
	router.GET("/orders/:id", controllers.GetOrder)
	router.PUT("/orders/:id", controllers.UpdateOrder)
	router.DELETE("/orders/:id", controllers.DeleteOrder)

	router.POST("/items", controllers.CreateItem)
	router.GET("/items", controllers.GetAllItems)
	router.GET("/items/:id", controllers.GetItem)
	router.PUT("/items/:id", controllers.UpdateItem)
	router.DELETE("/items/:id", controllers.DeleteItem)

	return router
}
