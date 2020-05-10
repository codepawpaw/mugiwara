package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	driver "../../driver"
	repository "../../repository"
	accountRepository "../../repository/account"
	"github.com/go-chi/chi"
)

func InitAccountHandler(db *driver.DB) *AccountHandler {
	return &AccountHandler{
		repository: accountRepository.InitAccountRepository(db.SQL),
	}
}

type AccountHandler struct {
	repository repository.AccountRepository
}

func (accountHandler *AccountHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	payload, err := accountHandler.repository.GetByID(r.Context(), int64(id))

	accountResponse, _ := json.Marshal(payload)

	response := construct(accountResponse, err)

	respondwithJSON(w, response.Status, response)
}

func respondwithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondwithJSON(w, code, map[string]string{"message": msg})
}
