package product

import "mime/multipart"

type ProductCreateRequest struct {
	UserID      int                  `form:"user_id" binding:"required"`
	CategoryID  int                  `form:"category_id" binding:"required"`
	Name        string               `form:"name" binding:"required"`
	Image       multipart.FileHeader `form:"image" binding:"required"`
	Description string               `form:"description" binding:"required"`
	Price       int                  `form:"price" binding:"required,number"`
	Stock       int                  `form:"stock" binding:"required,number"`
}
