package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Sosmed struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null;type:varchar(191)" json:"name" valid:"required~Name cannot be empty"`
	SosmedUrl string `gorm:"not null;type:varchar(191)" json:"sosmed_url" valid:"required~Social Media Url cannot be empty"`
	UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (s *Sosmed) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(s)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
