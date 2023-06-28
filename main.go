package main

import (
	"github.com/feebrian/gotodolist/config"
	"github.com/feebrian/gotodolist/router"
)

func main() {

	config.ConnectDB()   // to run automigration
	r := router.Router() // call router

	err := r.Run(":8080") // start development server
	if err != nil {
		panic(err)
	}

}
