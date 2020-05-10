package handler

import (
	"encoding/json"
	"net/http"

	driver "../../driver"
	dto "../../dto"
	models "../../models"
	repository "../../repository"
	accountRepository "../../repository/account"
	authRepository "../../repository/auth"
	userRepository "../../repository/user"
	jwtService "../../service/jwt"
	"github.com/dgrijalva/jwt-go"
)

func InitAuthHandler(db *driver.DB, jwtServiceObj *jwtService.JwtService) *AuthHandler {
	return &AuthHandler{
		accountRepository: accountRepository.InitAccountRepository(db.SQL),
		authRepository:    authRepository.InitAuthRepository(db.SQL),
		userRepository:    userRepository.InitUserRepository(db.SQL),
		jwtService:        jwtServiceObj,
	}
}

type AuthHandler struct {
	accountRepository repository.AccountRepository
	authRepository    repository.AuthRepository
	userRepository    repository.UserRepository
	jwtService        *jwtService.JwtService
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func (authHandler *AuthHandler) Login(response http.ResponseWriter, request *http.Request) {
	account := models.Account{}
	json.NewDecoder(request.Body).Decode(&account)

	accountResponse, err := authHandler.accountRepository.GetByUsernameAndPassword(request.Context(), string(account.Username), string(account.Password))

	emptyAuthResponse := dto.Auth{}
	authResponseJson, _ := json.Marshal(emptyAuthResponse)
	errorMessage := ""

	if accountResponse == nil || err != nil {
		errorMessage = "Incorrect username or password"
	} else {
		userReponse, err := authHandler.userRepository.GetByID(request.Context(), accountResponse.UserId)

		if err != nil {
			errorMessage = err.Error()
		}

		generatedToken := authHandler.jwtService.Encode(accountResponse.Username)

		authResponse := dto.Auth{
			Account: accountResponse,
			User:    userReponse,
			Token:   generatedToken,
		}

		json, _ := json.Marshal(authResponse)
		authResponseJson = json
	}

	responseJson := constructResponse(authResponseJson, errorMessage)
	respondwithJSON(response, responseJson.Status, responseJson)
}

func (authHandler *AuthHandler) Register(response http.ResponseWriter, request *http.Request) {
	authModel := models.Auth{}
	json.NewDecoder(request.Body).Decode(&authModel)

	if authModel.Account.Username == "" {
		respondwithJSON(response, 400, "account.username cannot be empty")
	} else {

		user := authModel.User
		account := authModel.Account

		authResponse, err := authHandler.authRepository.Create(request.Context(), &user, &account)

		generatedToken := authHandler.jwtService.Encode(account.Username)

		authResponse.Token = generatedToken

		authResponseJson, _ := json.Marshal(authResponse)

		responseJson := construct(authResponseJson, err)

		respondwithJSON(response, responseJson.Status, responseJson)
	}
}

func (authHandler *AuthHandler) Update(response http.ResponseWriter, request *http.Request) {
	authModel := models.Auth{}
	json.NewDecoder(request.Body).Decode(&authModel)

	user := authModel.User
	account := authModel.Account

	authResponse, err := authHandler.authRepository.Update(request.Context(), &user, &account)

	authResponseJson, _ := json.Marshal(authResponse)

	responseJson := construct(authResponseJson, err)

	respondwithJSON(response, responseJson.Status, responseJson)
}
