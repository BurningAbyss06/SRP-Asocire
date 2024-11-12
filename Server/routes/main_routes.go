package routes

import "github.com/gorilla/mux"

func Setup_routes(routes *mux.Router) {
	routes.HandleFunc("/", Hola)

	routes.HandleFunc("/user/get", Get_All_User).Methods("GET")
	routes.HandleFunc("/user/get/{Name}", Get_One_User).Methods("GET")
	routes.HandleFunc("/user/create", Post_User).Methods("POST")
	routes.HandleFunc("/user/delete/{Name}", Delete_User).Methods("DELETE")
}
