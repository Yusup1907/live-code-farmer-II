package repo

import (
	"database/sql"
	"fmt"
	"live-code-farmer-II/model"
)

type PlantRepo interface {
	GetAllPlants() ([]*model.PlantModel, error)
	CreatePlant(plant *model.PlantModel) error
	GetPlantById(int) (*model.PlantModel, error)
	GetPlantByName(string) (*model.PlantModel, error)
	UpdatePlant(id int, plant *model.PlantModel) error
	DeletePlant(int) error
}

type plantRepoImpl struct {
	db *sql.DB
}

func (plantRepo *plantRepoImpl) GetAllPlants() ([]*model.PlantModel, error) {
	qry := "SELECT id, name, created_at,updated_at,created_by,updated_by FROM plants ORDER BY id"

	rows, err := plantRepo.db.Query(qry)
	if err != nil {
		return nil, fmt.Errorf("GetAllPlants() : %w", err)

	}
	defer rows.Close()

	var arrPlants []*model.PlantModel
	for rows.Next() {
		plants := &model.PlantModel{}
		err := rows.Scan(&plants.ID, &plants.Name, &plants.CreatedAt, &plants.UpdatedAt, &plants.CreatedBy, &plants.UpdatedBy)
		if err != nil {
			return nil, fmt.Errorf("GetAllPlants() : %w", err)
		}
		arrPlants = append(arrPlants, plants)
	}

	return arrPlants, nil
}

func (plantRepo *plantRepoImpl) CreatePlant(plant *model.PlantModel) error {
	qry := `INSERT INTO plants (name, created_at, updated_at, created_by, updated_by) 
			VALUES ($1, $2, $3, $4, $5) RETURNING id`
	err := plantRepo.db.QueryRow(
		qry,
		plant.Name,
		plant.CreatedAt,
		plant.UpdatedAt,
		plant.CreatedBy,
		plant.UpdatedBy,
	).Scan(&plant.ID)
	if err != nil {
		return fmt.Errorf("CreatePlant() : %w", err)
	}

	return nil

}

func (plantRepo *plantRepoImpl) GetPlantById(id int) (*model.PlantModel, error) {
	qry := "SELECT id, name, created_at, updated_at, created_by, updated_by FROM plants WHERE id = $1"

	row := plantRepo.db.QueryRow(qry, id)

	plant := &model.PlantModel{}
	err := row.Scan(&plant.ID, &plant.Name, &plant.CreatedAt, &plant.UpdatedAt, &plant.CreatedBy, &plant.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Petani dengan ID tersebut tidak ditemukan
		}
		return nil, fmt.Errorf("GetFarmerByID() : %w", err)
	}

	return plant, nil
}

func (plantRepo *plantRepoImpl) GetPlantByName(name string) (*model.PlantModel, error) {
	qry := "SELECT id, name, created_at, updated_at, created_by, updated_by FROM plants WHERE name = $1"

	row := plantRepo.db.QueryRow(qry, name)

	plant := &model.PlantModel{}
	err := row.Scan(&plant.ID, &plant.Name, &plant.CreatedAt, &plant.UpdatedAt, &plant.CreatedBy, &plant.UpdatedBy)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Tanaman dengan ID tersebut tidak ditemukan
		}
		return nil, fmt.Errorf("GetPlantByName() : %w", err)
	}

	return plant, nil
}

func (plantRepo *plantRepoImpl) UpdatePlant(id int, plant *model.PlantModel) error {
	qry := "UPDATE plants SET name = $1, updated_by = $2, updated_at = $3 WHERE id = $4"
	_, err := plantRepo.db.Exec(qry, plant.Name, plant.UpdatedBy, plant.UpdatedAt, id)
	if err != nil {
		return fmt.Errorf("error on plantRepoImpl.UpdatePlant() : %w", err)
	}
	return nil
}

func (plantRepo *plantRepoImpl) DeletePlant(id int) error {
	qry := "DELETE FROM plants WHERE id = $1;"

	_, err := plantRepo.db.Exec(qry, id)
	if err != nil {
		return fmt.Errorf("DeletePlant() : %w", err)
	}

	return nil
}

func NewPlantRepo(db *sql.DB) PlantRepo {
	return &plantRepoImpl{
		db: db,
	}
}
