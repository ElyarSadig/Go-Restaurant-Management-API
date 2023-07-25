package controller

import (
	"net/http"
	"online_food_market/database"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetDeliveryInfo(context *gin.Context) {
	var deliveryInfo model.DeliveryInfo
	orderID := context.Param("order_id")

	if err := database.DB.First(&model.Order{}, orderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	if err := database.DB.Where("order_id=?", orderID).First(&deliveryInfo).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Order does not have deliveryInfo"})
		return
	}

	context.JSON(http.StatusOK, deliveryInfo)
}

func CreateDeliveryInfo(context *gin.Context) {
	var deliveryInfo model.DeliveryInfo
	orderID := context.Param("order_id")

	if err := database.DB.First(&model.Order{}, orderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	if err := context.ShouldBindJSON(&deliveryInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driverID := deliveryInfo.DriverID

	if err := database.DB.First(&model.Driver{}, driverID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "driver not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	_, err := deliveryInfo.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "cannot save to the database!"})
		return
	}

	context.JSON(http.StatusOK, deliveryInfo)
}

func UpdateDeliveryInfo(context *gin.Context) {
	var deliveryInfo model.DeliveryInfo
	deliveryID := context.Param("delivery_id")

	if err := database.DB.First(&deliveryInfo, deliveryID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "DeliveryInfo not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete deliveryInfo"})
		}
		return
	}

	if err := context.ShouldBindJSON(&deliveryInfo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	if err := database.DB.Save(&deliveryInfo).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, deliveryInfo)

}

func DeleteDeliveryInfo(context *gin.Context) {
	deliveryID := context.Param("delivery_id")
	var deliveryInfo model.DeliveryInfo

	if err := database.DB.First(&deliveryInfo, deliveryID).Delete(&deliveryInfo).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "DeliveryInfo not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete deliveryInfo"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "deliveryInfo deleted successfully"})
}
