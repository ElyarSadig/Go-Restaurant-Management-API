package controller

import (
	"net/http"
	"online_food_market/database"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllCustomers(context *gin.Context) {
	var customers []model.Customer

	if err := database.DB.Find(&customers).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}

	context.JSON(http.StatusOK, customers)
}

func GetCustomer(context *gin.Context) {
	var customer model.Customer
	customerID := context.Param("customer_id")

	if err := database.DB.First(&customer, customerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customer"})
		}
		return
	}

	context.JSON(http.StatusOK, customer)
}

func CreateCustomer(context *gin.Context) {
	var input model.Customer

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Incomplete JSON input. Please provide all required fields."})
		return
	}

	customer, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, customer)
}

func UpdateCustomer(context *gin.Context) {
	var customer model.Customer
	customerID := context.Param("customer_id")

	if err := database.DB.First(&customer, customerID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the customer"})
		}
		return
	}

	if err := context.ShouldBindJSON(&customer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&customer).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the customer"})
		return
	}

	context.JSON(http.StatusOK, customer)

}

func DeleteCustomer(context *gin.Context) {
	var customer model.Customer
	customerID := context.Param("customer_id")

	err := database.DB.First(&customer, customerID).Delete(&customer).Error
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Customer not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
