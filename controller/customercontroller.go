package controller

import (
	"go/build"
	"log"
	"net/http"
	"project/data/model"
	"project/data/service"
	"text/template"
)

var updatecustomer bool

//GetallVehicleWithBrandforview is..
func GetallVehicleWithBrandforview(w http.ResponseWriter, r *http.Request) {
	path := build.Default.GOPATH + "/src/project/template/customer/*"
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
	vehicles := service.GetParticlullarBrandVehiclewithR(r)
	tpl.ExecuteTemplate(w, "vehiclelist.html", struct {
		Vehicles []model.Vehicle
	}{vehicles})
}

//CustomerGetoneVehicleforview is...
func CustomerGetoneVehicleforview(w http.ResponseWriter, r *http.Request) {
	vehicle := service.GetOneVehicle(r)
	path := build.Default.GOPATH + "/src/project/template/customer/*"
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
	tpl.ExecuteTemplate(w, "viewvehicle.html", vehicle)
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
	path := build.Default.GOPATH + "/src/project/template/customer/*"
	tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
	tpl.ExecuteTemplate(w, "account.html", struct {
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
