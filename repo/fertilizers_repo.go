package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-II/model"
)

type FertilizersRepo interface {
	GetAllFertilizers() ([]*model.FertilizersModel, error)
	CreateFertilizer(ferti *model.FertilizersModel) error
	GetFertilizerById(int) (*model.FertilizersModel, error)
	GetFertilizerByName(string) (*model.FertilizersModel, error)
	UpdateFertilizer(id int, ferti *model.FertilizersModel) error
	DeleteFertilizer(id int) error
	ReduceStock(fertilizerID int, qty float64) error
}

type fertilizersRepoImpl struct {
	db *sql.DB
}

func (ferRepo *fertilizersRepoImpl) GetAllFertilizers() ([]*model.FertilizersModel, error) {
	qry := "SELECT id, name, stok, price, is_active, created_at,updated_at,created_by,updated_by FROM fertilizers ORDER BY id"

	rows, err := ferRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("GetAllFertilizers() : %w", err)

	}
	defer rows.Close()

	var arrFertilizer []*model.FertilizersModel
	for rows.Next() {
		ferti := &model.FertilizersModel{}
		err := rows.Scan(&ferti.ID, &ferti.Name, &ferti.Stok, &ferti.Price, &ferti.IsActive, &ferti.CreatedAt, &ferti.UpdatedAt, &ferti.CreatedBy, &ferti.UpdatedBy)
		if err != nil {
			return nil, fmt.Errorf("GetAllFertilizers() : %w", err)
		}
		arrFertilizer = append(arrFertilizer, ferti)
	}

	return arrFertilizer, nil
}

func (ferRepo *fertilizersRepoImpl) CreateFertilizer(ferti *model.FertilizersModel) error {
	qry := `INSERT INTO fertilizers (name, stok, price, is_active, created_at, updated_at, created_by, updated_by) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	err := ferRepo.db.QueryRow(
		qry,
		ferti.Name,
		ferti.Stok,
		ferti.Price,
		ferti.IsActive,
		ferti.CreatedAt,
		ferti.UpdatedAt,
		ferti.CreatedBy,
		ferti.UpdatedBy,
	).Scan(&ferti.ID)
	if err != nil {
		return fmt.Errorf("CreateFertilizer() : %w", err)
	}

	return nil

}

func (ferRepo *fertilizersRepoImpl) GetFertilizerById(id int) (*model.FertilizersModel, error) {
	qry := "SELECT id, name, stok, price, is_active, created_at, updated_at, created_by, updated_by FROM fertilizers WHERE id = $1"

	row := ferRepo.db.QueryRow(qry, id)

	ferti := &model.FertilizersModel{}
	err := row.Scan(&ferti.ID, &ferti.Name, &ferti.Stok, &ferti.Price, &ferti.IsActive, &ferti.CreatedAt, &ferti.UpdatedAt, &ferti.CreatedBy, &ferti.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetFertilizerById() : %w", err)
	}

	return ferti, nil
}

func (ferRepo *fertilizersRepoImpl) GetFertilizerByName(name string) (*model.FertilizersModel, error) {
	qry := "SELECT id, name, stok, price, is_active, created_at, updated_at, created_by, updated_by FROM fertilizers WHERE name = $1"

	row := ferRepo.db.QueryRow(qry, name)

	ferti := &model.FertilizersModel{}
	err := row.Scan(&ferti.ID, &ferti.Name, &ferti.Stok, &ferti.Price, &ferti.IsActive, &ferti.CreatedAt, &ferti.UpdatedAt, &ferti.CreatedBy, &ferti.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("GetFertilizerByName() : %w", err)
	}

	return ferti, nil
}

func (ferRepo *fertilizersRepoImpl) UpdateFertilizer(id int, ferti *model.FertilizersModel) error {
	qry := "UPDATE fertilizers SET name = $1, stok = $2, price = $3, is_active = $4, updated_by = $5, updated_at = $6 WHERE id = $7"
	_, err := ferRepo.db.Exec(qry, ferti.Name, ferti.Stok, ferti.Price, ferti.IsActive, ferti.UpdatedBy, ferti.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("error on fertilizersRepoImpl.UpdateFertilizer() : %w", err)
	}
	return nil
}

func (ferRepo *fertilizersRepoImpl) DeleteFertilizer(id int) error {
	qry := "DELETE FROM fertilizers WHERE id = $1;"

	_, err := ferRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("DeleteFertilizer() : %w", err)
	}

	return nil
}

func (ferRepo *fertilizersRepoImpl) ReduceStock(fertilizerID int, qty float64) error {
	query := "UPDATE fertilizers SET stok = stok - $1 WHERE id = $2"
	_, err := ferRepo.db.Exec(query, qty, fertilizerID)
	if err != nil {
		return err
	}

	return nil
}

func NewFertilizersRepo(db *sql.DB) FertilizersRepo {
	return &fertilizersRepoImpl{
		db: db,
	}
}
