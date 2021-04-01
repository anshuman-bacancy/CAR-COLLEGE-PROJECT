package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"project/data/model"
	"project/data/service"

	"github.com/gorilla/sessions"
)

var updateadmin bool
var adminregister bool
var adminregisteremailerror bool
var chekerror bool
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
	return service.GetCustomerNameByID(customerid)
}

func getbrand(brandid uint) string {
	return service.GetOneBrandNameByID(brandid)
}

func getbrandImage(brandid uint) string {
	return service.GetOneBrandImageByID(brandid)
}

func getVehicleName(vehicle uint) string {
	return service.GetOneVehicleNameByID(vehicle)
}

func getVehicleImage(vehicle uint) string {
	return service.GetOneVehicleImageByID(vehicle)
}

func getVehicleBrandName(vehicle uint) string {
	brandid := service.GetVehicleBrandID(vehicle)
	return getbrand(brandid)
}

func getVehicleBrandImage(vehicle uint) string {
	brandid := service.GetVehicleBrandID(vehicle)
	return getbrandImage(brandid)
}

//AdminIndexpageProcess is..
func AdminIndexpageProcess(w http.ResponseWriter, r *http.Request) {
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
	vehicles := service.GetAllVehicle()
	// path := build.Default.GOPATH + "/src/project/template/admin/*"
	// path := "template/admin/*"

	// tpl := template.Must(template.New("").Funcs(fm).ParseGlob(path))
	admintpl.ExecuteTemplate(w, "index.html", struct {
		HasMessage bool
		Message    string
		Vehicles   []model.Vehicle
	}{hasmessge, message, vehicles})
}

//NotFound is...
func NotFound(w http.ResponseWriter, r *http.Request) {
	admintpl.ExecuteTemplate(w, "404.html", nil)
}

//Login is...
func Login(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool
	if chekerror {
		hasmessge = true
		message = "Username or Password is wrong"
		chekerror = false
	}

	admintpl.ExecuteTemplate(w, "login.html", struct {
		HasMessage bool
		Message    string
	}{hasmessge, message})
}

//LoginPost is...
func LoginPost(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	adminemail := r.PostForm.Get("username")
	adminpassword := r.PostForm.Get("password")
	// fmt.Println(adminemail, adminpassword)
	admins := service.GetAllAdmin(r)
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
		fmt.Println("inside if")
		chekerror = true
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, "/admin/vehicle", http.StatusSeeOther)
}

//Logout is....
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "username")
	session.Options.MaxAge = -1
	delete(session.Values, "username")
	session.Save(r, w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

//CreateVehicleform is....
func CreateVehicleform(w http.ResponseWriter, r *http.Request) {
	brands := service.GetAllBrand(r)
	admintpl.ExecuteTemplate(w, "createvehicle.html", brands)
}

//AuthenticationAdmin is..
func AuthenticationAdmin(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "username")
		_, ok := session.Values["username"]
		if !ok {
			chekerror = true
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
		handler.ServeHTTP(w, r)
	}
}

//ServerError is...
func ServerError(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "error.html", nil)
}

//SaveVehicle is....
func SaveVehicle(w http.ResponseWriter, r *http.Request) {
	err := service.SaveVehicle(r)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	vehiclesave = true
	http.Redirect(w, r, "/admin/vehicle", http.StatusSeeOther)
}

//GetoneVehicleforview is...
func GetoneVehicleforview(w http.ResponseWriter, r *http.Request) {
	vehicle := service.GetOneVehicle(r)
	admintpl.ExecuteTemplate(w, "viewvehicle.html", vehicle)
}

//GetoneVehicleforedit is..
func GetoneVehicleforedit(w http.ResponseWriter, r *http.Request) {
	vehicle := service.GetOneVehicle(r)
	brands := service.GetAllBrand(r)
	admintpl.ExecuteTemplate(w, "editvehicle.html", struct {
		Vehicle model.Vehicle
		Brands  []model.Company
	}{vehicle, brands})
}

//DeleteVehicle is....
func DeleteVehicle(w http.ResponseWriter, r *http.Request) {
	service.DeleteOneVehicle(r)
	deletevehicle = true
}

