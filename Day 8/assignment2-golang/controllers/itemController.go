package controllers

import (
	"assignment2-golang/database"
	"assignment2-golang/models"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetAllItems(c *gin.Context) {
	var db = database.GetDB()

	var items []models.Item
	err := db.Find(&items).Error

	if err != nil {
		fmt.Println("Error getting item datas: ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func GetItem(c *gin.Context) {
	var db = database.GetDB()

	var item models.Item
	err := db.First(&item, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

func CreateItem(c *gin.Context) {
	var db = database.GetDB()
	var input models.Item

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	itemInput := models.Item{ItemCode: input.ItemCode, Description: input.Description, Quantity: input.Quantity, OrderID: input.OrderID, CreatedAt: time.Now(), UpdatedAt: time.Now()}
	db.Create(&itemInput)

	c.JSON(http.StatusOK, gin.H{"item": itemInput})
}

func UpdateItem(c *gin.Context) {
	var db = database.GetDB()

	var item models.Item
	err := db.First(&item, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input models.Item
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&item).Where("Id = ?", c.Param("id")).Updates(models.Item{ItemCode: input.ItemCode, Description: input.Description, Quantity: input.Quantity, OrderID: input.OrderID, UpdatedAt: time.Now()})

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("item with id %v has been successfully updated", c.Param("id")),
	})
}

func DeleteItem(c *gin.Context) {
	var db = database.GetDB()
	var itemDelete models.Item
	err := db.First(&itemDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&itemDelete)

	c.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("item with id %v has been successfully deleted", c.Param("id")),
	})
}
