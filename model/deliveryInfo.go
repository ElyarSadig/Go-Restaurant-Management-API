package model

import (
	"online_food_market/database"
	"time"

	"gorm.io/gorm"
)

type DeliveryInfo struct {
	gorm.Model
	OrderID      uint      `json:"order_id"`
	DriverID     uint      `json:"driver_id"`
	DeliveryDate time.Time `json:"delivery_date"`
	Status       string    `json:"status" gorm:"size:100"`
}

func (d *DeliveryInfo) Save() (*DeliveryInfo, error) {
	err := database.DB.Create(&d).Error
	if err != nil {
		return &DeliveryInfo{}, err
	}
	return d, nil
}
