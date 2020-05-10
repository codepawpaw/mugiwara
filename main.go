package main

import (
	"fmt"
	"net/http"
	"os"

	driver "./driver"

	ph "./handler/http"
	jwtService "./service/jwt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth"
)

func main() {
	dbName := os.Getenv("DB_NAME")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	connection, err := driver.ConnectSQL(dbHost, dbPort, "admin", dbPass, dbName)
	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)
	jwtServiceObj := jwtService.Init(tokenAuth)

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	authHandler := ph.InitAuthHandler(connection, jwtServiceObj)
	incidentHandler := ph.InitIncidentHandler(connection)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	r.Use(cors.Handler)

	r.Group(func(r chi.Router) {
		r.Route("/public", func(rt chi.Router) {
			rt.Route("/auth", func(route chi.Router) {
				route.Post("/login", authHandler.Login)
				route.Post("/register", authHandler.Register)
			})
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(jwtServiceObj.Verifier())
		r.Use(jwtServiceObj.Authenticator())

		r.Route("/v1", func(rt chi.Router) {
			rt.Route("/auth", func(route chi.Router) {
				route.Post("/update", authHandler.Update)
			})

			rt.Route("/incident", func(route chi.Router) {
				route.Get("/{city}", incidentHandler.GetByCity)
				route.Post("/", incidentHandler.Create)
				route.Put("/", incidentHandler.Update)
				route.Delete("/{city}/incident/{id:[0-9]+}", incidentHandler.Delete)
			})
		})
	})

	fmt.Println("Server listen at :8005")
	http.ListenAndServe(":8000", r)
}
