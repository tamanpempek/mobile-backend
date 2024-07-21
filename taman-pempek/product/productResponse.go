package product

type ProductResponse struct {
	ID          uint64 `json:"id"`
	UserID      int    `json:"user_id"`
	CategoryID  int    `json:"category_id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Price       int    `json:"price"`
	Stock       int    `json:"stock"`
}
