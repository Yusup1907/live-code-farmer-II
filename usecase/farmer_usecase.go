package usecase

import (
	"fmt"
	"live-code-farmer-II/apperror"
	"live-code-farmer-II/model"
	"live-code-farmer-II/repo"
)

type FarmerUsecase interface {
	GetAllFarmers() ([]*model.FarmerModel, error)
	CreateFarmer(farmer *model.FarmerModel) error
	GetFarmerById(int) (*model.FarmerModel, error)
	UpdateFarmer(id int, farmer *model.FarmerModel) error
	DeleteFarmer(id int) error
}

type farmerUsecaseImpl struct {
	farmRepo repo.FarmerRepo
}

func (farmUsecase *farmerUsecaseImpl) GetAllFarmers() ([]*model.FarmerModel, error) {
	return farmUsecase.farmRepo.GetAllFarmers()
}

func (farmUsecase *farmerUsecaseImpl) CreateFarmer(farmer *model.FarmerModel) error {
	svcDB, err := farmUsecase.farmRepo.GetfarmerByName(farmer.Name)
	if err != nil {
		return fmt.Errorf("farmerUsecaseImpl.CreateFarmer() : %w", err)
	}

	if svcDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data farmer dengan nama %v sudah ada", farmer.Name),
		}
	}

	return farmUsecase.farmRepo.CreateFarmer(farmer)
}

func (farmUsecase *farmerUsecaseImpl) GetFarmerById(id int) (*model.FarmerModel, error) {
	return farmUsecase.farmRepo.GetFarmerById(id)
}

func (farmUsecase *farmerUsecaseImpl) UpdateFarmer(id int, farmer *model.FarmerModel) error {
	farmDB, err := farmUsecase.farmRepo.GetFarmerById(id)
	if err != nil {
		return fmt.Errorf("farmerUsecaseImpl.UpdateService() : %w", err)
	}

	if farmDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data farmer dengan id %v tidak ada", farmer.ID),
		}
	}

	farmDb, err := farmUsecase.farmRepo.GetfarmerByName(farmer.Name)
	if err != nil {
		return fmt.Errorf("farmerUsecaseImpl.UpdateService() : %w", err)
	}

	if farmDb != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data farmer dengan nama %v sudah ada", farmer.Name),
		}
	}

	return farmUsecase.farmRepo.UpdateFarmer(id, farmer)
}

func (farmUsecase *farmerUsecaseImpl) DeleteFarmer(id int) error {
	farmDB, err := farmUsecase.farmRepo.GetFarmerById(id)
	if err != nil {
		return fmt.Errorf("farmerUsecaseImpl.UpdateService() : %w", err)
	}

	if farmDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data farmer dengan id %v tidak ada", id),
		}
	}

	return farmUsecase.farmRepo.DeleteFarmer(id)
}

func NewFarmerUseCase(farmRepo repo.FarmerRepo) FarmerUsecase {
	return &farmerUsecaseImpl{
		farmRepo: farmRepo,
	}
}
