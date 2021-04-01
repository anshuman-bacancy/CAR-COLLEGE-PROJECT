package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"project/common"
	"project/data/model"
	"strconv"

	"github.com/gorilla/mux"
)

//SaveCustomer is...
func SaveCustomer(r *http.Request) (bool, error) {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	sqldb := connection.DB()
	rows, err := sqldb.Query("SELECT email FROM customers")
	defer rows.Close()
	if err != nil {
		return false, err
	}
	emails := make(map[string]bool)
	for rows.Next() {
		var email string
		rows.Scan(&email)
		emails[email] = true
	}

	if emails[r.FormValue("email")] {
		return true, nil
	}
	customer := model.Customer{
		Name:     r.FormValue("name"),
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
		Mobile:   r.FormValue("mobilenumber"),
		City:     r.FormValue("city"),
	}
	connection.Create(&customer)
	return false, nil
}

//Returns all customers
func GetAllCustomer(r *http.Request) []model.Customer {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var customers []model.Customer
	connection.Find(&customers)
	return customers
}

// Deletes one customer
func DeleteOneCustomer(r *http.Request) {
	id := mux.Vars(r)["id"]
	var customer model.Customer
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	connection.Delete(&customer, id)
}

// Returns one customer
func GetOneCustomer(r *http.Request) model.Customer {
	id := mux.Vars(r)["id"]
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var customer model.Customer
	connection.First(&customer, id)
	return customer
}

// Returns one customer via email
func GetOneCustomerBYemail(email interface{}) model.Customer {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var customer model.Customer
	connection.Where("email = 	?", email).First(&customer)
	return customer
}

//CustomerUpdate is...
func CustomerUpdate(r *http.Request) ([]byte, error) {
	fmt.Println("called update")
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	id := mux.Vars(r)["id"]
	var customer model.Customer
	connection.First(&customer, id)
	bodydata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodydata, &customer)
	if err != nil {
		return nil, err
	}
	connection.Save(&customer)
	bytedata, err := json.MarshalIndent(customer, "", "  ")
	if err != nil {
		return nil, err
	}
	return bytedata, nil
}

//CustomerBookVehicle is....
func CustomerBookVehicle(r *http.Request, customer model.Customer) error {
	vehicleid, err := strconv.Atoi(r.FormValue("vehicleid"))
	if err != nil {
		return err
	}
	customerid := customer.ID
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	order := model.TestDrive{
		VehicleID:  uint(vehicleid),
		CustomerID: customerid,
	}
	connection.Create(&order)
	return nil
}

//GetParticlullarCustomerTestDrive is..
func GetParticlullarCustomerTestDrive(r *http.Request, customer model.Customer) []model.TestDrive {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var orders []model.TestDrive
	connection.Where("customer_id = ?", customer.ID).Find(&orders)
	return orders
}

func GetAllTestDrives(r *http.Request) []model.TestDrive {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var orders []model.TestDrive
	connection.Find(&orders)
	return orders
}

func GetCustomerNameByID(id uint) string {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var customer model.Customer
	connection.First(&customer, id)
	return customer.Name
}

func SaveCustomerTestDrive(customer model.Customer, vehicleId uint64, testDriveDate string) error {
	var testDrive model.TestDrive
	testDrive.CustomerID = customer.ID
	testDrive.VehicleID = uint(vehicleId)
	testDrive.TestDriveDate = testDriveDate

	connection := common.GetDatabase()
	defer common.Closedatabase(connection)

	connection.Create(&testDrive)
	return nil
}
