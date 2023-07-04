package handler

import (
	"errors"
	"fmt"
	"live-code-farmer-II/apperror"
	"live-code-farmer-II/model"
	"live-code-farmer-II/usecase"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type FarmerHandler interface {
}

type farmerHandlerImpl struct {
	farmUsecase usecase.FarmerUsecase
}

func (farmHandler farmerHandlerImpl) GetAllFarmers(ctx *gin.Context) {
	farmer, err := farmHandler.farmUsecase.GetAllFarmers()
	if err != nil {
		fmt.Printf("farmerHandlerImpl.GetAllFarmers() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving farmers data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    farmer,
	})
}

func (farmHandler farmerHandlerImpl) CreateFarmer(ctx *gin.Context) {
	var req model.CreateFarmerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	farmer := model.FarmerModel{
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   req.CreatedBy,
		UpdatedBy:   req.CreatedBy,
	}

	err := farmHandler.farmUsecase.CreateFarmer(&farmer)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("FarmerHandler.CreateFarmer() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("FarmerHandler.CreateFarmer() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Cannot Insert farmer because error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (farmHandler farmerHandlerImpl) GetFarmerById(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id tidak boleh kosong",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id harus angka",
		})
		return
	}
	farmer, err := farmHandler.farmUsecase.GetFarmerById(id)
	if err != nil {
		fmt.Printf("farmerHandlerImpl.GetFarmerById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving farmers data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    farmer,
	})
}

func (farmHandler farmerHandlerImpl) UpdateFarmer(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id cannot be empty",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id must be a number",
		})
		return
	}

	var farmer model.FarmerModel
	err = ctx.ShouldBindJSON(&farmer)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = farmHandler.farmUsecase.UpdateFarmer(id, &farmer)
	if err != nil {
		fmt.Printf("farmerHandlerImpl.GetFarmerById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Cannot Update farmer because error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (farmHandler farmerHandlerImpl) DeleteFarmer(ctx *gin.Context) {
	idText := ctx.Param("id")
	if idText == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id cannot be empty",
		})
		return
	}

	id, err := strconv.Atoi(idText)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success":      false,
			"errorMessage": "Id must be a number",
		})
		return
	}

	err = farmHandler.farmUsecase.DeleteFarmer(id)
	if err != nil {
		fmt.Printf("farmerHandlerImpl.DeleteFarmer()() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Cannot Delete service because error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewFarmerHandler(srv *gin.Engine, farmUsecase usecase.FarmerUsecase) FarmerHandler {
	farmHandler := &farmerHandlerImpl{
		farmUsecase: farmUsecase,
	}
	srv.GET("/farmer/:id", farmHandler.GetFarmerById)
	srv.GET("/farmer", farmHandler.GetAllFarmers)
	srv.POST("/farmer", farmHandler.CreateFarmer)
	srv.PUT("/farmer/:id", farmHandler.UpdateFarmer)
	srv.DELETE("/farmer/:id", farmHandler.DeleteFarmer)

	return farmHandler
}
