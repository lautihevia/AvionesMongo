package repositories

import (
	"context"
	"fmt"

	"AvionesMongo/model"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AirplaneRepositoryInterface interface {
	GetAirplane() ([]model.Airplane, error)
	// ObtenerAulaPorID(id string) (model.Aula, error)
	// EliminarAula(id primitive.ObjectID) (*mongo.DeleteResult, error)
	InsertAirplane(aula model.Airplane) (*mongo.InsertOneResult, error)
	// ModificarAula(aula model.Aula) (*mongo.UpdateResult, error)
}

// El propósito de este código es definir un repositorio para manejar
// operaciones relacionadas con aviones (AirplaneRepository).
// La estructura AirplaneRepository contiene un campo db que representa
// la conexión a la base de datos. La función NewAirplaneRepository es un
// constructor que inicializa una nueva instancia de AirplaneRepository con
// una conexión a la base de datos proporcionada.

type AirplaneRepository struct {
	db DB
}

func NewAirplaneRepository(db DB) *AirplaneRepository {
	return &AirplaneRepository{
		db: db,
	}
}

func (repository AirplaneRepository) GetAirplane() ([]model.Airplane, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("airplanes")
	filtro := bson.M{}

	cursor, err := collection.Find(context.TODO(), filtro)

	defer cursor.Close(context.Background())

	// Itera a través de los resultados
	var airplanes []model.Airplane
	for cursor.Next(context.Background()) {
		var airplane model.Airplane
		err := cursor.Decode(&airplane)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
		}
		airplanes = append(airplanes, airplane)
	}

	return airplanes, err
}

// func (repository AulaRepository) ObtenerAulaPorID(id string) (model.Aula, error) {
// 	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")
// 	objectID := utils.GetObjectIDFromStringID(id)
// 	filtro := bson.M{"_id": objectID}

// 	cursor, err := collection.Find(context.TODO(), filtro)

// 	defer cursor.Close(context.Background())

// 	// Itera a través de los resultados
// 	var aula model.Aula

// 	for cursor.Next(context.Background()) {
// 		err := cursor.Decode(&aula)
// 		if err != nil {
// 			fmt.Printf("Error: %v\n", err)
// 		}
// 	}

// 	return aula, err
// }

func (repository AirplaneRepository) InsertAirplane(airplane model.Airplane) (*mongo.InsertOneResult, error) {
	collection := repository.db.GetClient().Database("ejemplo").Collection("airplanes")
	resultado, err := collection.InsertOne(context.TODO(), airplane)
	return resultado, err
}

// func (repository AulaRepository) ModificarAula(aula model.Aula) (*mongo.UpdateResult, error) {
// 	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")

// 	filtro := bson.M{"_id": aula.ID}
// 	entidad := bson.M{"$set": bson.M{"nombre": aula.Nombre}}

// 	resultado, err := collection.UpdateOne(context.TODO(), filtro, entidad)

// 	return resultado, err
// }

// func (repository AulaRepository) EliminarAula(id primitive.ObjectID) (*mongo.DeleteResult, error) {
// 	collection := repository.db.GetClient().Database("ejemplo").Collection("aulas")

// 	filtro := bson.M{"_id": id}

// 	resultado, err := collection.DeleteOne(context.TODO(), filtro)

// 	return resultado, err
// }
