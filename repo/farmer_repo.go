package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-II/model"
)

type FarmerRepo interface {
	GetAllFarmers() ([]*model.FarmerModel, error)
	CreateFarmer(farmer *model.FarmerModel) error
	GetFarmerById(int) (*model.FarmerModel, error)
	GetfarmerByName(name string) (*model.FarmerModel, error)
	UpdateFarmer(id int, farmer *model.FarmerModel) error
	DeleteFarmer(int) error
}

type farmerRepoImpl struct {
	db *sql.DB
}

func (farmRepo *farmerRepoImpl) GetAllFarmers() ([]*model.FarmerModel, error) {
	qry := "SELECT id, name, address, phone_number,created_at,updated_at,created_by,updated_by FROM farmers ORDER BY id"

	rows, err := farmRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("GetFarGetAllFarmersmers() : %w", err)

	}
	defer rows.Close()

	var arrFarmer []*model.FarmerModel
	for rows.Next() {
		farmers := &model.FarmerModel{}
		err := rows.Scan(&farmers.ID, &farmers.Name, &farmers.Address, &farmers.PhoneNumber, &farmers.CreatedAt, &farmers.UpdatedAt, &farmers.CreatedBy, &farmers.UpdatedBy)
		if err != nil {
			return nil, fmt.Errorf("GetAllFarmers() : %w", err)
		}
		arrFarmer = append(arrFarmer, farmers)
	}

	return arrFarmer, nil
}

func (farmRepo *farmerRepoImpl) CreateFarmer(farmer *model.FarmerModel) error {
	qry := `INSERT INTO farmers (name, address, phone_number, created_at, updated_at, created_by, updated_by) 
			VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	err := farmRepo.db.QueryRow(
		qry,
		farmer.Name,
		farmer.Address,
		farmer.PhoneNumber,
		farmer.CreatedAt,
		farmer.UpdatedAt,
		farmer.CreatedBy,
		farmer.UpdatedBy,
	).Scan(&farmer.ID)
	if err != nil {
		return fmt.Errorf("CreateFarmer() : %w", err)
	}

	return nil

}

func (farmRepo *farmerRepoImpl) GetFarmerById(id int) (*model.FarmerModel, error) {
	qry := "SELECT id, name, address, phone_number, created_at, updated_at, created_by, updated_by FROM farmers WHERE id = $1"

	row := farmRepo.db.QueryRow(qry, id)

	farmer := &model.FarmerModel{}
	err := row.Scan(&farmer.ID, &farmer.Name, &farmer.Address, &farmer.PhoneNumber, &farmer.CreatedAt, &farmer.UpdatedAt, &farmer.CreatedBy, &farmer.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Petani dengan ID tersebut tidak ditemukan
		}
		return nil, fmt.Errorf("GetFarmerByID() : %w", err)
	}

	return farmer, nil
}

func (farmRepo *farmerRepoImpl) GetfarmerByName(name string) (*model.FarmerModel, error) {
	qry := "SELECT id, name, address, phone_number, created_at, updated_at, created_by, updated_by FROM farmers WHERE name = $1"

	row := farmRepo.db.QueryRow(qry, name)

	farmer := &model.FarmerModel{}
	err := row.Scan(&farmer.ID, &farmer.Name, &farmer.Address, &farmer.PhoneNumber, &farmer.CreatedAt, &farmer.UpdatedAt, &farmer.CreatedBy, &farmer.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Petani dengan ID tersebut tidak ditemukan
		}
		return nil, fmt.Errorf("GetfarmerByName() : %w", err)
	}

	return farmer, nil
}

func (farmRepo *farmerRepoImpl) UpdateFarmer(id int, farmer *model.FarmerModel) error {
	qry := "UPDATE farmers SET name = $1, address = $2, phone_number = $3, updated_by = $4, updated_at = $5 WHERE id = $6"
	_, err := farmRepo.db.Exec(qry, farmer.Name, farmer.Address, farmer.PhoneNumber, farmer.UpdatedBy, farmer.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("error on farmerRepoImpl.UpdateFarmer() : %w", err)
	}
	return nil
}

func (farmRepo *farmerRepoImpl) DeleteFarmer(id int) error {
	qry := "DELETE FROM farmers WHERE id = $1;"

	_, err := farmRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("deleteService() : %w", err)
	}

	return nil
}

func NewFarmerRepo(db *sql.DB) FarmerRepo {
	return &farmerRepoImpl{
		db: db,
	}
}
