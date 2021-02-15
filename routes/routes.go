package routes

import (
	"fmt"
	"go/build"
	"net/http"
	"project/controller"

	"github.com/gorilla/mux"
)

var r *mux.Router

//StartServer is started at 8082
func StartServer() {
	fmt.Println("Server is started at http://localhost:8083")
	http.ListenAndServe(":8083", r)
}

//CreateRouter is...
func CreateRouter() {
	r = mux.NewRouter()
}

//InitializeRoutesfrontendCustomer is..
func InitializeRoutesfrontendCustomer() {
	r.HandleFunc("/Registration", controller.CustomerRegister).Methods("GET")
	r.HandleFunc("/Login", controller.CustomerLogin).Methods("GET")
	r.HandleFunc("/customer/Logout", controller.CustomerLogout).Methods("GET")
	//GET ALL DATA
	r.HandleFunc("/customer/index", controller.AuthenticationCustomer(controller.CustomerIndexPage)).Methods("GET")

	//GET VIEW PAGE
	r.HandleFunc("/customer/brand/view/{id}", controller.AuthenticationCustomer(controller.GetallVehicleWithBrandforview)).Methods("GET")
	r.HandleFunc("/customer/vehicle/view/{id}", controller.AuthenticationCustomer(controller.CustomerGetoneVehicleforview)).Methods("GET")
	r.HandleFunc("/customer/account", controller.AuthenticationCustomer(controller.CustomerAccountforview)).Methods("GET")
	r.HandleFunc("/customer/forgotpassword", controller.CustomerForgotPassword).Methods("GET")
	r.HandleFunc("/customer/setpassword/{id}", controller.CustomerSetForgotPasswordPage).Methods("GET")
	r.HandleFunc("/success", controller.CustomerSuccess).Methods("GET")
}

//InitializeRoutesbackendCustomer is...
func InitializeRoutesbackendCustomer() {
	//POST REQUEST
	r.HandleFunc("/customer/register", controller.CustomerRegisterPOST).Methods("POST")
	r.HandleFunc("/customer/login", controller.CustomerLoginPost).Methods("POST")
	r.HandleFunc("/customer/validateemail", controller.CustomerValidateEmail).Methods("POST")
	//PUT REQUEST
	r.HandleFunc("/customer/{id}", controller.CustomerUpdate).Methods("PUT")
}

//InitializeRoutesfrontendAdmin is..
func InitializeRoutesfrontendAdmin() {
	//static
	path := build.Default.GOPATH + "/src/project/static/"
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(path)))
	r.PathPrefix("/static/").Handler(fs)
	http.Handle("/static/", r)

	r.HandleFunc("/", controller.HomePage).Methods("GET")
	r.HandleFunc("/admin", controller.Login).Methods("GET")
	r.HandleFunc("/admin/logout", controller.Logout).Methods("GET")
	r.HandleFunc("/error", controller.AuthenticationAdmin(controller.ServerError)).Methods("GET")

	//GET ALL DATA
	r.HandleFunc("/admin/vehicle", controller.AuthenticationAdmin(controller.AdminIndexpageProcess)).Methods("GET")
	r.HandleFunc("/admin/brand", controller.AuthenticationAdmin(controller.GetAllBrand)).Methods("GET")
	r.HandleFunc("/admin/customer", controller.AuthenticationAdmin(controller.GetAllCustomer)).Methods("GET")

	//GET CREATE PAGE
	r.HandleFunc("/admin/create/vehicle", controller.AuthenticationAdmin(controller.CreateVehicleform)).Methods("GET")
	r.HandleFunc("/admin/create/brand", controller.AuthenticationAdmin(controller.CreateBrandform)).Methods("GET")

	//GET EDIT PAGE
	r.HandleFunc("/admin/vehicle/{id}", controller.AuthenticationAdmin(controller.GetoneVehicleforedit)).Methods("GET")
	r.HandleFunc("/admin/brand/{id}", controller.AuthenticationAdmin(controller.GetoneBrandforedit)).Methods("GET")
	//GET VIEW PAGE
	r.HandleFunc("/admin/vehicle/view/{id}", controller.AuthenticationAdmin(controller.GetoneVehicleforview)).Methods("GET")
	r.HandleFunc("/admin/brand/view/{id}", controller.AuthenticationAdmin(controller.GetoneBrandforview)).Methods("GET")
	r.HandleFunc("/admin/customer/view/{id}", controller.AuthenticationAdmin(controller.GetoneCustomerforview)).Methods("GET")
}

//InitializeRoutesbackendAdmin is...
func InitializeRoutesbackendAdmin() {
	//POST METHODS
	r.HandleFunc("/admin/login", controller.LoginPost).Methods("POST")
	r.HandleFunc("/admin/brand", controller.AuthenticationAdmin(controller.SaveBrand)).Methods("POST")
	r.HandleFunc("/admin/vehicle", controller.AuthenticationAdmin(controller.SaveVehicle)).Methods("POST")

	//PUT METHODS
	r.HandleFunc("/admin/vehicle/{id}", controller.AuthenticationAdmin(controller.UpdateVehicle)).Methods("PUT")
	r.HandleFunc("/admin/brand/{id}", controller.AuthenticationAdmin(controller.UpdateBrand)).Methods("PUT")

	//DELETE METHODS
	r.HandleFunc("/admin/vehicle/{id}", controller.AuthenticationAdmin(controller.DeleteVehicle)).Methods("DELETE")
	r.HandleFunc("/admin/brand/{id}", controller.AuthenticationAdmin(controller.DeleteBrand)).Methods("DELETE")
	r.HandleFunc("/admin/customer/{id}", controller.AuthenticationAdmin(controller.DeleteCustomer)).Methods("DELETE")
	//NOT FOUND
	r.NotFoundHandler = http.HandlerFunc(controller.NotFound)
}
