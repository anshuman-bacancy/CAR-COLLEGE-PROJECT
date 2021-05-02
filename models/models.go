package models

import (
	"gorm.io/gorm"
)

type SalesPerson struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	City     string `json:"city"`
}

type Company struct {
	gorm.Model
	Name string `json:"name"`
	Logo string `json:"logo"`
}

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

type Customer struct {
	gorm.Model
	Name     string `json:"name"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	City     string `json:"city"`
}

type TestDrive struct {
	gorm.Model
	VehicleID     uint   `json:"vehicleid,string"`
	CustomerID    uint   `json:"Customerid,string"`
	TestDriveDate string `json:"testdrivedate,string"`
	Status        string `json:"Status"`
}

type TestDriveStatus struct {
	TestDriveID uint   `json:"TestDriveID,string"`
	Status      string `json:"Status"`
}

type Config struct {
	DB_DRIVER string `mapstructure:"DB_DRIVER"`
	DB_SOURCE string `mapstructure:"DB_SOURCE"`
}
