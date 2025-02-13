package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r *mux.Router

// starts server at 8084
func StartServer() {
	log.Println("Server is started at http://localhost:8084")
	http.ListenAndServe(":8084", r)
}

// initializes mux router
func CreateRouter() {
	r = mux.NewRouter()
}
