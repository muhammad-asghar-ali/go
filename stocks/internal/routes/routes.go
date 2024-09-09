package routes

import (
	"stocks/internal/handlers"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/stocks", handlers.GetStocks).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks", handlers.CreateStock).Methods("POST", "OPTIONS")
	router.HandleFunc("/stocks/{id:[0-9]+}", handlers.GetStockByID).Methods("GET", "OPTIONS")
	router.HandleFunc("/stocks/{id:[0-9]+}", handlers.UpdateStockByID).Methods("PUT", "OPTIONS")
	router.HandleFunc("/stocks/{id:[0-9]+}", handlers.DeleteStockByID).Methods("DELETE", "OPTIONS")

	return router
}
