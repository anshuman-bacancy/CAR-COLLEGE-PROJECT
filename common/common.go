package common

import (
	"database/sql"
	"fmt"
	"project/data/model"
	"strings"
	"text/template"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *sql.DB
var err error
var tpl *template.Template

//GetDatabase is return db connection
func GetDatabase() *gorm.DB {
	connection, err := gorm.Open("postgres", "postgres://postgres:password@localhost/CarProject?sslmode=disable")
	CheckError(err)
	sqldb := connection.DB()
	err = sqldb.Ping()
	CheckError(err)
	fmt.Println("connected to database")
	return connection
}

//Closedatabase is...
func Closedatabase(connection *gorm.DB) {
	sqldb := connection.DB()
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
	connection.AutoMigrate(&model.Company{})
	connection.AutoMigrate(&model.Vehicle{})
	connection.Model(&model.Vehicle{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
	connection.AutoMigrate(&model.Customer{})
	connection.AutoMigrate(&model.TestDrive{})
	connection.Model(&model.TestDrive{}).AddForeignKey("vehicle_id", "vehicles(id)", "CASCADE", "CASCADE")
	connection.Model(&model.TestDrive{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")
	defer Closedatabase(connection)
	fmt.Println("migration done")

	// connection.Create(&model.SalesPerson{
	// 	Name:     "anshuman",
	// 	Email:    "anshuman@gmail.com",
	// 	Password: "anshu",
	// 	City:     "Vadodara",
	// 	Mobile:   "1234567890"})
}

// FormatDate returns date in dd/mm/yyyy format
func FormatDate(date string) string {
	tempDate := strings.Replace(date, "/", "-", -1)
	dateLayout := "02-01-2006"
	dateFormat, _ := time.Parse(dateLayout, tempDate)
	testDriveDate := dateFormat.Format("02-01-2006")
	return testDriveDate
}
