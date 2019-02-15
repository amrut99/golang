package main

import (
	"router"

	"net/http"
)

func main() {
	router := router.InitRoutes()
	server := &http.Server {
		Handler : router,
		Addr : ":8080",
	}
	server.ListenAndServe()
}