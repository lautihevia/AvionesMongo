package handlers

import (
	"net/http"

	"AvionesMongo/dto"
	"AvionesMongo/services"

	"github.com/gin-gonic/gin"
)

type AirplaneHandler struct {
	airplaneService services.AirplaneInterface
}

func NewAirplaneHandler(airplaneService services.AirplaneInterface) *AirplaneHandler {
	return &AirplaneHandler{
		airplaneService: airplaneService,
	}
}

func (handler *AirplaneHandler) GetAirplanes(c *gin.Context) {
	airplanes := handler.airplaneService.GetAirplanes()

	c.JSON(http.StatusOK, airplanes)
}

// func (handler *AulaHandler) ObtenerAulaPorID(c *gin.Context) {
// 	id := c.Param("id")
// 	aulas := handler.aulaService.ObtenerAulaPorID(id)

// 	c.JSON(http.StatusOK, aulas)
// }

func (handler *AirplaneHandler) InsertAirplane(c *gin.Context) {
	var airplane dto.Airplane

	if err := c.ShouldBindJSON(&airplane); err != nil {
		// Si hay un error en el JSON, devuelve un error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resultado := handler.airplaneService.InsertAirplane(&airplane)

	c.JSON(http.StatusCreated, resultado)
}

// func (handler *AulaHandler) ModificarAula(c *gin.Context) {
// 	var aula dto.Aula

// 	if err := c.ShouldBindJSON(&aula); err != nil {
// 		// Si hay un error en el JSON, devuelve un error
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	aula.ID = c.Param("id")
// 	resultado := handler.aulaService.ModificarAula(&aula)

// 	c.JSON(http.StatusCreated, resultado)
// }

// func (handler *AulaHandler) EliminarAula(c *gin.Context) {
// 	id := c.Param("id")
// 	aulas := handler.aulaService.EliminarAula(id)

// 	c.JSON(http.StatusOK, aulas)
// }
