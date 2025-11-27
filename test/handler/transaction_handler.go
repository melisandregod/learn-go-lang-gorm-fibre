package handler 

import (
	"strconv" 
	"github.com/gofiber/fiber/v2" 
	"jajar-test/service"               
)

type TransactionHandler struct {
	svc *service.TransactionService 
}

func NewTransactionHandler(svc *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{svc: svc}
}

// helper แปลง string -> int 
func parseInt(s string, def int) int {
	v, err := strconv.Atoi(s)
	if err != nil {
		return def 
	}
	return v
}

// Get
func (h *TransactionHandler) GetTransactions(c *fiber.Ctx) error {

	// อ่าน query param 
	page := parseInt(c.Query("page", "1"), 1)   // default = 1
	limit := parseInt(c.Query("limit", "20"), 20) // default = 20

	// เรียก service 
	result, err := h.svc.List(page, limit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":        result.Data,
		"page":        result.Page,
		"limit":       result.Limit,
		"total":       result.Total,
		"total_pages": result.TotalPages,
	})
}
