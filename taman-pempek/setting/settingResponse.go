package setting

type SettingResponse struct {
	ID          uint64 `json:"id"`
	Image       string `json:"image"`
	Description string `json:"description"`
	Email       string `json:"email"`
	Instagram   string `json:"instagram"`
	Website     string `json:"website"`
}
