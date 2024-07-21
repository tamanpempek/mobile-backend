package setting

import (
	"errors"

	"gorm.io/gorm"
)

type SettingRepository interface {
	FindSettingByID(ID int) (Setting, error)
	UpdateSetting(setting Setting) (Setting, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindSettingByID(ID int) (Setting, error) {
	var setting Setting
	err := r.db.First(&setting, ID).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Setting{}, errors.New("Setting not found")
	}
	return setting, err
}

func (r *repository) UpdateSetting(setting Setting) (Setting, error) {
	err := r.db.Save(&setting).Error
	return setting, err
}
