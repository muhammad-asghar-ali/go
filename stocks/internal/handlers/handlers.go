package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"stocks/internal/config"
	"stocks/internal/models"
)

func GetStocks(w http.ResponseWriter, r *http.Request) {
	repo := models.NewStockRepository(config.DB())
	stocks, err := repo.List()
	if err != nil {
		http.Error(w, "Error retrieving stocks", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stocks); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func GetStockByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	repo := models.NewStockRepository(config.DB())
	stock, err := repo.Get(id)
	if err != nil {
		http.Error(w, "Stock not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stock); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	repo := models.NewStockRepository(config.DB())
	if err := repo.Create(&stock); err != nil {
		http.Error(w, "Error creating stock", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(stock); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func UpdateStockByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	var stock models.Stock
	if err := json.NewDecoder(r.Body).Decode(&stock); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	stock.ID = id

	repo := models.NewStockRepository(config.DB())
	if err := repo.Update(&stock); err != nil {
		http.Error(w, "Error updating stock", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(stock); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
	}
}

func DeleteStockByID(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid stock ID", http.StatusBadRequest)
		return
	}

	repo := models.NewStockRepository(config.DB())
	if err := repo.Delete(id); err != nil {
		http.Error(w, "Error deleting stock", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
