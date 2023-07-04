package model

import "time"

type BillHeaderModel struct {
	ID          int                `json:"id"`
	FarmerID    int                `json:"farmer_id" binding:"required"`
	Name        string             `json:"name"`
	Address     string             `json:"address"`
	PhoneNumber string             `json:"phone_number"`
	Date        time.Time          `json:"date"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
	CreatedBy   string             `json:"created_by"`
	UpdatedBy   string             `json:"updated_by"`
	Total       float64            `json:"total"`
	ArrDetails  []*BillDetailModel `json:"arr_details"`
}

type BillDetailModel struct {
	ID             int       `json:"id"`
	BillID         int       `json:"bill_id"`
	FertilizerID   int       `json:"fertilizer_id"`
	FertilizerName string    `json:"fertilizer_name"`
	Quantity       float64   `json:"quantity" `
	Price          float64   `json:"price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
	CreatedBy      string    `json:"created_by"`
	UpdatedBy      string    `json:"updated_by"`
	Total          float64   `json:"total"`
}

func (h *BillHeaderModel) CalculateTotal() {
	total := 0.0
	for _, detail := range h.ArrDetails {
		detail.Total = detail.Price * detail.Quantity
		total += detail.Total
	}
	h.Total = total
}
