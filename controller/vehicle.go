package controller

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"project/models"
	"project/services"

	"github.com/gorilla/sessions"
)

var updateadmin bool
var adminregister bool
var adminregisteremailerror bool
var checkError bool
var store = sessions.NewCookieStore([]byte("t0p-s3cr3ta"))
var vehiclesave bool
var deletevehicle bool
var updatevehicle bool
var brandsave bool
var deletebrand bool
var updatebrand bool
var deletecustomer bool

var Fm = template.FuncMap{
	"getbrand":             getbrand,
	"getVehicleName":       getVehicleName,
	"getVehicleImage":      getVehicleImage,
	"getVehicleBrandName":  getVehicleBrandName,
	"getVehicleBrandImage": getVehicleBrandImage,
	"getCustomerNameById":  getCustomerNameByID,
}

func getCustomerNameByID(customerid uint) string {
	return services.GetCustomerNameByID(customerid)
}

func getbrand(brandid uint) string {
	return services.GetOneBrandNameByID(brandid)
}

func getbrandImage(brandid uint) string {
	return services.GetOneBrandImageByID(brandid)
}

func getVehicleName(vehicle uint) string {
	return services.GetOneVehicleNameByID(vehicle)
}

func getVehicleImage(vehicle uint) string {
	return services.GetOneVehicleImageByID(vehicle)
}

func getVehicleBrandName(vehicle uint) string {
	brandid := services.GetVehicleBrandID(vehicle)
	return getbrand(brandid)
}

func getVehicleBrandImage(vehicle uint) string {
	brandid := services.GetVehicleBrandID(vehicle)
	return getbrandImage(brandid)
}

//AdminIndexPageProcess is..
func AdminIndexPageProcess(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if vehiclesave {
		vehiclesave = false
		hasmessge = true
		message = "Vehicle data stored successfully"
	}

	if deletevehicle {
		deletevehicle = false
		hasmessge = true
		message = "Vehicle deleted successfully"
	}

	if updatevehicle {
		updatevehicle = false
		hasmessge = true
		message = "Vehicle updated successfully"
	}
	vehicles := services.GetAllVehicle()
	admintpl.ExecuteTemplate(w, "index.html", struct {
		HasMessage bool
		Message    string
		Vehicles   []models.Vehicle
	}{hasmessge, message, vehicles})
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	admintpl.ExecuteTemplate(w, "404.html", nil)
}

// show admin login
func Login(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if checkError {
		hasmessge = true
		message = "Username or Password is wrong"
		checkError = false
	}

	admintpl.ExecuteTemplate(w, "login.html", struct {
		HasMessage bool
		Message    string
	}{hasmessge, message})
}

// admin login
func LoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	adminemail := r.PostForm.Get("username")
	adminpassword := r.PostForm.Get("password")
	admins := services.GetAllAdmin(r)
	if len(admins) == 0 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var accessadmin bool
	for _, admin := range admins {
		if admin.Email == adminemail && admin.Password == adminpassword {
			accessadmin = true
			session, _ := store.Get(r, "username")
			session.Values["username"] = admin.Email
			session.Save(r, w)
			break
		}
	}

	if !accessadmin {
		checkError = true
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/admin/vehicle", http.StatusSeeOther)
}

// admin logout
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "username")
	session.Options.MaxAge = -1
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// create vehicle
func CreateVehicleform(w http.ResponseWriter, r *http.Request) {
	brands := services.GetAllBrands(r)
	admintpl.ExecuteTemplate(w, "createvehicle.html", brands)
}

// admin authentication
func AuthenticationAdmin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "username")
		_, ok := session.Values["username"]
		if !ok {
			checkError = true
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

// serve 404 page
func ServerError(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "error.html", nil)
}

// save vehicle
func SaveVehicle(w http.ResponseWriter, r *http.Request) {
	err := services.SaveVehicle(r)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	vehiclesave = true
	http.Redirect(w, r, "/admin/vehicle", http.StatusSeeOther)
}

// view one vehicle
func GetOneVehicleForView(w http.ResponseWriter, r *http.Request) {
	vehicle := services.GetOneVehicle(r)
	admintpl.ExecuteTemplate(w, "viewvehicle.html", vehicle)
}

// view one vehicle for edit
func GetOneVehicleForEdit(w http.ResponseWriter, r *http.Request) {
	vehicle := services.GetOneVehicle(r)
	brands := services.GetAllBrands(r)
	admintpl.ExecuteTemplate(w, "editvehicle.html", struct {
		Vehicle models.Vehicle
		Brands  []models.Company
	}{vehicle, brands})
}

// delete one vehicle
func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	services.DeleteOneVehicle(r)
	deletevehicle = true
}

