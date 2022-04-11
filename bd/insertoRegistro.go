package bd

import (
	"context"
	"time"

	"github.com/hosnicolina/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*InsertoRegistro insetar el registro de usuario */
func InsertoRegistro(user models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoCN.Database("twitor")
	col := db.Collection("usuarios")

	user.Password, _ = EncriptarPassword(user.Password)

	result, err := col.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil

}
