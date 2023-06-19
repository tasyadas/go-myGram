package models

import (
	"time"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"user_id"`
	PhotoID   uint   `gorm:"foreignKey:PhotoID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"photo_id"`
	Message   string `gorm:"not null;type:varchar(191)" json:"message" valid:"required~Message cannot be empty"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(c)

	if errCreate != nil {
		err = errCreate
		return
	}

	return
}
