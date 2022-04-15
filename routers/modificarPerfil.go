package routers

import (
	"encoding/json"
	"net/http"

	"github.com/hosnicolina/twitter-go/bd"
	"github.com/hosnicolina/twitter-go/models"
)

/*ModificarPerfil Actualiza el perfil del usuario */
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "error al actualizar perfil "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool
	status, err = bd.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "error al actualizar perfil "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "error al actualizar perfil del usuario", 400)
		return
	}

	w.Header().Set("Context-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "message": "perfil actualizado con exito"})

}
