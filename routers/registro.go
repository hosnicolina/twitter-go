package routers

import (
	"encoding/json"
	"net/http"

	"github.com/hosnicolina/twitter-go/bd"
	"github.com/hosnicolina/twitter-go/models"
)

/*Registro funcion para crear el registro de la base de datos*/
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email es un campo obligatorio", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe tener un minimo 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)

	if encontrado {
		http.Error(w, "El usuario ya se encuentra registrado", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar realizar el registro de usuario "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
