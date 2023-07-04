package manager

import (
	"live-code-farmer-II/usecase"
	"sync"
)

type UsecaseManager interface {
	GetFarmerUsecase() usecase.FarmerUsecase
	GetPlantUsecase() usecase.PlantUsecase
	GetFertilizersUsecase() usecase.FertilizersUsecase
	GetBillUsecase() usecase.BillUsecase
}

type usecaseManager struct {
	repoManager RepoManager

	farmUsecase  usecase.FarmerUsecase
	plantUsecase usecase.PlantUsecase
	ferUsecase   usecase.FertilizersUsecase
	billUsecase  usecase.BillUsecase

	onceLoadFarmerUsecase sync.Once
	onceLoadPlantUsecase  sync.Once
	onceLoadFerUsecase    sync.Once
	onceLoadBillUsecase   sync.Once
}

func (um *usecaseManager) GetFarmerUsecase() usecase.FarmerUsecase {
	um.onceLoadFarmerUsecase.Do(func() {
		um.farmUsecase = usecase.NewFarmerUseCase(um.repoManager.GetFarmerRepo())
	})
	return um.farmUsecase
}

func (um *usecaseManager) GetPlantUsecase() usecase.PlantUsecase {
	um.onceLoadPlantUsecase.Do(func() {
		um.plantUsecase = usecase.NewPlantUseCase(um.repoManager.GetPlantRepo())
	})
	return um.plantUsecase
}

func (um *usecaseManager) GetFertilizersUsecase() usecase.FertilizersUsecase {
	um.onceLoadFerUsecase.Do(func() {
		um.ferUsecase = usecase.NewFertilizersUseCase(um.repoManager.GetFertilizersRepo())
	})
	return um.ferUsecase
}

func (um *usecaseManager) GetBillUsecase() usecase.BillUsecase {
	um.onceLoadBillUsecase.Do(func() {
		um.billUsecase = usecase.NewBillUsecase(um.repoManager.GetBillRepo(), um.repoManager.GetFertilizersRepo(), um.repoManager.GetFarmerRepo())
	})
	return um.billUsecase
}

// func (um *usecaseManager) GetTransactionUsecase() usecase.TransactionUsecase {
// 	um.onceLoadTransactionUsecase.Do(func() {
// 		trxRepo := um.repoManager.GetTransactionRepo()
// 		svcRepo := um.repoManager.GetServiceRepo()
// 		um.trxUsecase = usecase.NewTransactionUseCase(trxRepo, svcRepo)
// 	})
// 	return um.trxUsecase
// }

func NewUsecaseManager(repoManager RepoManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
