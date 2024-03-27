package main

import (
	"context"
	"gardenbug/internal/models"
	"gardenbug/internal/transport/rest"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	ctx := context.Background()

	conn, err := models.NewDatabaseConnection(ctx)

	router := chi.NewRouter()

	router.Route("/api", func(r chi.Router) {
		rest.SetupRoutes(r, conn)
	})

	http.ListenAndServe(":8080", router)
}
