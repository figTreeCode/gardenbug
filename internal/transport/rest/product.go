package rest

import (
	"encoding/json"
	"gardenbug/internal/services"
	"net/http"

	"github.com/jackc/pgx/v5"
)

type DatabaseHandler struct {
	conn *pgx.Conn
}

func (d *DatabaseHandler) getProducts(w http.ResponseWriter, r *http.Request) {
	products, err := services.GetActiveProducts(d.conn)
	if err != nil {
		http.Error(w, "Failed to fetch all active products", http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "applicaton/json")
	json.NewEncoder(w).Encode(products)
}
