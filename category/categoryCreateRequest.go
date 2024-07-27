package category

type CategoryCreateRequest struct {
	Name string `json:"name" binding:"required"`
}
