package bank

type BankUpdateRequest struct {
	Type   string `json:"type,omitempty"`
	Name   string `json:"name,omitempty"`
	Number string `json:"number,omitempty"`
}
