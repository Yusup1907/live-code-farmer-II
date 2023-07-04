package handler

import (
	"live-code-farmer-II/model"
	"live-code-farmer-II/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BillHandler interface {
}

type billHandler struct {
	billUsecase usecase.BillUsecase
}

func (h *billHandler) CreateBill(c *gin.Context) {
	var request struct {
		Header model.BillHeaderModel `json:"header"`
	}

	if err := c.ShouldBindJSON(&request.Header); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.billUsecase.CreateBill(&request.Header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bill created successfully"})
}

// func (h *billHandler) GetTotalIncomeToday(c *gin.Context) {
// 	totalIncome, err := h.billUsecase.GetTotalIncomeToday()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"total_income": totalIncome})
// }

func NewBillHandler(srv *gin.Engine, billUsecase usecase.BillUsecase) BillHandler {
	billHandler := &billHandler{
		billUsecase: billUsecase,
	}
	// s.GET("/bills", billHandler.GetAllBills)
	// s.GET("/bills/:id", billHandler.GetBillByID)
	// s.GET("/bills/today", billHandler.GetTotalIncomeToday)
	srv.POST("/bills/", billHandler.CreateBill)
	return billHandler
}
