package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/hosnicolina/twitter-go/bd"
	"github.com/hosnicolina/twitter-go/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "datos invalidos", http.StatusBadRequest)
		return
	}

	if len(mensaje.Mensaje) > 140 {
		http.Error(w, "el tweet no puede exceder los 140 caracteres de largo", http.StatusBadRequest)
		return
	}

	tweet := models.GraboTweet{
		Mensaje: mensaje.Mensaje,
		UserID:  IDUsuario,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsetoTweet(tweet)

	if err != nil {
		http.Error(w, "no se pudo crear el tweet", 400)
		return
	}

	if !status {
		http.Error(w, "no se pudo crear el tweet intente nuevamente", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":  true,
		"message": "tweet creado con exito",
	})

}
