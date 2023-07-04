package usecase

import (
	"fmt"
	"live-code-farmer-II/model"
	"live-code-farmer-II/repo"
)

type BillUsecase interface {
	CreateBill(header *model.BillHeaderModel) error
	UpdateTotalBillDetails(header *model.BillHeaderModel)
}

type billUsecase struct {
	billRepo repo.BillRepo
	ferRepo  repo.FertilizersRepo
	farmRepo repo.FarmerRepo
}

func (u *billUsecase) CreateBill(header *model.BillHeaderModel) error {
	cb := header.CreatedBy
	for _, detail := range header.ArrDetails {
		fertilizer, err := u.ferRepo.GetFertilizerById(detail.FertilizerID)
		if err != nil {
			return err
		}
		if fertilizer == nil {
			return fmt.Errorf("fertilizer with ID %d not found", detail.FertilizerID)
		}
		detail.FertilizerName = fertilizer.Name
		detail.Price = fertilizer.Price
		detail.CreatedBy = header.CreatedBy
		detail.UpdatedBy = cb

		err = u.ferRepo.ReduceStock(detail.FertilizerID, detail.Quantity)
		if err != nil {
			return err
		}
	}

	u.UpdateTotalBillDetails(header)
	header.CalculateTotal()

	err := u.billRepo.CreateBillHeader(header)
	if err != nil {
		return err
	}
	return nil
}

func (u *billUsecase) UpdateTotalBillDetails(header *model.BillHeaderModel) {
	for _, detail := range header.ArrDetails {
		detail.Total = detail.Price * detail.Quantity
	}
}

func NewBillUsecase(billRepo repo.BillRepo, ferRepo repo.FertilizersRepo, farmRepo repo.FarmerRepo) BillUsecase {
	return &billUsecase{
		billRepo: billRepo,
		ferRepo:  ferRepo,
		farmRepo: farmRepo,
	}
}
