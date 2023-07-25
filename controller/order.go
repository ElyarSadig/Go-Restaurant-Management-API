package controller

import (
	"net/http"
	"online_food_market/database"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllOrders(context *gin.Context) {
	var orders []model.Order
	customerID := context.Param("customer_id")

	if err := database.DB.First(&model.Customer{}, customerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	if err := database.DB.Where("customer_id=?", customerID).Find(&orders).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}

	context.JSON(http.StatusOK, orders)
}

func GetOrder(context *gin.Context) {
	var order model.Order
	orderID := context.Param("order_id")

	if err := database.DB.First(&order, orderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Order"})
		}
		return
	}

	context.JSON(http.StatusOK, order)
}

func CreateOrder(context *gin.Context) {
	var input model.Order
	customerID := context.Param("customer_id")

	if err := database.DB.First(&model.Customer{}, customerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Customer not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, order)
}

func UpdateOrder(context *gin.Context) {
	var order model.Order
	orderID := context.Param("order_id")

	if err := database.DB.First(&order, orderID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Order not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	if err := context.ShouldBindJSON(&order); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&order).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the order"})
		return
	}
	context.JSON(http.StatusOK, order)
}

func DeleteOrder(context *gin.Context) {
	var order model.Order
	orderID := context.Param("order_id")

	if err := database.DB.First(&order, orderID).Delete(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the order"})
		}
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
