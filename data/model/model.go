package model

import (
	"gorm.io/gorm"
)

//SalesPerson is...
type SalesPerson struct {
	gorm.Model
	Name        string `json:"name"`
	City        string `json:"city"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
}

//Vehicle is...
type Vehicle struct {
	gorm.Model
	Vin         string `json:"vin"`
	Year        string `json:"year"`
	ModelName   string `json:"model"`
	TitleNumber string `json:"title"`
	Price       string `json:"price"`
	FuelType    string `json:"fueltype"`
	Mileage     string `json:"mileage"`
	Stock       string `json:"stock"`
	Image       string `json:"image"`
}

//Customer is...
type Customer struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	City     string `json:"city"`
}
