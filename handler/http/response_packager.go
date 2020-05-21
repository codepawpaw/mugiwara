package handler

import (
	"encoding/json"
	"net/http"

	dto "../../dto"
)

func construct(data []byte, err error) *dto.HttpResponse {
	httpStatus := http.StatusOK
	errorMessage := ""
	responseData := string(data)

	if err != nil {
		responseData = ""
		httpStatus = http.StatusInternalServerError
		errorMessage = err.Error()
	}

	response := &dto.HttpResponse{
		Data:         responseData,
		ErrorMessage: errorMessage,
		Status:       httpStatus,
	}

	return response
}

func constructResponse(data []byte, err string) *dto.HttpResponse {
	httpStatus := http.StatusOK
	errorMessage := err
	responseData := string(data)

	if err != "" {
		responseData = ""
		httpStatus = http.StatusInternalServerError
	}

	response := &dto.HttpResponse{
		Data:         responseData,
		ErrorMessage: errorMessage,
		Status:       httpStatus,
	}

	return response
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(response http.ResponseWriter, err error) {
	jsonResponse := construct(nil, err)
	respondwithJSON(response, jsonResponse.Status, jsonResponse)
}
