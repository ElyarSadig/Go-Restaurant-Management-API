package model

import (
	"online_food_market/database"
	"time"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	CustomerID   uint      `json:"customer_id"`
	RestaurantID uint      `json:"restaurant_id"`
	Rating       int       `json:"rating"`
	Comments     string    `json:"comments" gorm:"type:text;not null"`
	ReviewDate   time.Time `json:"review_date"`
}

func (r *Review) Save() (*Review, error) {
	err := database.DB.Create(&r).Error
	if err != nil {
		return &Review{}, err
	}
	return r, nil
}
