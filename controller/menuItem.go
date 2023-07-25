package controller

import (
	"net/http"
	"online_food_market/database"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllMenuItems(context *gin.Context) {
	var menuItems []model.MenuItem
	restaurantID := context.Param("restaurant_id")

	if err := database.DB.First(&model.Restaurant{}, restaurantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Restaurant does not exists"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		}
		return
	}

	if err := database.DB.Where("restaurant_id=?", restaurantID).Find(&menuItems).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu items"})
		return
	}

	context.JSON(http.StatusOK, menuItems)
}

func GetMenuItem(context *gin.Context) {
	var menuItem model.MenuItem
	menuItemID := context.Param("menuitem_id")

	if err := database.DB.First(&menuItem, menuItemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "MenuItem not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menuItem"})
		}
		return
	}

	context.JSON(http.StatusOK, menuItem)
}

func CreateMenuItem(context *gin.Context) {
	var input model.MenuItem
	restaurantID := context.Param("restaurant_id")

	if err := database.DB.First(&model.Restaurant{}, restaurantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusBadRequest, gin.H{"error": "Restaurant not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		}
		return
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menuItem, err := input.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, menuItem)

}

func UpdateMenuItem(context *gin.Context) {
	var menuItem model.MenuItem
	menuItemID := context.Param("menuitem_id")

	if err := database.DB.First(&menuItem, menuItemID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "MenuItem not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the restaurant"})
		}
		return
	}

	var input struct {
		ItemName           string  `json:"item_name"`
		Description        string  `json:"description"`
		Price              float64 `json:"price"`
		AvailabilityStatus bool    `json:"availability_status"`
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	menuItem.ItemName = input.ItemName
	menuItem.Description = input.Description
	menuItem.Price = input.Price
	menuItem.AvailabilityStatus = input.AvailabilityStatus

	if err := database.DB.Save(&menuItem).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the menuItem"})
		return
	}

	context.JSON(http.StatusOK, menuItem)
}

func DeleteMenuItem(context *gin.Context) {
	var menuItem model.MenuItem
	menuItemID := context.Param("menuitem_id")

	if err := database.DB.First(&menuItem, menuItemID).Delete(&menuItem).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "MenuItem not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the menuItem"})
		}
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "MenuItem deleted successfully"})
}
