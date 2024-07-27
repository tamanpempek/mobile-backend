package bank

type BankCreateRequest struct {
	UserID int    `json:"user_id" binding:"required"`
	Type   string `json:"type" binding:"required"`
	Name   string `json:"name" binding:"required"`
	Number string `json:"number" binding:"required"`
}
