package handler

import (
	"live-code-farmer-II/usecase"

	"github.com/gin-gonic/gin"
)

type BillHandler interface {
}

type billHandler struct {
	billUsecase usecase.BillUsecase
}

func NewBillHandler(s *gin.Engine, billUsecase usecase.BillUsecase) *BillHandler {
	billHandler := &billHandler{
		billUsecase: billUsecase,
	}
	// s.GET("/bills", billHandler.GetAllBills)
	// s.GET("/bills/:id", billHandler.GetBillByID)
	// s.GET("/bills/today", billHandler.GetTotalIncomeToday)
	// s.POST("/bills/", billHandler.CreateBill)
	return billHandler
}
