package routes

import (
	"github.com/gorilla/mux"

	"stocks/internal/handlers"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/stocks", handlers.GetStocks).Methods("GET", "OPTION")
	router.HandleFunc("/stocks", handlers.CreateStock).Methods("POST", "OPTION")
	router.HandleFunc("/stocks/{id}", handlers.GetStockByID).Methods("GET", "OPTION")
	router.HandleFunc("/stocks/{id}", handlers.UpdateStockByID).Methods("PUT", "OPTION")
	router.HandleFunc("/stocks/{id}", handlers.DeleteStockByID).Methods("DELETE", "OPTION")

	return router
}
