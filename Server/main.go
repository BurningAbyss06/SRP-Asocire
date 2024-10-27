package main

import (
	"log"
	"net/http"

	"github.com/BurningAbyss06/SRP-Asocire/Server/database"
	"github.com/BurningAbyss06/SRP-Asocire/Server/routes"
	"github.com/gorilla/mux"
)

func main() {
	DB := database.DBConnection()
	if DB.Error != nil {
		panic(DB.Error.Error())
	} else {
		log.Println("Se Conecto correctamente a la DB")
	}
	r := mux.NewRouter()
	routes.Setup_routes(r)

	http.ListenAndServe(":8000", r)
}