// update vehicle
func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	data, err := services.UpdateVehicle(r)
	if err != nil {
		log.Fatalln(err)
		return
	}
	updatevehicle = true
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// brand form
func CreateBrandForm(w http.ResponseWriter, r *http.Request) {
	admintpl.ExecuteTemplate(w, "createbrand.html", nil)
}

// save brand
func SaveBrand(w http.ResponseWriter, r *http.Request) {
	err := services.SaveCompanyLogo(r)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	brandsave = true
	http.Redirect(w, r, "/admin/brand", http.StatusSeeOther)
}

// get all brands
func GetAllBrands(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool

	if brandsave {
		hasmessge = true
		message = "Brand data stored successfully"
		brandsave = false
	}

	if updatebrand {
		updatebrand = false
		hasmessge = true
		message = "Brand updated successfully"
	}

	if deletebrand {
		deletebrand = false
		hasmessge = true
		message = "Brand deleted successfully"
	}

	brands := services.GetAllBrands(r)

	admintpl.ExecuteTemplate(w, "brandlist.html", struct {
		HasMessage bool
		Message    string
		Brands     []models.Company
	}{hasmessge, message, brands})
}

// delete brand
func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	services.DeleteOneBrand(r)
	deletebrand = true
}

// edit brand
func GetOneBrandForEdit(w http.ResponseWriter, r *http.Request) {
	brand := services.GetOneBrand(r)
	admintpl.ExecuteTemplate(w, "editbrand.html", brand)
}

// update brand
func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	data, err := services.UpdateBrand(r)
	if err != nil {
		log.Fatalln(err)
		return
	}
	updatebrand = true
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// get one brand
func GetOneBrandForView(w http.ResponseWriter, r *http.Request) {
	brand := services.GetOneBrand(r)
	vehicles := services.GetParticularBrandVehicle(brand.ID)
	admintpl.ExecuteTemplate(w, "viewbrand.html", struct {
		Brand    models.Company
		Vehicles []models.Vehicle
	}{brand, vehicles})
}

// get all customer
func GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool

	if deletecustomer {
		hasmessge = true
		message = "Customer deleted successfully"
		deletecustomer = false
	}

	customer := services.GetAllCustomers(r)
	admintpl.ExecuteTemplate(w, "customerlist.html", struct {
		HasMessage bool
		Message    string
		Customers  []models.Customer
	}{hasmessge, message, customer})
}

// delete customer
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	deletecustomer = true
	services.DeleteOneCustomer(r)
}

// get one customer
func GetOneCustomerForView(w http.ResponseWriter, r *http.Request) {
	customer := services.GetOneCustomer(r)
	custTestDrives := services.GetParticularCustomerTestDrive(r, customer)
	admintpl.ExecuteTemplate(w, "viewcustomer.html", struct {
		Customer           models.Customer
		CustomerTestDrives []models.TestDrive
	}{customer, custTestDrives})
}

// get all customer test drives
func GetAllCustomersOrders(w http.ResponseWriter, r *http.Request) {
	orders := services.GetAllTestDrives(r)
	admintpl.ExecuteTemplate(w, "order.html", struct {
		Orders []models.TestDrive
	}{orders})
}

// register admin
func AdminRegister(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	var hasmessgesuccess bool
	if adminregisteremailerror {
		hasmessge = true
		message = "Email is already taken"
		adminregisteremailerror = false
	}

	if adminregister {
		hasmessgesuccess = true
		message = "Admin Registration successfully"
		adminregister = false
	}

	admintpl.ExecuteTemplate(w, "register.html", struct {
		HasMessage       bool
		Message          string
		Hasmessgesuccess bool
	}{hasmessge, message, hasmessgesuccess})
}

func AdminRegisterPOST(w http.ResponseWriter, r *http.Request) {
	emailcheck, err := services.SaveAdmin(r)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	if emailcheck {
		adminregisteremailerror = true
		http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
		return
	}
	adminregister = true
	http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
}

// admin account
func GetAdminAccountPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "username")
	email, _ := session.Values["username"]
	admin := services.GetOneAdminByEmail(email)
	var message string
	var hasmessge bool
	if updateadmin {
		hasmessge = true
		message = "Profile Updated successfully"
		updateadmin = false
	}
	admintpl.ExecuteTemplate(w, "account.html", struct {
		HasMessage bool
		Message    string
		Admin      models.SalesPerson
	}{hasmessge, message, admin})
}

// update admin
func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	customer, err := services.AdminUpdate(r)
	if err != nil {
		log.Fatalln(err)
		return
	}
	updateadmin = true
	w.Header().Set("Content-Type", "application/json")
	w.Write(customer)
}

// update test drive
func UpdateCustomerTestDriveStatus(w http.ResponseWriter, r *http.Request) {
	var data models.TestDriveStatus
	json.NewDecoder(r.Body).Decode(&data)
	services.UpdateCustomerTestDriveStatus(data)
}
