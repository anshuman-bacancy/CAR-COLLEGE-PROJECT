package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"project/common"
	"project/data/model"
	"project/data/service"
	"strconv"
)

var (
	updatecustomer bool
	ordersave      bool
)

var tpl *template.Template

// returns vehicle with brands
func GetallVehicleWithBrandforview(w http.ResponseWriter, r *http.Request) {
	vehicles := service.GetParticlullarBrandVehiclewithR(r)
	custtpl.ExecuteTemplate(w, "vehiclelist.html", vehicles)
}

// returns one vehicle for customer
func CustomerGetoneVehicleforview(w http.ResponseWriter, r *http.Request) {
	vehicle := service.GetOneVehicle(r)
	custtpl.ExecuteTemplate(w, "viewvehicle.html", vehicle)
}

// shows customer account
func CustomerAccountforview(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	email, _ := session.Values["customer"]
	customer := service.GetOneCustomerBYemail(email)
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
	customer, err := service.CustomerUpdate(r)
	if err != nil {
		log.Fatalln(err)
		return
	}
	updatecustomer = true
	w.Header().Set("Content-Type", "application/json")
	w.Write(customer)
}

// shows all test drives for customer
func CustomerGetallOrder(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if ordersave {
		ordersave = false
		message = "Your test drive has been booked successfully."
		hasmessge = true
	}
	session, _ := storecustomer.Get(r, "customerusername")
	email, _ := session.Values["customer"]
	customer := service.GetOneCustomerBYemail(email)
	orders := service.GetParticlullarCustomerTestDrive(r, customer)
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
	customer := service.GetOneCustomerBYemail(email)

	//save to db
	ordersave = true
	err := service.SaveCustomerTestDrive(customer, vehicleId, testDriveDate)
	if err != nil {
		log.Println(err)
	}
	http.Redirect(w, r, "/customer/orders", 302)
}

// returns list of vehicles in dropdown 
func CarCompare(w http.ResponseWriter, r *http.Request) {
	allVehicles := service.GetAllVehicle()
	custtpl.ExecuteTemplate(w, "compareCar.html", allVehicles)
}

// returns vehicle for car compare
func CustomerGetVehicle(w http.ResponseWriter, r *http.Request) {
	vehicle := service.GetOneVehicle(r)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicle)
}