//UpdateVehicle is....
func UpdateVehicle(w http.ResponseWriter, r *http.Request) {
	data, err := service.UpdateVehicle(r)
	if err != nil {
		log.Fatalln(err)
		return
	}
	updatevehicle = true
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//CreateBrandform is..
func CreateBrandform(w http.ResponseWriter, r *http.Request) {
	admintpl.ExecuteTemplate(w, "createbrand.html", nil)
}

//SaveBrand is..
func SaveBrand(w http.ResponseWriter, r *http.Request) {
	err := service.SaveCompanyLogo(r)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}
	brandsave = true
	http.Redirect(w, r, "/admin/brand", http.StatusSeeOther)
}

//GetAllBrand is....
func GetAllBrand(w http.ResponseWriter, r *http.Request) {
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

	brands := service.GetAllBrand(r)

	fmt.Println(brands)

	admintpl.ExecuteTemplate(w, "brandlist.html", struct {
		HasMessage bool
		Message    string
		Brands     []model.Company
	}{hasmessge, message, brands})
}

//DeleteBrand is....
func DeleteBrand(w http.ResponseWriter, r *http.Request) {
	service.DeleteOneBrand(r)
	deletebrand = true
}

//GetoneBrandforedit is..
func GetoneBrandforedit(w http.ResponseWriter, r *http.Request) {
	brand := service.GetOneBrand(r)
	admintpl.ExecuteTemplate(w, "editbrand.html", brand)
}

//UpdateBrand is...
func UpdateBrand(w http.ResponseWriter, r *http.Request) {
	data, err := service.UpdateBrand(r)
	if err != nil {
		log.Fatalln(err)
		return
	}
	updatebrand = true
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

//GetoneBrandforview is...
func GetoneBrandforview(w http.ResponseWriter, r *http.Request) {
	brand := service.GetOneBrand(r)
	vehicles := service.GetParticlullarBrandVehicle(brand.ID)
	admintpl.ExecuteTemplate(w, "viewbrand.html", struct {
		Brand    model.Company
		Vehicles []model.Vehicle
	}{brand, vehicles})
}

//GetAllCustomer is....
func GetAllCustomer(w http.ResponseWriter, r *http.Request) {
	var message string
	var hasmessge bool

	if deletecustomer {
		hasmessge = true
		message = "Customer deleted successfully"
		deletecustomer = false
	}

	customer := service.GetAllCustomer(r)
	admintpl.ExecuteTemplate(w, "customerlist.html", struct {
		HasMessage bool
		Message    string
		Customers  []model.Customer
	}{hasmessge, message, customer})
}

//DeleteCustomer is...
func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	deletecustomer = true
	service.DeleteOneCustomer(r)
}

//GetoneCustomerforview is...
func GetoneCustomerforview(w http.ResponseWriter, r *http.Request) {
	customer := service.GetOneCustomer(r)
	admintpl.ExecuteTemplate(w, "viewcustomer.html", struct {
		Customer model.Customer
	}{customer})
}

//GetAllCustomerOrders is..
func GetAllCustomerOrders(w http.ResponseWriter, r *http.Request) {
	orders := service.GetAllTestDrives(r)
	admintpl.ExecuteTemplate(w, "order.html", struct {
		Orders []model.TestDrive
	}{orders})
}

//AdminRegister is..
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

//AdminRegisterPOST is...
func AdminRegisterPOST(w http.ResponseWriter, r *http.Request) {
	emailcheck, err := service.SaveAdmin(r)
	if err != nil {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
		return
	}

	if emailcheck {
		adminregisteremailerror = true
		http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
		return
	}
	//fmt.Fprintf(w, "Register data")
	adminregister = true
	http.Redirect(w, r, "/admin/register", http.StatusSeeOther)
}

//GetAdminAccountPage is..
func GetAdminAccountPage(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "username")
	email, _ := session.Values["username"]
	admin := service.GetOneAdminBYemail(email)
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
		Admin      model.SalesPerson
	}{hasmessge, message, admin})
}

//UpdateAdmin is..
func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	customer, err := service.AdminUpdate(r)
	if err != nil {
		log.Fatalln(err)
		return
	}
	updateadmin = true
	w.Header().Set("Content-Type", "application/json")
	w.Write(customer)
}
