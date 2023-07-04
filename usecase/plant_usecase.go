package usecase

import (
	"fmt"
	"live-code-farmer-II/apperror"
	"live-code-farmer-II/model"
	"live-code-farmer-II/repo"
)

type PlantUsecase interface {
	GetAllPlants() ([]*model.PlantModel, error)
	CreatePlant(plant *model.PlantModel) error
	GetPlantById(int) (*model.PlantModel, error)
	UpdatePlant(id int, plant *model.PlantModel) error
	DeletePlant(id int) error
}

type plantUsecaseImpl struct {
	plantRepo repo.PlantRepo
}

func (plantUsecase *plantUsecaseImpl) GetAllPlants() ([]*model.PlantModel, error) {
	return plantUsecase.plantRepo.GetAllPlants()
}

func (plantUsecase *plantUsecaseImpl) CreatePlant(plant *model.PlantModel) error {
	svcDB, err := plantUsecase.plantRepo.GetPlantByName(plant.Name)
	if err != nil {
		return fmt.Errorf("plantUsecaseImpl.CreatePlant() : %w", err)
	}

	if svcDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data tanaman dengan nama %v sudah ada", plant.Name),
		}
	}

	return plantUsecase.plantRepo.CreatePlant(plant)
}

func (plantUsecase *plantUsecaseImpl) GetPlantById(id int) (*model.PlantModel, error) {
	return plantUsecase.plantRepo.GetPlantById(id)
}

func (plantUsecase *plantUsecaseImpl) UpdatePlant(id int, plant *model.PlantModel) error {
	plantDB, err := plantUsecase.plantRepo.GetPlantById(id)
	if err != nil {
		return fmt.Errorf("plantUsecaseImpl.UpdatePlant() : %w", err)
	}

	if plantDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data tanaman dengan id %v tidak ada", plant.ID),
		}
	}

	plantDb, err := plantUsecase.plantRepo.GetPlantByName(plant.Name)
	if err != nil {
		return fmt.Errorf("plantUsecaseImpl.UpdatePlant() : %w", err)
	}

	if plantDb != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data tanaman dengan nama %v sudah ada", plant.Name),
		}
	}

	return plantUsecase.plantRepo.UpdatePlant(id, plant)
}

func (plantUsecase *plantUsecaseImpl) DeletePlant(id int) error {
	plantDB, err := plantUsecase.plantRepo.GetPlantById(id)
	if err != nil {
		return fmt.Errorf("farmerUsecaseImpl.UpdateService() : %w", err)
	}

	if plantDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data farmer dengan id %v tidak ada", id),
		}
	}

	return plantUsecase.plantRepo.DeletePlant(id)
}

func NewPlantUseCase(plantRepo repo.PlantRepo) PlantUsecase {
	return &plantUsecaseImpl{
		plantRepo: plantRepo,
	}
}
