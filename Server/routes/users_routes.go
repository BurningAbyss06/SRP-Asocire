package routes

import (
	"encoding/json"
	"net/http"

	"github.com/BurningAbyss06/SRP-Asocire/Server/database"
	"github.com/BurningAbyss06/SRP-Asocire/Server/models"
	"github.com/BurningAbyss06/SRP-Asocire/Server/utilities"
	"github.com/gorilla/mux"
)

func Get_All_User(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	database.DBConnection().Find(&users)

	json.NewEncoder(w).Encode(&users)
}

func Get_One_User(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DBConnection().Where("Name=?", params["Name"]).First(&user)

	if user.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado"))
	}

	json.NewEncoder(w).Encode(&user)

}

func Post_User(w http.ResponseWriter, r *http.Request) {
	var newuser models.NewUser
	json.NewDecoder(r.Body).Decode(&newuser)
	salt := utilities.NewSalt()
	pas, err := utilities.HashPassword(newuser.Password, salt)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al momento de Crear el usuario"))
	}

	user := models.User{
		Name:     newuser.Name,
		Password: pas,
		Salt:     salt,
	}

	create_user := database.DBConnection().Create(&user)

	if err := create_user.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error al momento de Crear el usuario"))
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&user)

}

func Delete_User(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	database.DBConnection().Where("Name=?", params["Name"]).First(&user)

	if user.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Usuario no encontrado"))
	}

	deleteuser := database.DBConnection().Delete(&user)
	if err := deleteuser.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("No se pudo eliminar el usuario"))
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Se ha eliminado el usuario correctamente"))
	}

}
