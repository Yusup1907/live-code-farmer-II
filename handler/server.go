package handler

import (
	"live-code-farmer-II/manager"

	"github.com/gin-gonic/gin"
)

type Server interface {
	Run()
}

type server struct {
	usecaseManager manager.UsecaseManager
	srv            *gin.Engine
}

func (s *server) Run() {
	NewFarmerHandler(s.srv, s.usecaseManager.GetFarmerUsecase())
	NewPlantHandler(s.srv, s.usecaseManager.GetPlantUsecase())
	NewFertilizersHandler(s.srv, s.usecaseManager.GetFertilizersUsecase())
	NewFertilizerPriceHandler(s.srv, s.usecaseManager.GetFertilizerPriceUsecase())

	s.srv.Run()
}

func NewServer() Server {
	infra := manager.NewInfraManager()
	repo := manager.NewRepoManager(infra)
	usecase := manager.NewUsecaseManager(repo)

	srv := gin.Default()

	return &server{
		usecaseManager: usecase,
		srv:            srv,
	}

}
