package user

type UserCreateRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Whatsapp string `json:"whatsapp" binding:"required"`
	Gender   string `json:"gender" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
