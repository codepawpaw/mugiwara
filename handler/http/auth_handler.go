package handler

import (
	"encoding/json"
	"net/http"

	driver "../../driver"
	models "../../models"
	repository "../../repository"
	authRepository "../../repository/auth"
	jwtService "../../service/jwt"
	"github.com/dgrijalva/jwt-go"
)

func InitAuthHandler(db *driver.DB, jwtServiceObj *jwtService.JwtService) *AuthHandler {
	return &AuthHandler{
		authRepository: authRepository.InitAuthRepository(db.SQL),
		jwtService:     jwtServiceObj,
	}
}

type AuthHandler struct {
	authRepository repository.AuthRepository
	jwtService     *jwtService.JwtService
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (authHandler *AuthHandler) Login(response http.ResponseWriter, request *http.Request) {
	userParam := models.User{}
	json.NewDecoder(request.Body).Decode(&userParam)

	generatedToken := authHandler.jwtService.Encode(userParam.Email)
	userParam.JWTToken = generatedToken

	userData, err := authHandler.authRepository.GetByEmail(request.Context(), userParam.Email)

	if err != nil {
		respondWithError(response, err)
	}

	if userData != nil {
		userJsonData, _ := json.Marshal(userData)
		userResponse := construct(userJsonData, err)
		respondwithJSON(response, userResponse.Status, userResponse)
	} else {
		userData, err := authHandler.authRepository.Create(request.Context(), &userParam)

		if err != nil {
			respondWithError(response, err)
		}

		userJsonData, _ := json.Marshal(userData)
		userResponse := construct(userJsonData, err)
		respondwithJSON(response, userResponse.Status, userResponse)
	}
}
