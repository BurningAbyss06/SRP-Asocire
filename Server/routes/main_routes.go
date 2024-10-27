package routes

import "github.com/gorilla/mux"

func Setup_routes(routes *mux.Router) {
	routes.HandleFunc("/", Hola)
}
