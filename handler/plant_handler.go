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

type PlantHandler interface {
}

type plantHandlerImpl struct {
	plantUsecase usecase.PlantUsecase
}

func (plantHandler plantHandlerImpl) GetAllPlants(ctx *gin.Context) {
	plant, err := plantHandler.plantUsecase.GetAllPlants()
	if err != nil {
		fmt.Printf("plantHandlerImpl.GetAllPlants() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving plants data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    plant,
	})
}

func (plantHandler plantHandlerImpl) CreatePlant(ctx *gin.Context) {
	var req model.CreatePlantRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	plant := model.PlantModel{
		Name:      req.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: req.CreatedBy,
		UpdatedBy: req.CreatedBy,
	}

	err := plantHandler.plantUsecase.CreatePlant(&plant)
	if err != nil {
		appError := apperror.AppError{}
		if errors.As(err, &appError) {
			fmt.Printf("plantHandler.CreatePlant() 1 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": appError.Error(),
			})
		} else {
			fmt.Printf("plantHandler.CreatePlant() 2 : %v ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success":      false,
				"errorMessage": "Cannot Insert plant because error",
			})
		}
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (plantHandler plantHandlerImpl) GetPlantById(ctx *gin.Context) {
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
	plant, err := plantHandler.plantUsecase.GetPlantById(id)
	if err != nil {
		fmt.Printf("plantHandlerImpl.GetPlantById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "An error occurred when retrieving plants data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    plant,
	})
}

func (plantHandler plantHandlerImpl) UpdatePlant(ctx *gin.Context) {
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

	var plant model.PlantModel
	err = ctx.ShouldBindJSON(&plant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = plantHandler.plantUsecase.UpdatePlant(id, &plant)
	if err != nil {
		fmt.Printf("plantHandlerImpl.GetPlantById() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Cannot Update plants because error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func (plantHandler plantHandlerImpl) DeletePlant(ctx *gin.Context) {
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

	err = plantHandler.plantUsecase.DeletePlant(id)
	if err != nil {
		fmt.Printf("plantHandlerImpl.DeletePlant()() : %v ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success":      false,
			"errorMessage": "Cannot Delete plant because error",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
	})
}

func NewPlantHandler(srv *gin.Engine, plantUsecase usecase.PlantUsecase) PlantHandler {
	plantHandler := &plantHandlerImpl{
		plantUsecase: plantUsecase,
	}
	srv.GET("/plant/:id", plantHandler.GetPlantById)
	srv.GET("/plant", plantHandler.GetAllPlants)
	srv.POST("/plant", plantHandler.CreatePlant)
	srv.PUT("/plant/:id", plantHandler.UpdatePlant)
	srv.DELETE("/plant/:id", plantHandler.DeletePlant)

	return plantHandler
}
