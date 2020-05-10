package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	jwtAuth "github.com/go-chi/jwtauth"
)

var secret = []byte("secret")

func Init(tokenAuth *jwtAuth.JWTAuth) *JwtService {
	return &JwtService{
		tokenAuth: tokenAuth,
	}
}

type JwtService struct {
	tokenAuth *jwtAuth.JWTAuth
}

func (jwtService *JwtService) Encode(username string) string {
	claim := jwt.MapClaims{"username": username}
	_, token, _ := jwtService.tokenAuth.Encode(claim)

	return token
}

func (jwtService *JwtService) Verifier() func(http.Handler) http.Handler {
	return jwtAuth.Verifier(jwtService.tokenAuth)
}

func (JwtService *JwtService) Authenticator() func(http.Handler) http.Handler {
	return jwtAuth.Authenticator
}
