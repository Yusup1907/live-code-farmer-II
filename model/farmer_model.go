package model

import "time"

type FarmerModel struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedBy   string    `json:"created_by"`
	UpdatedBy   string    `json:"updated_by"`
}

type CreateFarmerRequest struct {
	Name        string `json:"name" binding:"required,min=3"`
	Address     string `json:"address" binding:"required"`
	PhoneNumber string `json:"phone_number" binding:"numeric"`
	CreatedBy   string `json:"created_by"`
}
