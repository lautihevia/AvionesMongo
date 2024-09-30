package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

//este archivo se llama ids porque es para los IDs de los objetos

//primitive.ObjectID es un tipo de dato definido en el paquete
//go.mongodb.org/mongo-driver/bson/primitive de la biblioteca
//del controlador de MongoDB para Go. Representa un identificador
//único de objeto (ObjectID) utilizado por MongoDB para
//identificar de manera única los documentos dentro de una colección.

//Las funciones son de conversion de string
//a primitive.ObjectID y viceversa.

func GetObjectIDFromStringID(id string) primitive.ObjectID {
	objectID, _ := primitive.ObjectIDFromHex(id)

	return objectID
}

func GetStringIDFromObjectID(id primitive.ObjectID) string {
	return id.Hex()
}
