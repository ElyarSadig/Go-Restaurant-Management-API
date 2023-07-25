package model

import (
	"online_food_market/database"

	"gorm.io/gorm"
)

type MenuItem struct {
	gorm.Model
	RestaurantID       uint    `json:"restaurant_id"`
	ItemName           string  `json:"item_name" gorm:"size:255;not null" binding:"required"`
	Description        string  `json:"description" gorm:"type:text"`
	Price              float64 `json:"price" binding:"required"`
	AvailabilityStatus bool    `json:"availability_status"`
}

func (m *MenuItem) Save() (*MenuItem, error) {
	err := database.DB.Create(&m).Error
	if err != nil {
		return &MenuItem{}, err
	}

	return m, nil
}
