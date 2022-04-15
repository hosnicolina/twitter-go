package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hosnicolina/twitter-go/bd"
)

/*LeoTweets ruta para devolver los tweets */
func LeoTweets(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro id", http.StatusBadRequest)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "debe enviar el parametro pagina con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pag := int64(pagina)

	tweets, ok := bd.LeoTweets(IDUsuario, pag)

	if !ok {
		http.Error(w, "error al procesar los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweets)

}
