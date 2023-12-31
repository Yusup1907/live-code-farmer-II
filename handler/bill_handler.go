package handler

import (
	"live-code-farmer-II/model"
	"live-code-farmer-II/usecase"
	"net/http"
	"strconv"
	"time"

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

func (h *billHandler) GetTotalIncomeToday(c *gin.Context) {
	totalIncome, err := h.billUsecase.GetTotalIncomeToday()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_income": totalIncome})
}

func (h *billHandler) GetTotalIncomeMonthly(c *gin.Context) {
	year := c.Query("year")
	month := c.Query("month")

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}

	monthInt, err := strconv.Atoi(month)
	if err != nil || monthInt < 1 || monthInt > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month"})
		return
	}

	totalIncome, err := h.billUsecase.GetTotalIncomeMonthly(yearInt, time.Month(monthInt))
	if err != nil {
		// Handle error dari use case
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_income": totalIncome})
}

func (h *billHandler) GetTotalIncomeYearly(c *gin.Context) {
	year := c.Query("year")

	yearInt, err := strconv.Atoi(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year"})
		return
	}

	totalIncome, err := h.billUsecase.GetTotalIncomeYearly(yearInt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get total income"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"total_income": totalIncome})
}

func (h *billHandler) GetBillByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid bill ID"})
		return
	}

	bill, err := h.billUsecase.GetBillByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if bill == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bill not found"})
		return
	}

	c.JSON(http.StatusOK, bill)
}

func (h *billHandler) GetAllBills(c *gin.Context) {
	bills, err := h.billUsecase.GetAllBills()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, bills)
}

func NewBillHandler(srv *gin.Engine, billUsecase usecase.BillUsecase) BillHandler {
	billHandler := &billHandler{
		billUsecase: billUsecase,
	}
	srv.GET("/bills", billHandler.GetAllBills)
	srv.GET("/bills/:id", billHandler.GetBillByID)
	srv.GET("/bills/today", billHandler.GetTotalIncomeToday)
	srv.GET("/bills/monthly", billHandler.GetTotalIncomeMonthly)
	srv.GET("/bills/yearly", billHandler.GetTotalIncomeYearly)
	srv.POST("/bills/", billHandler.CreateBill)
	return billHandler
}
