package service 

import "jajar-test/repository" 


type TransactionService struct {
	repo *repository.TransactionRepository
}


func NewTransactionService(repo *repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

// Struct response pagination 
type TransactionPage struct {
	Data       interface{} 
	Page       int         
	Limit      int         
	Total      int64       
	TotalPages int         
}

// List ทำ pagination + เรียก repository
func (s *TransactionService) List(page, limit int) (*TransactionPage, error) {

	if page < 1 { 
		page = 1
	}

	if limit <= 0 { 
		limit = 20 
	}

	if limit > 100 {
		limit = 100
	}

	// ดึงข้อมูลจาก repository
	data, total, err := s.repo.GetPaginated(page, limit)
	if err != nil {
		return nil, err
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	return &TransactionPage{
		Data:       data,
		Page:       page,
		Limit:      limit,
		Total:      total,
		TotalPages: totalPages,
	}, nil
}
