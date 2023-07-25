package model

import (
	"online_food_market/database"

	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	DriverName         string         `json:"driver_name" gorm:"size:255;not null"`
	ContactPhone       string         `json:"contact_phone" gorm:"size:20 unique"`
	VehiclePlateNumber string         `json:"plate_number" gorm:"size:20"`
	DeliveryInfos      []DeliveryInfo `json:"-"`
}

func (driver *Driver) Save() (*Driver, error) {
	err := database.DB.Create(&driver).Error
	if err != nil {
		return &Driver{}, err
	}
	return driver, nil
}
