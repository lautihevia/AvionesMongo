package repositories

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// El propósito de esta estructura es
// encapsular la conexión al cliente de
// MongoDB dentro de un tipo más manejable
// (MongoDB). Esto puede ser útil para organizar
// el código y pasar la conexión a MongoDB a diferentes
// partes de la aplicación de manera más sencilla.

type MongoDB struct {
	Client *mongo.Client
}

// Crear y devolver una nueva instancia de la estructura MongoDB
// y establecer una conexión con la base de datos MongoDB

func NewMongoDB() *MongoDB {
	instance := &MongoDB{}
	instance.Connect()

	return instance
}

// (mongoDB *MongoDB): Indica que GetClient es un método que p
// uede ser llamado en cualquier instancia de MongoDB.
// La función GetClient es un método de la estructura MongoDB que
// devuelve el cliente de MongoDB almacenado en el campo Client
// de la estructura.

func (mongoDB *MongoDB) GetClient() *mongo.Client {
	return mongoDB.Client
}

// La función Connect establece una conexión a una base de datos
// MongoDB y asigna el cliente conectado al campo Client de la
// estructura MongoDB.
// La dejamos privada, se ejecuta cuando se crea el objeto
func (mongoDB *MongoDB) Connect() error {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		return err
	}

	// Se verifica la conexión al servidor MongoDB enviando
	// un comando ping.
	// Si el ping falla, se devuelve el error
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return err
	}

	mongoDB.Client = client

	return nil
}

// La función Disconnect cierra la conexiona a la base de datos
func (mongoDB *MongoDB) Disconnect() error {
	return mongoDB.Client.Disconnect(context.Background())
}
