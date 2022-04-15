package models

/*Tweet Model de tweet */
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
