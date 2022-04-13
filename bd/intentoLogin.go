package bd

import (
	"github.com/hosnicolina/twitter-go/models"
	"golang.org/x/crypto/bcrypt"
)

/*IntentoLogin realiza el crequeo en la base de datos*/
func IntentoLogin(email, password string) (models.Usuario, bool) {
	usu, econtrado, _ := ChequeoYaExisteUsuario(email)

	if !econtrado {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true

}
