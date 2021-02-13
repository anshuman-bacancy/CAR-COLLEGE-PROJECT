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
	fmt.Println("Server is started at 8082")
	http.ListenAndServe(":8082", r)
}

//CreateRouter is...
func CreateRouter() {
	r = mux.NewRouter()
}

//InitializeRoutesfrontend is..
func InitializeRoutesfrontend() {
	path := build.Default.GOPATH + "/src/project/static/"
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir(path)))
	r.PathPrefix("/static/").Handler(fs)

	http.Handle("/static/", r)
	r.HandleFunc("/", controller.HomePage).Methods("GET")
	r.HandleFunc("/admin", controller.Login).Methods("GET")
	r.HandleFunc("/admin/logout", controller.Logout).Methods("GET")
	r.HandleFunc("/error", controller.Authentication(controller.ServerError)).Methods("GET")
	r.HandleFunc("/admin/vehicle", controller.Authentication(controller.AdminIndexpageProcess)).Methods("GET")
	r.HandleFunc("/admin/create/vehicle", controller.Authentication(controller.CreateVehicleform)).Methods("GET")
	r.HandleFunc("/admin/vehicle/view/{id}", controller.Authentication(controller.GetoneVehicleforview)).Methods("GET")
	r.HandleFunc("/admin/vehicle/{id}", controller.Authentication(controller.GetoneVehicleforedit)).Methods("GET")

}

//InitializeRoutesbackend is...
func InitializeRoutesbackend() {

	r.HandleFunc("/admin/login", controller.LoginPost).Methods("POST")
	r.HandleFunc("/admin/vehicle", controller.Authentication(controller.SaveVehicle)).Methods("POST")
	r.HandleFunc("/admin/vehicle/{id}", controller.UpdateVehicle).Methods("PUT")
	r.HandleFunc("/admin/vehicle/{id}", controller.Authentication(controller.DeleteVehicle)).Methods("DELETE")
	r.NotFoundHandler = http.HandlerFunc(controller.NotFound)
}
