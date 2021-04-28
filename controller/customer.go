package controller

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"project/common"
	model "project/models"
	"project/services"
	"strconv"
)

var (
	updatecustomer bool
	ordersave      bool
	tpl *template.Template
)


// returns vehicle with brands
func GetAllVehicleWithBrandForView(w http.ResponseWriter, r *http.Request) {
	vehicles := services.GetParticularBrandVehiclewithR(r)
	custtpl.ExecuteTemplate(w, "vehiclelist.html", vehicles)
}

// returns one vehicle for customer
func CustomerGetoneVehicleforview(w http.ResponseWriter, r *http.Request) {
	vehicle := services.GetOneVehicle(r)
	custtpl.ExecuteTemplate(w, "viewvehicle.html", vehicle)
}

// shows customer account
func CustomerAccountForView(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	email, _ := session.Values["customer"]
	customer := services.GetOneCustomerBYemail(email)
	var message string
	var hasmessge bool
	if updatecustomer {
		hasmessge = true
		message = "Profile Updated successfully"
		updatecustomer = false
	}
	custtpl.ExecuteTemplate(w, "account.html", struct {
		HasMessage bool
		Message    string
		Customer   model.Customer
	}{hasmessge, message, customer})
}

// customer update
func CustomerUpdate(w http.ResponseWriter, r *http.Request) {
	customer, err := services.CustomerUpdate(r)
	if err != nil {
		log.Fatalln(err)
		return
	}
	updatecustomer = true
	w.Header().Set("Content-Type", "application/json")
	w.Write(customer)
}

// shows all test drives for customer
func CustomerGetAllOrders(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if ordersave {
		ordersave = false
		message = "Your test drive has been booked successfully."
		hasmessge = true
	}
	session, _ := storecustomer.Get(r, "customerusername")
	email, _ := session.Values["customer"]
	customer := services.GetOneCustomerBYemail(email)
	orders := services.GetParticularCustomerTestDrive(r, customer)
	custtpl.ExecuteTemplate(w, "order.html", struct {
		Orders     []model.TestDrive
		HasMessage bool
		Message    string
	}{orders, hasmessge, message})
}

// saves customer test drive
func CustomerTestDrive(w http.ResponseWriter, r *http.Request) {
	tempVehicleId := r.FormValue("vehicleId")
	vehicleId, _ := strconv.ParseUint(tempVehicleId, 10, 64)
	testDriveDate := common.FormatDate(r.FormValue("testDriveDate"))

	// get customer
	session, _ := storecustomer.Get(r, "customerusername")
	email, _ := session.Values["customer"]
	customer := services.GetOneCustomerBYemail(email)

	//save to db
	ordersave = true
	err := services.SaveCustomerTestDrive(customer, vehicleId, testDriveDate)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/customer/orders", 302)
}

// returns list of vehicles in dropdown
func CompareCar(w http.ResponseWriter, r *http.Request) {
	allVehicles := services.GetAllVehicle()
	ids := make(map[string]uint, 0)

	for _, vehicle := range allVehicles {
		ids["car"+strconv.Itoa(int(vehicle.ID))] = vehicle.ID
	}

	fmt.Println(ids)

	custtpl.ExecuteTemplate(w, "compareCar.html", struct {
		AllVehicles []model.Vehicle
		Ids         map[string]uint
	}{allVehicles, ids})
}

// returns vehicle for car compare
func CustomerGetVehicle(w http.ResponseWriter, r *http.Request) {
	vehicle := services.GetOneVehicle(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicle)
}
