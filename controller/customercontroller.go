package controller

import (
	"html/template"
	"log"
	"net/http"
	"project/data/model"
	"project/data/service"
)

var (
	updatecustomer bool
	ordersave      bool
)

var tpl *template.Template

func init() {
	admintpl = template.Must(template.New("").Funcs(fm).ParseGlob(("template/admin/*")))
	hometpl = template.Must(template.New("").Funcs(fm).ParseGlob(("template/home/*")))
	custtpl = template.Must(template.New("").Funcs(fm).ParseGlob(("template/customer/*")))
}

//GetallVehicleWithBrandforview is..
func GetallVehicleWithBrandforview(w http.ResponseWriter, r *http.Request) {
	vehicles := service.GetParticlullarBrandVehiclewithR(r)
	custtpl.ExecuteTemplate(w, "vehiclelist.html", struct {
		Vehicles []model.Vehicle
	}{vehicles})
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
	orders := service.GetParticlullarCustomerOrder(r, customer)
	custtpl.ExecuteTemplate(w, "order.html", struct {
		Orders     []model.Order
		HasMessage bool
		Message    string
	}{orders, hasmessge, message})
}
