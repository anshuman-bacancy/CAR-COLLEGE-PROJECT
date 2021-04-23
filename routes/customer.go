package routes

import (
	"project/controller"
)

// customer frontend routes
func InitializeCustomerViewRoutes() {
	//GET ALL DATA
	r.HandleFunc("/Registration", controller.CustomerRegister).Methods("GET")
	r.HandleFunc("/Login", controller.CustomerLogin).Methods("GET")
	r.HandleFunc("/customer/Logout", controller.CustomerLogout).Methods("GET")
	r.HandleFunc("/customer/index", controller.AuthenticationCustomer(controller.CustomerIndexPage)).Methods("GET")
	r.HandleFunc("/customer/orders", controller.AuthenticationCustomer(controller.CustomerGetAllOrders)).Methods("GET")
	r.HandleFunc("/customer/compare", controller.AuthenticationCustomer(controller.CompareCar)).Methods("GET")

	//GET VIEW PAGE
	r.HandleFunc("/customer/brand/view/{id}", controller.AuthenticationCustomer(controller.GetAllVehicleWithBrandForView)).Methods("GET")
	r.HandleFunc("/customer/vehicle/view/{id}", controller.AuthenticationCustomer(controller.CustomerGetoneVehicleforview)).Methods("GET")
	r.HandleFunc("/customer/account", controller.AuthenticationCustomer(controller.CustomerAccountForView)).Methods("GET")
	r.HandleFunc("/customer/forgotpassword", controller.CustomerForgotPassword).Methods("GET")
	r.HandleFunc("/customer/setpassword/{id}", controller.CustomerSetForgotPasswordPage).Methods("GET")
	r.HandleFunc("/customer/getVehicle/{id}", controller.CustomerGetVehicle).Methods("GET")
	r.HandleFunc("/success", controller.CustomerSuccess).Methods("GET")

}

// customer backend routes
func InitializeCustomerBackendRoutes() {
	//POST REQUEST
	r.HandleFunc("/customer/register", controller.CustomerRegisterPOST).Methods("POST")
	r.HandleFunc("/customer/login", controller.CustomerLoginPost).Methods("POST")
	r.HandleFunc("/customer/validateemail", controller.CustomerValidateEmail).Methods("POST")
	r.HandleFunc("/customer/book/vehicle", controller.AuthenticationCustomer(controller.CustomerTestDrive)).Methods("POST")
	r.HandleFunc("/customer/bookTestDrive", controller.AuthenticationCustomer(controller.CustomerTestDrive)).Methods("POST")

	//PUT REQUEST
	r.HandleFunc("/customer/{id}", controller.CustomerUpdate).Methods("PUT")
}
