package common

import (
	"database/sql"
	"log"
	"project/models"
	"strings"
	"text/template"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	db  *sql.DB
	err error
	tpl *template.Template
)

//GetDatabase is return db connection
func GetDatabase() *gorm.DB {
	connection, err := gorm.Open("postgres", "postgres://postgres:password@localhost/CarProject?sslmode=disable")
	CheckError(err)
	sqldb := connection.DB()
	err = sqldb.Ping()
	CheckError(err)
	log.Println("connected to database")
	return connection
}

//Closedatabase closes the database connection
func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}

// panics error based on severity
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

//Initialmigration migrates modelss to database
func InitialMigration() {
	connection := GetDatabase()
	connection.AutoMigrate(&models.SalesPerson{})
	connection.AutoMigrate(&models.Company{})
	connection.AutoMigrate(&models.Vehicle{})
	connection.Model(&models.Vehicle{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
	connection.AutoMigrate(&models.Customer{})
	connection.AutoMigrate(&models.TestDrive{})
	connection.Model(&models.TestDrive{}).AddForeignKey("vehicle_id", "vehicles(id)", "CASCADE", "CASCADE")
	connection.Model(&models.TestDrive{}).AddForeignKey("customer_id", "customers(id)", "CASCADE", "CASCADE")
	defer CloseDatabase(connection)

	log.Println("Migration done... ")

	// connection.Create(&models.SalesPerson{
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
	testDriveDate := dateFormat.Format(dateLayout)
	return testDriveDate
}
