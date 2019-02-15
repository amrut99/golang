package router

import (
	
	"github.com/gorilla/mux"

	"mapstore"
	"controller"
)


func InitRoutes() *mux.Router{

	// Create instance of controler types
	controller := controller.CustomerController {
		Store : mapstore.NewMapStore(),
	}
	router := mux.NewRouter()
	router.HandleFunc("/customers", controller.Create)
	router.HandleFunc("/get", controller.PrintAll)
	return router
}