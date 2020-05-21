package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	connection "../../connection"
	driver "../../driver"
	dto "../../dto"
	models "../../models"
	repository "../../repository"
	incidentRepository "../../repository/incident"
	"github.com/go-chi/chi"
)

func InitIncidentHandler(db *driver.DB) *IncidentHandler {
	return &IncidentHandler{
		repository: incidentRepository.InitIncidentRepository(db.SQL),
	}
}

type IncidentHandler struct {
	repository repository.IncidentRepository
}

func (incidentHandler *IncidentHandler) Create(w http.ResponseWriter, r *http.Request) {
	incident := models.Incident{}
	json.NewDecoder(r.Body).Decode(&incident)

	incidentData, err := incidentHandler.repository.Create(r.Context(), &incident)

	if err != nil {
		respondWithError(w, err)
	}

	incidentResponse, _ := json.Marshal(incidentData)

	if err == nil {
		redisResponse := &dto.RedisResponse{
			Type:   "create",
			Data:   string(incidentResponse),
			Status: http.StatusOK,
		}

		redisResponseMarshaled, _ := json.Marshal(redisResponse)
		connection.GetRedis().Publish("city#"+incidentData.CityName, string(redisResponseMarshaled))
	}

	response := construct(incidentResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (incidentHandler *IncidentHandler) Update(w http.ResponseWriter, r *http.Request) {
	incident := models.Incident{}
	json.NewDecoder(r.Body).Decode(&incident)
	payload, err := incidentHandler.repository.Update(r.Context(), &incident)

	if err != nil {
		respondWithError(w, err)
	}

	incidentResponse, _ := json.Marshal(payload)

	if err == nil {
		redisResponse := &dto.RedisResponse{
			Type:   "update",
			Data:   string(incidentResponse),
			Status: http.StatusOK,
		}

		redisResponseMarshaled, _ := json.Marshal(redisResponse)
		connection.GetRedis().Publish("city#"+payload.CityName, string(redisResponseMarshaled))
	}

	response := construct(incidentResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (incidentHandler *IncidentHandler) GetByCity(w http.ResponseWriter, r *http.Request) {
	city := chi.URLParam(r, "city")

	payload, err := incidentHandler.repository.GetByCityName(r.Context(), string(city))

	if err != nil {
		respondWithError(w, err)
	}

	incidentResponse, _ := json.Marshal(payload)

	response := construct(incidentResponse, err)

	respondwithJSON(w, response.Status, response)
}

func (incidentHandler *IncidentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	cityName := chi.URLParam(r, "city")

	payload, err := incidentHandler.repository.Delete(r.Context(), int64(id))

	if err != nil {
		respondWithError(w, err)
	}

	if payload == true {
		incidentData := models.Incident{
			ID: int64(id),
		}

		data, _ := json.Marshal(incidentData)

		redisResponse := &dto.RedisResponse{
			Type:   "delete",
			Data:   string(data),
			Status: http.StatusOK,
		}

		redisResponseMarshaled, _ := json.Marshal(redisResponse)
		connection.GetRedis().Publish("city#"+cityName, string(redisResponseMarshaled))
	}

	incidentResponse, _ := json.Marshal(payload)

	response := construct(incidentResponse, err)

	respondwithJSON(w, response.Status, response)

}
