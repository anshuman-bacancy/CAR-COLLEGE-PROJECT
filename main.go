package main

import (
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

	routes.StartServer()
}
