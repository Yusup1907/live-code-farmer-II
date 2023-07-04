package manager

import (
	"live-code-farmer-II/repo"
	"sync"
)

type RepoManager interface {
	GetFarmerRepo() repo.FarmerRepo
	GetPlantRepo() repo.PlantRepo
	GetFertilizersRepo() repo.FertilizersRepo
	GetBillRepo() repo.BillRepo
}

type repoManager struct {
	infraManager InfraManager

	farmRepo           repo.FarmerRepo
	plantRepo          repo.PlantRepo
	ferRepo            repo.FertilizersRepo
	billRepo repo.BillRepo
	onceLoadFarmerRepo sync.Once
	onceLoadPlantRepo  sync.Once
	onceLoadPferRepo   sync.Once
	onceLoadBillRepo   sync.Once
}

func (rm *repoManager) GetFarmerRepo() repo.FarmerRepo {
	rm.onceLoadFarmerRepo.Do(func() {
		rm.farmRepo = repo.NewFarmerRepo(rm.infraManager.GetDB())
	})
	return rm.farmRepo
}

func (rm *repoManager) GetPlantRepo() repo.PlantRepo {
	rm.onceLoadPlantRepo.Do(func() {
		rm.plantRepo = repo.NewPlantRepo(rm.infraManager.GetDB())
	})
	return rm.plantRepo
}

func (rm *repoManager) GetFertilizersRepo() repo.FertilizersRepo {
	rm.onceLoadPferRepo.Do(func() {
		rm.ferRepo = repo.NewFertilizersRepo(rm.infraManager.GetDB())
	})
	return rm.ferRepo
}

func (rm *repoManager) GetBillRepo() repo.BillRepo {
	rm.onceLoadBillRepo.Do(func() {
		rm.billRepo = repo.NewBillRepo(rm.infraManager.GetDB())
	})
	return rm.billRepo
}

func NewRepoManager(infraManager InfraManager) RepoManager {
	return &repoManager{
		infraManager: infraManager,
	}
}
