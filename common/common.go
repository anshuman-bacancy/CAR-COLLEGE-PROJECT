package common

import (
	"database/sql"
	"fmt"
	"project/data/model"
	"text/template"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *sql.DB
var err error
var tpl *template.Template

//GetDatabase is return db connection
func GetDatabase() *gorm.DB {
	connection, err := gorm.Open(postgres.Open("postgres://postgres:1312@localhost/CarProject?sslmode=disable"), &gorm.Config{})
	CheckError(err)
	sqldb, err := connection.DB()
	CheckError(err)
	err = sqldb.Ping()
	CheckError(err)
	fmt.Println("connected to database")
	return connection
}

//Closedatabase is...
func Closedatabase(connection *gorm.DB) {
	sqldb, err := connection.DB()
	CheckError(err)
	sqldb.Close()
}

//CheckError is...
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//Initialmigration is....
func Initialmigration() {
	connection := GetDatabase()
	connection.AutoMigrate(&model.SalesPerson{})
	connection.AutoMigrate(&model.Vehicle{})
	connection.AutoMigrate(&model.Customer{})
	defer Closedatabase(connection)
	fmt.Println("migration done")
}
