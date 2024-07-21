package setting

import "mime/multipart"

type SettingUpdateRequest struct {
	Image       *multipart.FileHeader `form:"image,omitempty"`
	Description string                `form:"description,omitempty"`
	Email       string                `form:"email,omitempty"`
	Instagram   string                `form:"instagram,omitempty"`
	Website     string                `form:"website,omitempty"`
}
