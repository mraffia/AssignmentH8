package controllers

import (
	"assignment2-golang/database"
	"assignment2-golang/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllOrders(c *gin.Context) {
	var db = database.GetDB()

	var orders []models.Order
	err := db.Find(&orders).Error

	if err != nil {
		fmt.Println("Error getting order datas: ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func GetOrder(c *gin.Context) {
	var db = database.GetDB()

	var order models.Order
	err := db.First(&order, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func CreateOrder(c *gin.Context) {
	var db = database.GetDB()
	var input models.Order

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderInput := models.Order{CustomerName: input.CustomerName, OrderedAt: input.OrderedAt, Items: input.Items, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&orderInput)

	c.JSON(http.StatusOK, gin.H{"order": orderInput})
}

func UpdateOrder(c *gin.Context) {
	var db = database.GetDB()

	var order models.Order
	err := db.First(&order, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Order
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&order).Where("Id = ?", c.Param("id")).Updates(models.Order{CustomerName: input.CustomerName, OrderedAt: input.OrderedAt, Items: input.Items, UpdatedAt: time.Now()})

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order with id %v has been successfully updated", c.Param("id")),
	})
}

func DeleteOrder(c *gin.Context) {
	var db = database.GetDB()
	var orderDelete models.Order
	err := db.First(&orderDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&orderDelete)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("order with id %v has been successfully deleted", c.Param("id")),
	})
}
