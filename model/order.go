package model

import (
	"online_food_market/database"

	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	CustomerID   uint         `json:"customer_id"`
	RestaurantID uint         `json:"restaurant_id"`
	TotalAmount  float64      `json:"total_amount"`
	DeliveryInfo DeliveryInfo `json:"-"`
}

func (o *Order) Save() (*Order, error) {
	err := database.DB.Create(&o).Error
	if err != nil {
		return &Order{}, err
	}
	return o, nil
}
