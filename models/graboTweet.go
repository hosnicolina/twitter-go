package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*GraboTweet modelo del tweet*/
type GraboTweet struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	UserID  string             `bson:"userid" json:"userid,omitempty"`
	Mensaje string             `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time          `bson:"fecha" json:"fecha,omitempty"`
}
