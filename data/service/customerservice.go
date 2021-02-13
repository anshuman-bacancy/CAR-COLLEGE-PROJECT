package service

import (
	"net/http"
	"project/common"
	"project/data/model"
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
