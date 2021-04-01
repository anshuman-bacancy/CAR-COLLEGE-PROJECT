package controller

import (
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

//GetallVehicleWithBrandforview is..
func GetallVehicleWithBrandforview(w http.ResponseWriter, r *http.Request) {
	vehicles := service.GetParticlullarBrandVehiclewithR(r)
	// custtpl.ExecuteTemplate(w, "vehiclelist.html", struct {
	// 	Vehicles []model.Vehicle
	// }{vehicles})

	custtpl.ExecuteTemplate(w, "vehiclelist.html", vehicles)
}

//CustomerGetoneVehicleforview is...
func CustomerGetoneVehicleforview(w http.ResponseWriter, r *http.Request) {
	vehicle := service.GetOneVehicle(r)
	custtpl.ExecuteTemplate(w, "viewvehicle.html", vehicle)
}

//CustomerAccountforview is..
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

//CustomerUpdate is...
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

//CustomerBookVehicle is.
func CustomerBookVehicle(w http.ResponseWriter, r *http.Request) {
	session, _ := storecustomer.Get(r, "customerusername")
	email, _ := session.Values["customer"]
	customer := service.GetOneCustomerBYemail(email)
	err := service.CustomerBookVehicle(r, customer)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	ordersave = true
	http.Redirect(w, r, "/customer/orders", http.StatusSeeOther)
}

//CustomerGetallOrder is..
func CustomerGetallOrder(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if ordersave {
		ordersave = false
		message = "Your order process successfully"
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

func CustomerTestDrive(w http.ResponseWriter, r *http.Request) {
	tempVehicleId := r.FormValue("vehicleId")
	vehicleId, _ := strconv.ParseUint(tempVehicleId, 10, 64)
	testDriveDate := common.FormatDate(r.FormValue("testDriveDate"))
	// fmt.Println(vehicleId, testDriveDate)

	// get customer
	session, _ := storecustomer.Get(r, "customerusername")
	email, _ := session.Values["customer"]
	customer := service.GetOneCustomerBYemail(email)

	//save to db
	err := service.SaveCustomerTestDrive(customer, vehicleId, testDriveDate)
	if err != nil {
		log.Println(err)
	}

	// redirect to customer/index
	http.Redirect(w, r, "/customer/index", 302)
}
