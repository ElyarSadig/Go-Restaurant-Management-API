package model

import (
	"online_food_market/database"

	"gorm.io/gorm"
)

type Restaurant struct {
	gorm.Model
	RestaurantName string     `json:"restaurant_name" gorm:"size:255;not null"`
	CuisineType    string     `json:"cuisine_type" gorm:"size:100"`
	ContactEmail   string     `json:"contact_email" gorm:"size:255;unique;not null"`
	ContactPhone   string     `json:"contact_phone" gorm:"size:20"`
	Address        string     `json:"address" gorm:"size:255"`
	MenuItems      []MenuItem `json:"-"`
	Reviews        []Review   `json:"-"`
	Orders         []Order    `json:"-"`
}

func (r *Restaurant) Save() (*Restaurant, error) {
	err := database.DB.Create(&r).Error
	if err != nil {
		return &Restaurant{}, err
	}
	return r, nil
}
