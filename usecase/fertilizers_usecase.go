package usecase

import (
	"fmt"
	"live-code-farmer-II/apperror"
	"live-code-farmer-II/model"
	"live-code-farmer-II/repo"
)

type FertilizersUsecase interface {
	GetAllFertilizers() ([]*model.FertilizersModel, error)
	CreateFertilizer(ferti *model.FertilizersModel) error
	GetFertilizerById(id int) (*model.FertilizersModel, error)
	UpdateFertilizer(id int, ferti *model.FertilizersModel) error
	DeleteFertilizer(id int) error
}

type fertilizerUsecaseImpl struct {
	ferRepo repo.FertilizersRepo
}

func (ferUsecase *fertilizerUsecaseImpl) GetAllFertilizers() ([]*model.FertilizersModel, error) {
	return ferUsecase.ferRepo.GetAllFertilizers()
}

func (ferUsecase *fertilizerUsecaseImpl) CreateFertilizer(ferti *model.FertilizersModel) error {
	svcDB, err := ferUsecase.ferRepo.GetFertilizerByName(ferti.Name)
	if err != nil {
		return fmt.Errorf("fertilizerUsecaseImpl.CreateFertilizer() : %w", err)
	}

	if svcDB != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data Fertilizers dengan nama %v sudah ada", ferti.Name),
		}
	}

	return ferUsecase.ferRepo.CreateFertilizer(ferti)
}

func (ferUsecase *fertilizerUsecaseImpl) GetFertilizerById(id int) (*model.FertilizersModel, error) {
	return ferUsecase.ferRepo.GetFertilizerById(id)
}

func (ferUsecase *fertilizerUsecaseImpl) UpdateFertilizer(id int, ferti *model.FertilizersModel) error {
	farmDB, err := ferUsecase.ferRepo.GetFertilizerById(id)
	if err != nil {
		return fmt.Errorf("fertilizerUsecaseImpl.UpdateFertilizer() : %w", err)
	}

	if farmDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data ferti dengan id %v tidak ada", ferti.ID),
		}
	}

	farmDb, err := ferUsecase.ferRepo.GetFertilizerByName(ferti.Name)
	if err != nil {
		return fmt.Errorf("fertilizerUsecaseImpl.UpdateFertilizer() : %w", err)
	}

	if farmDb != nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data ferti dengan nama %v sudah ada", ferti.Name),
		}
	}

	return ferUsecase.ferRepo.UpdateFertilizer(id, ferti)
}

func (ferUsecase *fertilizerUsecaseImpl) DeleteFertilizer(id int) error {
	farmDB, err := ferUsecase.ferRepo.GetFertilizerById(id)
	if err != nil {
		return fmt.Errorf("fertilizerUsecaseImpl.DeleteFertilizer() : %w", err)
	}

	if farmDB == nil {
		return apperror.AppError{
			ErrorCode:    1,
			ErrorMassage: fmt.Sprintf("data ferti dengan id %v tidak ada", id),
		}
	}

	return ferUsecase.ferRepo.DeleteFertilizer(id)
}

func NewFertilizersUseCase(ferRepo repo.FertilizersRepo) FertilizersUsecase {
	return &fertilizerUsecaseImpl{
		ferRepo: ferRepo,
	}
}
