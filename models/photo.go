package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Photo struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null;type:varchar(191)" json:"title" valid:"required~Title cannot be empty"`
	Caption   string `gorm:"not null;type:varchar(191)" json:"caption"`
	PhotoUrl  string `gorm:"not null;type:varchar(191)" json:"photo_url" valid:"required~Photo Url cannot be empty"`
	UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Photo) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(p)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
