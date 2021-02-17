package model

import (
	"gorm.io/gorm"
)

//SalesPerson is...
type SalesPerson struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	City     string `json:"city"`
}

//Company is...
type Company struct {
	gorm.Model
	Name string `json:"name"`
	Logo string `json:"logo"`
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
	CompanyID   uint   `json:"companyid,string"`
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

//Order is...
type Order struct {
	gorm.Model
	VehicleID  uint `json:"vehicleid,string"`
	CustomerID uint `json:"Customerid,string"`
}
