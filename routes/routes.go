package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var r *mux.Router

// starts server at 8084
func StartServer(port string) {
	log.Println("Server is started at http://localhost"+port)
	http.ListenAndServe(port, r)
}

// initializes mux router
func CreateRouter() {
	r = mux.NewRouter()
}
