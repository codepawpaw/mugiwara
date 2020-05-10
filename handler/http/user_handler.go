package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	repository "../../repository"
	userRepository "../../repository/user"
	"github.com/go-chi/chi"
)

func InitUserHandler(db *driver.DB) *UserHandler {
	return &UserHandler{
		repository: userRepository.InitUserRepository(db.SQL),
	}
}

type UserHandler struct {
	repository repository.UserRepository
}

func (userHandler *UserHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := userHandler.repository.GetByID(r.Context(), int64(id))

	userResponse, _ := json.Marshal(payload)

	response := construct(userResponse, err)

	respondwithJSON(w, response.Status, response)
}
