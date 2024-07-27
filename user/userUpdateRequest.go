package user

type UserUpdateRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Whatsapp string `json:"whatsapp,omitempty"`
	Gender   string `json:"gender,omitempty"`
	Role     string `json:"role,omitempty"`
}
