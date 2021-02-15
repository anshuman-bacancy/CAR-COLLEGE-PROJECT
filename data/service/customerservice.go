package service

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"project/common"
	"project/data/model"

	"github.com/gorilla/mux"
)

//SaveCustomer is...
func SaveCustomer(r *http.Request) (bool, error) {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	sqldb := connection.DB()
	// if err != nil {
	// 	return false, err
	// }
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

//GetAllCustomer is..
func GetAllCustomer(r *http.Request) []model.Customer {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var customers []model.Customer
	connection.Find(&customers)
	return customers
}

//DeleteOneCustomer is..
func DeleteOneCustomer(r *http.Request) {
	id := mux.Vars(r)["id"]
	var customer model.Customer
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	connection.Delete(&customer, id)
}

//GetOneCustomer is...
func GetOneCustomer(r *http.Request) model.Customer {
	id := mux.Vars(r)["id"]
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var customer model.Customer
	connection.First(&customer, id)
	return customer
}

//GetOneCustomerBYemail is...
func GetOneCustomerBYemail(email interface{}) model.Customer {
	connection := common.GetDatabase()
	defer common.Closedatabase(connection)
	var customer model.Customer
	connection.Where("email = 	?", email).First(&customer)
	return customer
}

//CustomerUpdate is...
func CustomerUpdate(r *http.Request) ([]byte, error) {
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
