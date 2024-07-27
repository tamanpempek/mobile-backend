package product

import "mime/multipart"

type ProductUpdateRequest struct {
	CategoryID  int                   `form:"category_id,omitempty"`
	Name        string                `form:"name,omitempty"`
	Image       *multipart.FileHeader `form:"image,omitempty"`
	Description string                `form:"description,omitempty"`
	Price       int                   `form:"price,omitempty"`
	Stock       int                   `form:"stock,omitempty"`
}
