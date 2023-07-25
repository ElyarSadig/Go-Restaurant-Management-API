package model

import (
	"online_food_market/database"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	CustomerName string   `json:"customer_name" gorm:"size:255;not null"`
	Email        string   `json:"email" gorm:"size:255;unique;not null"`
	Phone        string   `json:"phone" gorm:"size:20"`
	Address      string   `json:"address" gorm:"size:255"`
	Orders       []Order  `json:"-"`
	Reviews      []Review `json:"-"`
}

func (c *Customer) Save() (*Customer, error) {
	err := database.DB.Create(&c).Error
	if err != nil {
		return &Customer{}, err
	}
	return c, nil
}
