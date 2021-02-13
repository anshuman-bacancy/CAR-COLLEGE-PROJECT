package main

import (
	"project/common"
	"project/routes"
)

func main() {
	common.Initialmigration()
	routes.CreateRouter()
	routes.InitializeRoutesfrontendCustomer()
	routes.InitializeRoutesbackendCustomer()
	routes.InitializeRoutesbackendAdmin()
	routes.InitializeRoutesfrontendAdmin()
	routes.StartServer()
}
