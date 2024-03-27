package rest

import (
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
)

func SetupRoutes(r chi.Router, conn *pgx.Conn) {
	r.Get("/products", (&DatabaseHandler{conn}).getProducts)
}
