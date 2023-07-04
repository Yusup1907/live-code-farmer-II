package model

import "time"

type PlantModel struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedBy string    `json:"updated_by"`
}

type CreatePlantRequest struct {
	Name      string `json:"name" binding:"required,min=3"`
	CreatedBy string `json:"created_by"`
}
