package main

import (
	"os"
	"project/common"
	"project/routes"
)

func main() {
	common.InitialMigration()
	routes.CreateRouter()

	routes.InitializeCustomerViewRoutes()
	routes.InitializeCustomerBackendRoutes()

	routes.InitializeAdminViewRoutes()
	routes.InitializeAdminBackendRoutes()

	var port string

	if len(os.Args) == 2 {
		port = string(":" + os.Args[1])
	} else {
		port = ":8080"
	}

	routes.StartServer(port)
}
