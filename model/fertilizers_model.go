package model

import "time"

type FertilizersModel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Stok      int       `json:"stok"`
	Price     *int      `json:"price"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

type CreatefertilizersRequest struct {
	Name      string `json:"name" binding:"required,min=3"`
	Stok      int    `json:"stok" binding:"required"`
	Price     *int   `json:"price" binding:"required"`
	IsActive  bool   `json:"is_active"`
	CreatedBy string `json:"created_by"`
}
