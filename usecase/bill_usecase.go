package usecase

import (
	"fmt"
	"live-code-farmer-II/model"
	"live-code-farmer-II/repo"
	"time"
)

type BillUsecase interface {
	CreateBill(header *model.BillHeaderModel) error
	UpdateTotalBillDetails(header *model.BillHeaderModel)
	GetTotalIncomeToday() (float64, error)
	GetTotalIncomeMonthly(year int, month time.Month) (float64, error)
	GetTotalIncomeYearly(year int) (float64, error)
	GetBillByID(id int) (*model.BillHeaderModel, error)
	GetAllBills() ([]*model.BillHeaderModel, error)
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

func (u *billUsecase) GetTotalIncomeToday() (float64, error) {
	totalIncome, err := u.billRepo.GetTotalIncomeToday()
	if err != nil {
		return 0, err
	}

	return totalIncome, nil
}

func (u *billUsecase) GetTotalIncomeMonthly(year int, month time.Month) (float64, error) {
	totalIncome, err := u.billRepo.GetTotalIncomeMonthly(year, month)
	if err != nil {
		return 0, err
	}

	return totalIncome, nil
}

func (u *billUsecase) GetTotalIncomeYearly(year int) (float64, error) {
	totalIncome, err := u.billRepo.GetTotalIncomeYearly(year)
	if err != nil {
		return 0, err
	}

	return totalIncome, nil
}

func (u *billUsecase) GetBillByID(id int) (*model.BillHeaderModel, error) {
	bill, err := u.billRepo.GetBillByID(id)
	if err != nil {
		return nil, err
	}

	return bill, nil
}

func (u *billUsecase) GetAllBills() ([]*model.BillHeaderModel, error) {
	bills, err := u.billRepo.GetAllBills()
	if err != nil {
		return nil, err
	}

	details, err := u.billRepo.GetAllBillDetails()
	if err != nil {
		return nil, err
	}

	for _, bill := range bills {
		for _, detail := range details {
			if detail.BillID == bill.ID {
				bill.ArrDetails = append(bill.ArrDetails, detail)
			}
		}
	}

	return bills, nil
}

func NewBillUsecase(billRepo repo.BillRepo, ferRepo repo.FertilizersRepo, farmRepo repo.FarmerRepo) BillUsecase {
	return &billUsecase{
		billRepo: billRepo,
		ferRepo:  ferRepo,
		farmRepo: farmRepo,
	}
}
