package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//ejemplo del maxi
// type Aula struct {
// 	ID     primitive.ObjectID `bson:"_id,omitempty"` //omitempty agrega el campo solo si no es nil
// 	Nombre string             `bson:"nombre"`
// }

// Airplane es la estructura del modelo de avión.
type Airplane struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	Airline        string             `bson:"airline"`
	PassengerCount int                `bson:"passengerCount"`
	CreationDate   time.Time          `bson:"creationDate"`
	UpdateDate     time.Time          `bson:"updateDate"`
}
