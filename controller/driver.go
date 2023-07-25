package controller

import (
	"net/http"
	"online_food_market/database"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllDrivers(context *gin.Context) {
	var drivers []model.Driver

	if err := database.DB.Find(&drivers).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch drivers"})
		return
	}

	context.JSON(http.StatusOK, drivers)
}

func GetDriver(context *gin.Context) {
	var driver model.Driver
	driverID := context.Param("driver_id")

	if err := database.DB.First(&driver, driverID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	context.JSON(http.StatusOK, driver)
}

func CreateDriver(context *gin.Context) {
	var input model.Driver

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, driver)
}

func UpdateDriver(context *gin.Context) {
	var driver model.Driver
	driverID := context.Param("driver_id")

	if err := database.DB.First(&driver, driverID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the driver"})
		}
		return
	}

	if err := context.ShouldBindJSON(&driver); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&driver).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the driver"})
		return
	}

	context.JSON(http.StatusOK, driver)
}

func DeleteDriver(context *gin.Context) {
	var driver model.Driver
	driverID := context.Param("driver_id")

	if err := database.DB.First(&driver, driverID).Delete(&driver).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Driver not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the driver"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Driver deleted successfully"})
}
