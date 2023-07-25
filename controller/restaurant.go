package controller

import (
	"net/http"
	"online_food_market/database"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllRestaurants(context *gin.Context) {
	var restaurants []model.Restaurant

	if err := database.DB.Find(&restaurants).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch customers"})
		return
	}

	context.JSON(http.StatusOK, restaurants)
}

func GetRestaurant(context *gin.Context) {
	var restaurant model.Restaurant

	restaurantID := context.Param("restaurant_id")

	if err := database.DB.First(&restaurant, restaurantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Restaurant not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurant"})
		}
		return
	}

	context.JSON(http.StatusOK, restaurant)
}

func CreateRestaurant(context *gin.Context) {
	var input model.Restaurant

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurant := model.Restaurant{
		RestaurantName: input.RestaurantName,
		ContactEmail:   input.ContactEmail,
		CuisineType:    input.CuisineType,
		ContactPhone:   input.ContactPhone,
		Address:        input.Address,
	}

	saved, err := restaurant.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusCreated, saved)

}

func UpdateRestaurant(context *gin.Context) {
	var restaurant model.Restaurant
	restaurantID := context.Param("restaurant_id")

	if err := database.DB.First(&restaurant, restaurantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Restaurnat not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch the restaurant"})
		}
		return
	}

	var input struct {
		RestauranName string `json:"restaurant_name"`
		CuisineType   string `json:"cuisine_type"`
		ContactPhone  string `json:"contact_phone"`
		Address       string `json:"address"`
	}

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	restaurant.Address = input.Address
	restaurant.CuisineType = input.CuisineType
	restaurant.ContactPhone = input.ContactPhone
	restaurant.RestaurantName = input.RestauranName

	if err := database.DB.Save(&restaurant).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update the restaurant"})
		return
	}

	context.JSON(http.StatusOK, restaurant)
}

func DeleteRestaurant(context *gin.Context) {
	var restaurant model.Restaurant
	restaurantID := context.Param("restaurant_id")

	if err := database.DB.First(&restaurant, restaurantID).Delete(&restaurant).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "Restaurnat not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete the restaurant"})
		}
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Restaurant deleted successfully"})
}
