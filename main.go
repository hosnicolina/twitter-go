package main

import (
	"log"

	"github.com/hosnicolina/twitter-go/bd"
	"github.com/hosnicolina/twitter-go/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
