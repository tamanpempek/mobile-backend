package setting

import (
	"errors"

	"gorm.io/gorm"
)

type SettingService interface {
	FindSettingByID(ID int) (Setting, error)
	UpdateSetting(ID int, setting SettingUpdateRequest) (Setting, error)
}

type service struct {
	settingRepository SettingRepository
}

func NewService(settingRepository SettingRepository) *service {
	return &service{settingRepository}
}

func (s *service) FindSettingByID(ID int) (Setting, error) {
	return s.settingRepository.FindSettingByID(ID)
}

func (s *service) UpdateSetting(ID int, settingRequest SettingUpdateRequest) (Setting, error) {
	setting, err := s.settingRepository.FindSettingByID(ID)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return Setting{}, errors.New("Setting not found")
	}
	if err != nil {
		return Setting{}, err
	}
	if settingRequest.Image != nil {
		setting.Image = settingRequest.Image.Filename
	}
	if settingRequest.Description != "" {
		setting.Description = settingRequest.Description
	}
	if settingRequest.Email != "" {
		setting.Email = settingRequest.Email
	}
	if settingRequest.Instagram != "" {
		setting.Instagram = settingRequest.Instagram
	}
	if settingRequest.Website != "" {
		setting.Website = settingRequest.Website
	}

	return s.settingRepository.UpdateSetting(setting)
}
