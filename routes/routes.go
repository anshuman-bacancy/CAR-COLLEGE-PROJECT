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
	r.HandleFunc("/customer/index", controller.AuthenticationCustomer(controller.CustomerIndexPage)).Methods("GET")
}

//InitializeRoutesbackendCustomer is...
func InitializeRoutesbackendCustomer() {
	r.HandleFunc("/customer/register", controller.CustomerRegisterPOST).Methods("POST")
	r.HandleFunc("/customer/login", controller.CustomerLoginPost).Methods("POST")
}

//InitializeRoutesfrontendAdmin is..
func InitializeRoutesfrontendAdmin() {
	path := build.Default.GOPATH + "/src/project/static/"
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(path)))
	r.PathPrefix("/static/").Handler(fs)

	http.Handle("/static/", r)
	r.HandleFunc("/", controller.HomePage).Methods("GET")
	r.HandleFunc("/admin", controller.Login).Methods("GET")
	r.HandleFunc("/admin/logout", controller.Logout).Methods("GET")
	r.HandleFunc("/error", controller.AuthenticationAdmin(controller.ServerError)).Methods("GET")
	r.HandleFunc("/admin/vehicle", controller.AuthenticationAdmin(controller.AdminIndexpageProcess)).Methods("GET")
	r.HandleFunc("/admin/create/vehicle", controller.AuthenticationAdmin(controller.CreateVehicleform)).Methods("GET")
	r.HandleFunc("/admin/vehicle/view/{id}", controller.AuthenticationAdmin(controller.GetoneVehicleforview)).Methods("GET")
	r.HandleFunc("/admin/vehicle/{id}", controller.AuthenticationAdmin(controller.GetoneVehicleforedit)).Methods("GET")
	r.HandleFunc("/admin/logo", controller.AuthenticationAdmin(controller.CreateLogoCompany)).Methods("GET")
}

//InitializeRoutesbackendAdmin is...
func InitializeRoutesbackendAdmin() {
	r.HandleFunc("/admin/login", controller.LoginPost).Methods("POST")
	r.HandleFunc("/admin/logo", controller.CreateLogoPost).Methods("POST")
	r.HandleFunc("/admin/vehicle", controller.AuthenticationAdmin(controller.SaveVehicle)).Methods("POST")
	r.HandleFunc("/admin/vehicle/{id}", controller.AuthenticationAdmin(controller.UpdateVehicle)).Methods("PUT")
	r.HandleFunc("/admin/vehicle/{id}", controller.AuthenticationAdmin(controller.DeleteVehicle)).Methods("DELETE")
	r.NotFoundHandler = http.HandlerFunc(controller.NotFound)
}
