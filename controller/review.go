package controller

import (
	"net/http"
	"online_food_market/database"
	"online_food_market/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllReviews(context *gin.Context) {
	var reviews []model.Review
	restaurantID := context.Param("restaurant_id")

	if err := database.DB.First(&model.Restaurant{}, restaurantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	if err := database.DB.Where("restaurant_id=?", restaurantID).Find(&reviews).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch restaurant reviews"})
		return
	}

	context.JSON(http.StatusOK, reviews)
}

func GetReview(context *gin.Context) {
	var review model.Review
	reviewID := context.Param("review_id")

	if err := database.DB.First(&review, reviewID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	context.JSON(http.StatusOK, review)
}

func CreateReview(context *gin.Context) {
	var review model.Review
	restaurantID := context.Param("restaurant_id")

	if err := database.DB.First(&model.Restaurant{}, restaurantID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "restaurant not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	if err := context.ShouldBindJSON(&review); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := review.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		return
	}

	context.JSON(http.StatusCreated, review)

}

func UpdateReview(context *gin.Context) {
	var review model.Review
	reviewID := context.Param("review_id")

	if err := database.DB.First(&review, reviewID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			context.JSON(http.StatusNotFound, gin.H{"error": "review not found"})
		} else {
			context.JSON(http.StatusInternalServerError, gin.H{"error": "Something went wrong"})
		}
		return
	}

	if err := context.ShouldBindJSON(&review); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&review).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot save to the database"})
		return
	}

	context.JSON(http.StatusOK, review)
}

func DeleteReview(context *gin.Context) {
	var review model.Review
	reviewID := context.Param("review_id")

	if err := database.DB.First(&review, reviewID).Delete(&review).Error; err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Review not found"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "review deleted successfully"})
}
