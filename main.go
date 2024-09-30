package main

import (
	"AvionesMongo/handlers"
	"AvionesMongo/repositories"
	"AvionesMongo/services"

	"github.com/gin-gonic/gin"
)

var (
	airplaneHandler *handlers.AirplaneHandler
	router          *gin.Engine
)

func main() {
	router = gin.Default()
	//Iniciar objetos de handler
	dependencies()
	//Iniciar rutas
	mappingRoutes()

	router.Run(":8080")
}

func mappingRoutes() {
	//Listado de rutas
	router.GET("/airplanes", airplaneHandler.GetAirplanes)
	// router.GET("/aulas/:id", aulaHandler.ObtenerAulaPorID)
	router.POST("/airplanes", airplaneHandler.InsertAirplane)
	// router.PUT("/aulas/:id", aulaHandler.ModificarAula)
	// router.DELETE("/aulas/:id", aulaHandler.EliminarAula)
}

// Generacion de los objetos que se van a usar en la api
func dependencies() {
	//Definicion de variables de interface
	var database repositories.DB
	var airplaneRepository repositories.AirplaneRepositoryInterface
	var airplaneService services.AirplaneInterface

	//Creamos los objetos reales y los pasamos como parametro
	database = repositories.NewMongoDB()
	airplaneRepository = repositories.NewAirplaneRepository(database)
	airplaneService = services.NewAirplaneService(airplaneRepository)
	airplaneHandler = handlers.NewAirplaneHandler(airplaneService)
}
