package main

import (
	"project/common"
	"project/routes"
)

func main() {
	common.InitialMigration()
	routes.CreateRouter()
	routes.InitializeRoutesFrontendCustomer()
	routes.InitializeRoutesBackendCustomer()
	routes.InitializeRoutesBackendAdmin()
	routes.InitializeRoutesFrontendAdmin()
	routes.StartServer()
}
