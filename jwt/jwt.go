package jwt

import (
	"time"

	jwt "github.com/golang-jwt/jwt"
	"github.com/hosnicolina/twitter-go/models"
)

/*GeneroJWT genera un jwt */
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("HosniBoniekColinaRodriguez")

	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"fecha_nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioweb":         t.Apellidos,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	signedToken, err := token.SignedString(miClave)

	if err != nil {
		return "", err
	}

	return signedToken, nil

}
