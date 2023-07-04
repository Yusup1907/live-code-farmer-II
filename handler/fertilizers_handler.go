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

type FertilizersHandler interface {
}

type fertilizersHandlerImpl struct {
	ferUsecase usecase.FertilizersUsecase
}

func (ferHandler fertilizersHandlerImpl) GetAllFertilizers(ctx *gin.Context) {
	ferti, err := ferHandler.ferUsecase.GetAllFertilizers()
	if err != nil {
		fmt.Printf("fertilizersHandlerImpl.GetAllFertilizers() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving ferti data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ferti,
	})
}

func (ferHandler fertilizersHandlerImpl) CreateFertilizer(ctx *gin.Context) {
	var req model.CreatefertilizersRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ferti := model.FertilizersModel{
		Name:      req.Name,
		Stok:      req.Stok,
		Price:     req.Price,
		IsActive:  req.IsActive,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	}

	err := ferHandler.ferUsecase.CreateFertilizer(&ferti)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("ferHandler.CreateFertilizer() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("ferHandler.CreateFertilizer() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Cannot Insert ferti because error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ferHandler fertilizersHandlerImpl) GetFertilizerById(ctx *gin.Context) {
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
	ferti, err := ferHandler.ferUsecase.GetFertilizerById(id)
	if err != nil {
		fmt.Printf("fertilizersHandlerImpl.GetFertilizerById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving ferti data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    ferti,
	})
}

func (ferHandler fertilizersHandlerImpl) UpdateFertilizer(ctx *gin.Context) {
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

	var ferti model.FertilizersModel
	err = ctx.ShouldBindJSON(&ferti)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = ferHandler.ferUsecase.UpdateFertilizer(id, &ferti)
	if err != nil {
		fmt.Printf("fertilizersHandlerImpl.UpdateFertilizer() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Cannot Update ferti because error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (ferHandler fertilizersHandlerImpl) DeleteFertilizer(ctx *gin.Context) {
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

	err = ferHandler.ferUsecase.DeleteFertilizer(id)
	if err != nil {
		fmt.Printf("farHandlerImpl.DeleteFertilizer()() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Cannot Delete Fertilizer because error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewFertilizersHandler(srv *gin.Engine, ferUsecase usecase.FertilizersUsecase) FertilizersHandler {
	fertHandler := &fertilizersHandlerImpl{
		ferUsecase: ferUsecase,
	}
	srv.GET("/fertilizer/:id", fertHandler.GetFertilizerById)
	srv.GET("/fertilizer", fertHandler.GetAllFertilizers)
	srv.POST("/fertilizer", fertHandler.CreateFertilizer)
	srv.PUT("/fertilizer/:id", fertHandler.UpdateFertilizer)
	srv.DELETE("/fertilizer/:id", fertHandler.DeleteFertilizer)

	return fertHandler
}
