package repository

import (
	"gorm.io/gorm"
	"jajar-test/model"
)

type TransactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetPaginated(page, limit int) ([]model.Transaction, int64, error) {

	offset := (page - 1) * limit // คำนวณ offset

	var total int64
	// count table
	if err := r.db.Model(&model.Transaction{}).Count(&total).Error; err != nil {
		return nil, 0, err 
	}

	var txs []model.Transaction
	// query
	err := r.db.
		Order("created_time DESC").
		Limit(limit).
		Offset(offset).
		Find(&txs).
		Error

	if err != nil {
		return nil, 0, err
	}

	return txs, total, nil
}
