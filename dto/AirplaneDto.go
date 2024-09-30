package dto

import (
	"AvionesMongo/model"
	"AvionesMongo/utils"
)

type Airplane struct {
	ID             string
	PassengerCount int
	Airline        string
}

// Sirve para Crear una nueva instancia del
// tipo dto.Airplane a partir de una
// instancia del tipo model.Airplane.

func NewAirplane(airplane model.Airplane) *Airplane {
	return &Airplane{
		ID:             utils.GetStringIDFromObjectID(airplane.ID),
		PassengerCount: airplane.PassengerCount,
		Airline:        airplane.Airline,
	}
}

// Convertir una instancia del tipo
// dto.Airplane a una instancia del
// tipo model.Airplane.

func (airplane Airplane) GetModel() model.Airplane {
	return model.Airplane{
		ID:             utils.GetObjectIDFromStringID(airplane.ID),
		PassengerCount: airplane.PassengerCount,
		Airline:        airplane.Airline,
	}
}

//Atencion, explicacion del GetModel
// En Go, la sintaxis (airplane Airplane) antes del
// nombre de la función GetModel indica que GetModel es
// un método con un receptor de tipo Airplane.
// Esto significa que GetModel es un método que puede
// ser llamado en cualquier instancia de Airplane.

// Diferencia entre Método y Función

// - Método: Un método es una función que tiene un receptor.
// El receptor es el tipo al que el método está asociado.
// En este caso, GetModel es un método del tipo Airplane.

// - Función: Una función no tiene un receptor y puede
// ser llamada independientemente de cualquier tipo.

// Se llama a un método en una instancia de un tipo
// a := Airplane{ID: "some-id", PassengerCount: 100}
// modelAirplane := a.GetModel()
// en este es instancia punto get model

// Si fuera una funcion, se llamaria asi:
// a := Airplane{ID: "some-id", PassengerCount: 100}
// modelAirplane := GetModel(a)
// en este es llamar a la funcion y pasarle la instancia
