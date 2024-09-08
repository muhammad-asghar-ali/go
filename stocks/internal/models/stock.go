package models

import (
	"time"

	"github.com/jmoiron/sqlx"
)

type (
	Stock struct {
		ID        int64     `json:"id"`
		Name      string    `json:"name"`
		Price     string    `json:"price"`
		Company   string    `json:"company"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	StockRepository struct {
		DB *sqlx.DB
	}
)

func NewStockRepository(db *sqlx.DB) *StockRepository {
	return &StockRepository{DB: db}
}

func (r *StockRepository) Create(stock *Stock) error {
	query := `INSERT INTO stocks (name, price, company) VALUES ($1, $2, $3) RETURNING id`
	err := r.DB.QueryRow(query, stock.Name, stock.Price, stock.Company).Scan(&stock.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *StockRepository) Get(id int64) (*Stock, error) {
	query := `SELECT id, name, price, company, created_at, updated_at FROM stocks WHERE id = $1`
	stock := &Stock{}
	err := r.DB.Get(stock, query, id)
	if err != nil {
		return nil, err
	}
	return stock, nil
}

func (r *StockRepository) Update(stock *Stock) error {
	query := `UPDATE stocks SET name = $1, price = $2, company = $3, updated_at = CURRENT_TIMESTAMP WHERE id = $4`
	_, err := r.DB.Exec(query, stock.Name, stock.Price, stock.Company, stock.ID)
	return err
}

func (r *StockRepository) Delete(id int64) error {
	query := `DELETE FROM stocks WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *StockRepository) List() ([]Stock, error) {
	query := `SELECT id, name, price, company, created_at, updated_at FROM stocks`
	var stocks []Stock
	err := r.DB.Select(&stocks, query)
	if err != nil {
		return nil, err
	}
	return stocks, nil
}
