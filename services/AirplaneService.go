package services

import (
	"AvionesMongo/dto"
	"AvionesMongo/repositories"
)

type AirplaneInterface interface {
	GetAirplanes() []*dto.Airplane
	// ObtenerAulaPorID(id string) *dto.Aula
	// EliminarAula(id string) bool
	InsertAirplane(aula *dto.Airplane) bool
	// ModificarAula(aula *dto.Aula) bool
}

type AirplaneService struct {
	airplaneRepository repositories.AirplaneRepositoryInterface
}

func NewAirplaneService(airplaneRepository repositories.AirplaneRepositoryInterface) *AirplaneService {
	return &AirplaneService{
		airplaneRepository: airplaneRepository,
	}
}

func (service *AirplaneService) GetAirplanes() []*dto.Airplane {
	airplanesDB, _ := service.airplaneRepository.GetAirplane()
	var airplanes []*dto.Airplane
	for _, airplaneDB := range airplanesDB {
		airplane := dto.NewAirplane(airplaneDB)
		airplanes = append(airplanes, airplane)
	}

	return airplanes
}

// func (service *AulaService) ObtenerAulaPorID(id string) *dto.Aula {
// 	aulaDB, err := service.aulaRepository.ObtenerAulaPorID(id)

// 	var aula *dto.Aula
// 	if err == nil {
// 		aula = dto.NewAula(aulaDB)
// 	}

// 	return aula
// }

func (service *AirplaneService) InsertAirplane(airplane *dto.Airplane) bool {
	service.airplaneRepository.InsertAirplane(airplane.GetModel())

	return true
}

// func (service *AulaService) ModificarAula(aula *dto.Aula) bool {
// 	service.aulaRepository.ModificarAula(aula.GetModel())

// 	return true
// }

// func (service *AulaService) EliminarAula(id string) bool {
// 	service.aulaRepository.EliminarAula(utils.GetObjectIDFromStringID(id))

// 	return true
// }
