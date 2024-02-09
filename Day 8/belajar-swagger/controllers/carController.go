package controllers

import (
	"belajar-swagger/database"
	"belajar-swagger/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAllCars godoc
// @Summary Get details
// @Description Get details of all car
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} models.Car
// @Router /cars [get]
func GetAllCars(c *gin.Context) {
	var db = database.GetDB()

	var cars []models.Car
	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car datas: ", err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"data": cars})
}

// GetOneCar godoc
// @Summary Get details for a given id
// @Description Get details of a car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} models.Car
// @Router /cars/{Id} [get]
func GetOneCar(c *gin.Context) {
	var db = database.GetDB()

	var carOne models.Car
	err := db.First(&carOne, "Id = ?", c.Param("id")).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data One": carOne})
}

// CreateCars godoc
// @Summary Post details
// @Description Post details of a car
// @Tags cars
// @Accept json
// @Produce json
// @Param models.Car body models.Car true "create car"
// @Success 200 {object} models.Car
// @Router /cars [post]
func CreateCars(c *gin.Context) {
	var db = database.GetDB()
	// Validate input
	var input models.Car

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create car
	carInput := models.Car{Merk: input.Merk, Pemilik: input.Pemilik, Harga: input.Harga, Typecars: input.Typecars}
	db.Create(&carInput)

	c.JSON(http.StatusOK, gin.H{"data": carInput})
}

// UpdateCars godoc
// @Summary Update car identified by the given Id
// @Description Update the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be updated"
// @Success 200 {object} models.Car
// @Router /cars/{id} [patch]
func UpdateCars(c *gin.Context) {
	var db = database.GetDB()

	var car models.Car
	err := db.First(&car, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	//Validate input
	var input models.Car
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}

// DeleteCars godoc
// @Summary Delete car identified by the given Id
// @Description Delete the car corresponding to the input Id
// @Tags cars
// @Accept json
// @Produce json
// @Param id path int true "ID of the car to be deleted"
// @Success 204 "No Content"
// @Router /cars/{id} [delete]
func DeleteCars(c *gin.Context) {
	var db = database.GetDB()
	// Get model if exist
	var carDelete models.Car
	err := db.First(&carDelete, "Id = ?", c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&carDelete)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
