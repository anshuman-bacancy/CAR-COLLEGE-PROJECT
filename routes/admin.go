package routes

import (
	"net/http"
	"project/controller"
)

// admin view routes
func InitializeAdminViewRoutes() {
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("./static")))
	r.PathPrefix("/static/").Handler(fs)
	http.Handle("/static/", r)

	r.HandleFunc("/", controller.HomePage).Methods("GET")
	r.HandleFunc("/admin", controller.Login).Methods("GET")
	r.HandleFunc("/admin/logout", controller.Logout).Methods("GET")
	r.HandleFunc("/error", controller.ServerError).Methods("GET")
	r.HandleFunc("/admin/register", controller.AuthenticationAdmin(controller.AdminRegister)).Methods("GET")

	//GET ALL DATA
	r.HandleFunc("/admin/vehicle", controller.AuthenticationAdmin(controller.AdminIndexpageProcess)).Methods("GET")
	r.HandleFunc("/admin/brand", controller.AuthenticationAdmin(controller.GetAllBrand)).Methods("GET")
	r.HandleFunc("/admin/customer", controller.AuthenticationAdmin(controller.GetAllCustomer)).Methods("GET")
	r.HandleFunc("/admin/orders", controller.AuthenticationAdmin(controller.GetAllCustomerOrders)).Methods("GET")

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
	r.HandleFunc("/admin/account", controller.AuthenticationAdmin(controller.GetAdminAccountPage)).Methods("GET")
}

// admin backend routes
func InitializeAdminBackendRoutes() {
	//POST METHODS
	r.HandleFunc("/admin/login", controller.LoginPost).Methods("POST")
	r.HandleFunc("/admin/register", controller.AdminRegisterPOST).Methods("POST")
	r.HandleFunc("/admin/brand", controller.AuthenticationAdmin(controller.SaveBrand)).Methods("POST")
	r.HandleFunc("/admin/vehicle", controller.AuthenticationAdmin(controller.SaveVehicle)).Methods("POST")

	//PUT METHODS
	r.HandleFunc("/admin/vehicle/{id}", controller.AuthenticationAdmin(controller.UpdateVehicle)).Methods("PUT")
	r.HandleFunc("/admin/brand/{id}", controller.AuthenticationAdmin(controller.UpdateBrand)).Methods("PUT")
	r.HandleFunc("/admin/{id}", controller.AuthenticationAdmin(controller.UpdateAdmin)).Methods("PUT")
	r.HandleFunc("/admin/updateTestDrive/", controller.AuthenticationAdmin(controller.UpdateCustomerTestDriveStatus)).Methods("PUT")

	//DELETE METHODS
	r.HandleFunc("/admin/vehicle/{id}", controller.AuthenticationAdmin(controller.DeleteVehicle)).Methods("DELETE")
	r.HandleFunc("/admin/brand/{id}", controller.AuthenticationAdmin(controller.DeleteBrand)).Methods("DELETE")
	r.HandleFunc("/admin/customer/{id}", controller.AuthenticationAdmin(controller.DeleteCustomer)).Methods("DELETE")

	//NOT FOUND
	r.NotFoundHandler = http.HandlerFunc(controller.NotFound)
}
