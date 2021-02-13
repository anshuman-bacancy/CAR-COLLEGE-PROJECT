package main

import (
	"project/common"
	"project/routes"
)

func main() {
	common.Initialmigration()
	routes.CreateRouter()
	routes.InitializeRoutesbackend()
	routes.InitializeRoutesfrontend()
	routes.StartServer()
}
