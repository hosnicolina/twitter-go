package routers

import (
	"errors"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/hosnicolina/twitter-go/bd"
	"github.com/hosnicolina/twitter-go/models"
)

var (
	Email     string
	IDUsuario string
)

/*ProcesoToken validar token*/
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("HosniBoniekColinaRodriguez")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, "", errors.New("el token es invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)

		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return claims, false, "", errors.New("el token es invalido")
	}

	return claims, false, "", err

}
