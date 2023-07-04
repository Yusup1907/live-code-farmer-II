package repo

import (
	"database/sql"
	"live-code-farmer-II/model"
	"time"
)

type BillRepo interface {
	CreateBillHeader(header *model.BillHeaderModel) error
	GetBillByID(id int) (*model.BillHeaderModel, error)
	GetAllBills() ([]*model.BillHeaderModel, error)
	GetAllBillDetails() ([]*model.BillDetailModel, error)
	GetTotalIncomeToday() (float64, error)
}

type billRepo struct {
	db *sql.DB
}

func (r *billRepo) CreateBillHeader(header *model.BillHeaderModel) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	currentDate := time.Now().Format("2006-01-02")
	query := `INSERT INTO bills (farmer_id, name, address, phone_number, date, created_at, updated_at, created_by, updated_by) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`
	err = tx.QueryRow(query, header.FarmerID, header.Name, header.Address, header.PhoneNumber, currentDate, time.Now(), time.Now(), header.CreatedBy, header.CreatedBy).Scan(&header.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	query = "INSERT INTO bill_details (bill_id, fertilizer_id, fertilizer_name, quantity, price, created_at, updated_at, created_by, updated_by,total) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)"
	for _, detail := range header.ArrDetails {
		_, err = tx.Exec(query, header.ID, detail.FertilizerID, detail.FertilizerName, detail.Quantity, detail.Price, time.Now(), time.Now(), detail.CreatedBy, detail.CreatedBy, detail.Total)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	updateQuery := "UPDATE bills SET total = $1 WHERE id = $2"
	_, err = tx.Exec(updateQuery, header.Total, header.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (r *billRepo) GetBillByID(id int) (*model.BillHeaderModel, error) {
	headerQuery := "SELECT id, farmer_id, name, address, phone_number, date, created_at, updated_at, created_by, updated_by, total FROM bills WHERE id = $1"
	headerRow := r.db.QueryRow(headerQuery, id)

	header := &model.BillHeaderModel{}
	err := headerRow.Scan(&header.ID, &header.FarmerID, &header.Name, &header.Address, &header.PhoneNumber, &header.Date, &header.CreatedAt, &header.UpdatedAt, &header.CreatedBy, &header.UpdatedBy, &header.Total)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	detailQuery := "SELECT id, bill_id, fertilizer_id, fertilizer_name, quantity, price,created_at,updated_at,created_by,updated_by, total FROM bill_details WHERE bill_id = $1"
	rows, err := r.db.Query(detailQuery, header.ID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	details := []*model.BillDetailModel{}
	for rows.Next() {
		detail := &model.BillDetailModel{}
		err := rows.Scan(&detail.ID, &detail.BillID, &detail.FertilizerID, &detail.FertilizerName, &detail.Quantity, &detail.Price, &detail.CreatedAt, &detail.UpdatedAt, &detail.CreatedBy, &detail.UpdatedBy, &detail.Total)
		if err != nil {
			return nil, err
		}

		details = append(details, detail)
	}

	header.ArrDetails = details

	return header, nil
}

func (r *billRepo) GetAllBills() ([]*model.BillHeaderModel, error) {
	query := "SELECT id, farmer_id, name, address, phone_number, date, created_at, updated_at, created_by, updated_by, total FROM bills ORDER BY id DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	bills := []*model.BillHeaderModel{}
	for rows.Next() {
		bill := &model.BillHeaderModel{}
		err := rows.Scan(&bill.ID, &bill.FarmerID, &bill.Name, &bill.Address, &bill.PhoneNumber, &bill.Date, &bill.CreatedAt, &bill.UpdatedAt, &bill.CreatedBy, &bill.UpdatedBy, &bill.Total)
		if err != nil {
			return nil, err
		}

		bills = append(bills, bill)
	}

	return bills, nil
}

func (r *billRepo) GetAllBillDetails() ([]*model.BillDetailModel, error) {
	query := "SELECT id, bill_id, fertilizer_id, fertilizer_name, quantity, price,created_at,updated_at,created_by,updated_by, total FROM bill_details order by id DESC"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	details := []*model.BillDetailModel{}
	for rows.Next() {
		detail := &model.BillDetailModel{}
		err := rows.Scan(&detail.ID, &detail.BillID, &detail.FertilizerID, &detail.FertilizerName, &detail.Quantity, &detail.Price, &detail.CreatedAt, &detail.UpdatedAt, &detail.CreatedBy, &detail.UpdatedBy, &detail.Total)
		if err != nil {
			return nil, err
		}

		details = append(details, detail)
	}

	return details, nil
}

func (r *billRepo) GetTotalIncomeToday() (float64, error) {
	query := "SELECT COALESCE(SUM(total), 0) FROM bills WHERE DATE(date) = DATE(NOW())"
	var totalIncome float64
	err := r.db.QueryRow(query).Scan(&totalIncome)
	if err != nil {
		return 0, err
	}

	return totalIncome, nil
}

func NewBillRepo(db *sql.DB) BillRepo {
	return &billRepo{
		db: db,
	}
}
